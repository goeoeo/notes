---
categories: 
- docker
tags:
- docker-install-ware
- zookeeper
---

# docker 安装zookeeper 


## 单机版
```
docker run -d -p 2181:2181 --name some-zookeeper --restart=always zookeeper
```

## zookeeper 简单使用
```shell
cd /apache-zookeeper-3.7.0-bin/bin && zkCli.sh

# 查看节点 
ls /server


```