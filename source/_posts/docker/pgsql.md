---
categories:
- docker
tags:
- docker-install-ware
- pgsql
---

# docker 安装pgsql


## 单机版
```
docker run --name mypostgres -e POSTGRES_PASSWORD=p@ss52Dnb --restart=always -d -p 5432:5432 postgres 
```
