---
title: pgsql
categories:
- docker
tags:
- docker-install-ware
- pgsql
---

# docker 安装pgsql


## 单机版
```
docker run --name postgres10.4 -e POSTGRES_PASSWORD=root -e POSTGRES_USER=root --restart=always -d  postgres:10.4 
```
