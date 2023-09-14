---
title: kafak
categories: 
- docker
tags:
- docker-install-ware
- kafak
---

# docker 安装kafak

## 单机版
```
docker run -p 2181:2181 -p 9092:9092 --env ADVERTISED_HOST=192.168.11.36 --env ADVERTISED_PORT=9092 -d spotify/kafka
```

