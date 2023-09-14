---
title: redis
categories: 
- docker
tags:
- docker-install-ware
- redis
---

# docker 安装redis


## 参数
* appendonly 是否开启数据持久化
* requirepass 用户名密码 默认:不使用
## 单机版

```
docker run --name redis -p 16000:6379 -d --restart=always -v /usr/local/docker/redis/redis.conf:/etc/redis/redis.conf  redis:latest redis-server /etc/redis/redis.conf --appendonly yes 
```

## 主从版
```
docker run --name redis -p 6379:6379 -d --restart=always -v /usr/local/docker/redis/redis.conf:/etc/redis/redis.conf  redis:latest redis-server /etc/redis/redis.conf --appendonly yes --requirepass "123456"
docker run --name redis -p 6380:6379 -d --restart=always -v /usr/local/docker/redis/redis-slave.conf:/etc/redis/redis.conf  redis:latest redis-server /etc/redis/redis.conf --appendonly yes --requirepass "123456"
```