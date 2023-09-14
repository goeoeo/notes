---
title: MySQL问题集

categories: 
- mysql
tags:
- 索引
---
# 解决第一次连接MySQL连不上和连接速度慢
所谓反向解析是这样的：
mysql接收到连接请求后，获得的是客户 端的ip，为了更好的匹配mysql.user里的权限记录（某些是用hostname定义的）。
如果mysql服务器设置了dns服务器，并且客户端ip在dns上并没有相应的hostname，那么这个过程很慢，导致连接等待。

添加skip-name-resolve以后就跳过着一个过程了
```
[mysqld]
skip-name-resolve
```
<!--more-->

