---
title: 如何利用nginx进行自动化负载均衡

categories: 
- nginx

tags:
- 自动负载均衡
- Nginx
---

基于Docker + Consul + Nginx + Consul-Template + Registrator的服务自动负载均衡实现 
<!--more-->

## docker-composer相关知识点
* ports 暴露容器端口到主机的任意端口或指定端口, 用法如下
```
ports:
 
    - "80:80" # 绑定容器的80端口到主机的80端口
     
    - "9000:8080" # 绑定容器的8080端口到主机的9000端口
     
    - "443" # 绑定容器的443端口到主机的任意端口，容器启动时随机分配绑定的主机端口

```
* links 使得docker容器间通过指定的名字来和目标容器通信 
> 如果使用links则容器间通信 只能通过该使用容器的ip地址来通信或者通过宿主机的ip加上容器暴露出的端口号来通信 （这两种都有弊端）
```
links:
    -  consul_server_master:consul # 当前容器中使用consul 即可获取容器consul_server_master的ip地址,便于通信
```
* depends_on 决定容器的依赖, 也就是指定当前容器的启动顺序，必须依赖启动后才会启动
```
depends_on:
    - lb
    - registrator
```

## consul 相关知识点
* Client模式 就是客户端模式。是 Consul 节点的一种模式，这种模式下，所有注册到当前节点的服务会被转发到 Server，本身是不持久化这些信息。
* Server模式 表明这个 Consul 是个 Server ，这种模式下，功能和 Client 都一样，唯一不同的是，它会把所有的信息持久化的本地，这样遇到故障，信息是可以被保留的。同时可以同步多个consul客户端的数据
    > Server 节点之间的数据一致性保证协议使用的是 raft，而 zookeeper 用的 paxos，etcd采用的也是raft
```
 command: consul agent -server -bootstrap-expect 1 -advertise 192.168.1.181 -node consul_server_master -data-dir /tmp/data-dir -client 0.0.0.0 -ui
 # -bootstrap-expect=2 表示节点个数为2个
 # -node=consul-server-1 表示节点名称为consul-server-1
 # -client=0.0.0.0 表示允许连接的客户端 IP
 # -bind=10.211.55.2 表示服务端 IP为10.211.55.2
 # -datacenter=dc1 数据中心名称
 # -join=10.211.55.4 表示加入10.211.55.4节点的集群
```

* 服务发现协议 Consul 采用 http 和 DNS 协议，etcd 只支持 http 
* 服务注册 Consul 支持两种方式实现服务注册，一种是通过 Consul 的服务注册 Http API，由服务自己调用 API 实现注册，另一种方式是通过 json 格式的配置文件实现注册，将需要注册的服务以 json 格式的配置文件给出。Consul 官方建议使用第二种方式。
* 服务定义参数

| 环境变量Key   | 环境变量Value | 说明 |
| :---:   | :---: | :---: | 
|SERVICE_ID|web-001|可以为GUID或者可读性更强变量，保证不重复|
|SERVICE_NAME|web|如果ID没有设置，Consul会将name作为id，则有可能注册失败|
|SERVICE_TAGS|nodejs,web|服务的标签，用逗号分隔，开发者可以根据标签来查询一些信息|
|SERVICE_IGNORE|Boolean|是否忽略本Container，可以为一些不需要注册的Container添加此属性|


## 实现原理
* 通过 Nginx 自身实现负载均衡和请求转发
* 通过 Consul-template 的 config 功能实时监控 Consul 集群节点的服务和数据的变化；
  实时的用 Consul 节点的信息替换 Nginx 配置文件的模板，并重新加载配置文件
> Consul-template 和 nginx 必须安装在同一台机器上，因为 Consul-template 需要动态修改 nginx 的配置文件 nginx.conf，然后执行 nginx -s reload 命令进行路由更新，达到动态负载均衡的目的
 
*  registrator，它可以通过跟本地的 docker 引擎通信，来获取本地启动的容器信息，并且注册到指定的服务发现管理端。

数据流向： docker 容器数据-> registrator -> consul ->consul-template -> nginx

## 镜像构建
* Consul：consul:latest
* Registrator：gliderlabs/registrator:latest
* Nginx和Consul-template：liberalman/nginx-consul-template:latest
* Web: yeasy/simple-web:latest


## docker-compose.yml 
```
version: '3'
#backend web application, scale this with docker-compose scale web=3
services:
    web:
      image: yeasy/simple-web:latest
      environment:
        - SERVICE_80_NAME=my-web-server
      ports:
      - "80"
      depends_on:
          - lb
          - registrator  

    #load balancer will automatically update the config using consul-template
    lb:
      image: liberalman/nginx-consul-template:latest
      hostname: lb
      environment:
          - SERVICE_IGNORE=true
      links:
      -  consul_server_master:consul
      ports:
      - "80:80"
      depends_on:
          - consul_server_master
          
    consul_server_master:
      image:  consul:latest
      hostname: consul_server_master
      environment:
          - SERVICE_IGNORE=true
      ports:
      - "8500:8500"
      command: consul agent -server -bootstrap-expect 1 -node consul_server_master -data-dir /tmp/consul -client 0.0.0.0 -ui

    # listen on local docker sock to register the container with public ports to the consul service
    registrator:
      image: gliderlabs/registrator:master
      hostname: registrator
      environment:
          - SERVICE_IGNORE=true
      links:
      - consul_server_master:consul
      depends_on:
          - consul_server_master
      volumes:
      - "/var/run/docker.sock:/tmp/docker.sock"
      command: -internal consul://consul:8500
```
> 注意：我们使用的第三方镜像 liberalman/nginx-consul-template，Nginx 会把名称为 my-web-server的服务容器作为后台转发的目标服务器,当然你也可以自己制作镜像指定模板. 

## 测试步骤
1. docker-compose 模板所在目录，执行
```bash
$ sudo docker-compose up
```
2. 访问 http://localhost 可以看到一个 web 页面，提示实际访问的目标地址。
```
#2019-12-20 04:12:30: 19 requests from <LOCAL: 172.17.0.6> to WebServer <172.17.0.4>
```

3. 增加web 测试负载
```bash
$ sudo docker-compose scale web=3
```

4. 重复刷新http://localhost 观察ip变化


## 总结测试
最先我未去指定容器的启动顺序（deponds_on） 导致docker-compose up 后访问http://localhost 报错 nginx 502
原因是因为 registrator容器先于web容器启动，导致未将web容器注册到consul 所以生成的nginx配置文件中，没用web参与负载

本文搭建的consul为单节点版本 ， 集群版 [参考文章1](https://www.jianshu.com/p/fa41434d444a)

## 参考文章

1. [基于Docker + Consul + Nginx + Consul-Template的服务负载均衡实现](https://www.jianshu.com/p/fa41434d444a)
2. [用 consul + consul-template + registrator + nginx 打造真正可动态扩展的服务架构 ](https://my.oschina.net/xiaominmin/blog/1597660)
3. [8分钟学会Consul集群搭建及微服务概念 ](https://www.sohu.com/a/282625515_468635)