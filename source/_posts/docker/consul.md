---
title: consul

categories: 
- docker

tags:
- docker-install-ware
- consul
---

# docker 安装consul

## 参数
```
8500 http 端口，用于 http 接口和 web ui
8300 server rpc 端口，同一数据中心 consul server 之间通过该端口通信
8301 serf lan 端口，同一数据中心 consul client 通过该端口通信
8302 serf wan 端口，不同数据中心 consul server 通过该端口通信
8600 dns 端口，用于服务发现
-bbostrap-expect 2: 集群至少两台服务器，才能选举集群leader
-ui：运行 web 控制台
-bind： 监听网口，0.0.0.0 表示所有网口，如果不指定默认为127.0.0.1，则无法和容器通信
-client ： 限制某些网口可以访问
```

## 单机版
```
docker run --name consul -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600 consul agent -server -bootstrap-expect 1 -ui -bind=0.0.0.0 -client=0.0.0.0
```