---
title: etcd

categories: 
- docker
tags:
- docker-install-ware
- etcd
---
# docker安装ectd 

## etcd 启动参数说明
* -name：方便理解的节点名称，默认为 default，在集群中应该保持唯一，可以使用 hostname
* -data-dir：服务运行数据保存的路径，默认为 ${name}.etcd
* -snapshot-count：指定有多少事务（transaction）被提交时，触发截取快照保存到磁盘
* -heartbeat-interval：leader 多久发送一次心跳到 followers。默认值是 100ms
* -eletion-timeout：重新投票的超时时间，如果follower在该时间间隔没有收到心跳包，会触发重新投票，默认为 1000 ms
* -listen-peer-urls：和同伴通信的地址，比如 http://ip:2380，如果有多个，使用逗号分隔。需要所有节点都能够访问，所以不要使用 localhost
* -listen-client-urls：对外提供服务的地址：比如 http://ip:2379,http://127.0.0.1:2379，客户端会连接到这里和etcd交互
* -advertise-client-urls：对外公告的该节点客户端监听地址，这个值会告诉集群中其他节点
* -initial-advertise-peer-urls：该节点同伴监听地址，这个值会告诉集群中其他节点
* -initial-cluster：集群中所有节点的信息，格式为 node1=http://ip1:2380,node2=http://ip2:2380,…。需要注意的是，这里的 node1 是节点的--name指定的名字；后面的ip1:2380 是--initial-advertise-peer-urls 指定的值
* -initial-cluster-state：新建集群的时候，这个值为 new；假如已经存在的集群，这个值为existing
* -initial-cluster-token：创建集群的token，这个值每个集群保持唯一。这样的话，如果你要重新创建集群，即使配置和之前一样，也会再次生成新的集群和节点 uuid；否则会导致多个集群之间的冲突，造成未知的错误

<!--more-->

## 单机版
```
docker run -d  -p 2379:2379 -p 2380:2380 --name etcd quay.io/coreos/etcd:v3.5.0 /usr/local/bin/etcd -name qf2200-client0  -advertise-client-urls http://0.0.0.0:2379 -listen-client-urls http://0.0.0.0:2379 -listen-peer-urls http://0.0.0.0:2380  -initial-advertise-peer-urls http://0.0.0.0:2380 

 docker run --platform=linux/amd64  --rm -p 2379:2379 -p 2380:2380 --name etcd bitnami/etcd:3.5.4 
 

```


## 集群

测试环境：[play-with-docker](https://labs.play-with-docker.com/)

### docker-machine 准备换环境
play-with-docker太卡了，自己本地搭建测试环境
```
$ docker-machine create -d virtualbox manager1 && 
docker-machine create -d virtualbox worker1 && 
docker-machine create -d virtualbox worker2

$ docker-machine ls
NAME       ACTIVE   DRIVER       STATE     URL                         SWARM   DOCKER     ERRORS
manager1   -        virtualbox   Running   tcp://192.168.99.101:2376           v19.03.5   
worker1    -        virtualbox   Running   tcp://192.168.99.102:2376           v19.03.5   
worker2    -        virtualbox   Running   tcp://192.168.99.103:2376           v19.03.5 

```
  

### manager1 主机(node1 192.168.99.101) :
```
docker run -d --name etcd \
    -p 2379:2379 \
    -p 2380:2380 \
    --volume=etcd-data:/etcd-data \
    quay.io/coreos/etcd \
    /usr/local/bin/etcd \
    -name node1 \
    -data-dir=/etcd-data \
    -initial-advertise-peer-urls http://192.168.99.101:2380 \
    -listen-peer-urls http://0.0.0.0:2380 \
    -advertise-client-urls http://192.168.99.101:2379 \
    -listen-client-urls http://0.0.0.0:2379 \
    -initial-cluster-state new \
    -initial-cluster-token docker-etcd \
    -initial-cluster node1=http://192.168.99.101:2380,node2=http://192.168.99.102:2380,node3=http://192.168.99.103:2380
```

### worker1 主机（node2 192.168.99.102）：
```
docker run -d --name etcd \
    -p 2379:2379 \
    -p 2380:2380 \
    --volume=etcd-data:/etcd-data \
    quay.io/coreos/etcd \
    /usr/local/bin/etcd \
    -name node2 \
    -data-dir=/etcd-data \
    -initial-advertise-peer-urls http://192.168.99.102:2380 \
    -listen-peer-urls http://0.0.0.0:2380 \
    -advertise-client-urls http://192.168.99.102:2379 \
    -listen-client-urls http://0.0.0.0:2379 \
    -initial-cluster-state existing \
    -initial-cluster-token docker-etcd \
    -initial-cluster node1=http://192.168.99.101:2380,node2=http://192.168.99.102:2380,node3=http://192.168.99.103:2380

```

### worker2 主机（node3 192.168.99.103）：
```
docker run -d --name etcd \
    -p 2379:2379 \
    -p 2380:2380 \
    --volume=etcd-data:/etcd-data \
    quay.io/coreos/etcd \
    /usr/local/bin/etcd \
    -name node3 \
    -data-dir=/etcd-data \
    -initial-advertise-peer-urls http://192.168.99.103:2380 \ 
    -listen-peer-urls http://0.0.0.0:2380 \
    -advertise-client-urls http://192.168.99.103:2379 \
    -listen-client-urls http://0.0.0.0:2379 \
    -initial-cluster-state existing \
    -initial-cluster-token docker-etcd \
    -initial-cluster node1=http://192.168.99.101:2380,node2=http://192.168.99.102:2380,node3=http://192.168.99.103:2380

```
