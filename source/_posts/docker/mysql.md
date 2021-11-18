---
categories: 
- docker
tags:
- docker-install-ware
- mysql
---

# docker 安装mysql

## 单机版 
```
docker run -it -d -p 3306:3306 --name docker_mysql -e MYSQL_ROOT_PASSWORD=123456  --restart=always  mysql 
docker run -it -d -p 13306:3306 --name newbilling_mysql -e MYSQL_ROOT_PASSWORD=p@ss52Dnb  --restart=always  mysql 
```

docker run -it -d -p 33063:3306 --name mysql5.6.47-1 -e MYSQL_ROOT_PASSWORD=123456  --restart=always  mysql:5.6.47 





```
docker run --name postgres -e POSTGRES_PASSWORD=p@ss52Dnb -p 5432:5432 --restart=always -d postgres
```