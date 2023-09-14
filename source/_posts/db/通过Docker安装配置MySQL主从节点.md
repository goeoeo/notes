---
title: 通过Docker安装配置MySQL主从节点
categories: 
- mysql

tags:
- mysql主从同步
---
## 拉取最新版mysql镜像
```bash
docker pull mysql 
```
## mysql主服务器
### 启动容器
```bash
docker run -p 3307:3306 --name mysql_master -e MYSQL_ROOT_PASSWORD=123456 -d mysql:latest
```
<!--more-->

### 修改配置并添加slave用户
1.进入容器
```
docker exec -it d45457d26cad /bin/bash
```

2.修改配置    
```bash
apt update && apt install  vim -y
vim /etc/mysql/my.cnf

[mysqld]
## 同一局域网内注意要唯一
server-id=100  
## 开启二进制日志功能，可以随便取（关键）
log-bin=mysql-bin
```
3.添加slave用户，用于slave读取binlog
```bash
mysql -uroot -p123456 

#mysql5.7以下使用 CREATE USER 'slave'@'%' IDENTIFIED BY '123456';
#mysql5.8修改了默认的认证方式为caching_sha2_password ,如果使用此方式创建slave用户，
#那么在后面slave连接master的时候就需要使用加密方式进行连接，否则会报错：Authentication requires secure connection

CREATE USER 'slave'@'%' IDENTIFIED WITH mysql_native_password BY '123456';

GRANT REPLICATION SLAVE, REPLICATION CLIENT ON *.* TO 'slave'@'%';
```

4.重启容器
```
docker restart d45457d26cad
```

## mysql从服务器

### 启动容器
```bash
docker run -p 3308:3306 --name mysql_slave -e MYSQL_ROOT_PASSWORD=123456 -d mysql:latest
```
### 修改配置
1.进入容器
```bash
docker exec -it 3e80d15d8b00 /bin/bash
```
2.修改配置    
```bash
apt update && apt install  vim -y
vim /etc/mysql/my.cnf

[mysqld]
## 设置server_id,注意要唯一
server-id=101  
## 开启二进制日志功能，以备Slave作为其它Slave的Master时使用
log-bin=mysql-slave-bin   
## relay_log配置中继日志
relay_log=edu-mysql-relay-bin  
```
3.重启容器
```bash
docker restart 3e80d15d8b00
```    
## 配置slave使其连上master    
```bash
change master to master_host='192.168.11.36', master_user='slave', master_password='123456', master_port=3307, master_log_file='mysql-bin.000001', master_log_pos= 605, master_connect_retry=30;

start slave 

show slave status \G 查看主从状态

```
> 参考文章
[基于Docker的Mysql主从复制搭建](https://www.cnblogs.com/songwenjie/p/9371422.html)

