---
title: Let's Encrypt

categories:
- http
---

# 简介

Let's Encrypt 是一个由非营利性组织 互联网安全研究小组（ISRG）提供的免费、自动化和开放的证书颁发机构（CA）。

简单的说，借助 Let's Encrypt 颁发的证书可以为我们的网站免费启用 HTTPS(SSL/TLS) 。

Let's Encrypt免费证书的签发/续签都是脚本自动化的，官方提供了几种证书的申请方式方法，点击此处 快速浏览。

官方推荐使用 Certbot 客户端来签发证书，这种方式可参考文档自行尝试，不做评价。

我这里直接使用第三方客户端 acme.sh 申请，据了解这种方式可能是目前 Let's Encrypt 免费证书客户端最简单、最智能的 shell 脚本，可以自动发布和续订 Let's Encrypt 中的免费证书。

<!--more-->


# 安装
1. 安装Certbot  
```shell
sudo snap install --classic certbot
```
2. 准备Certbot命令
```shell
sudo ln -s /snap/bin/certbot /usr/bin/certbot
```
3. 获取并安装证书
```shell
sudo certbot --nginx
```

# nginx 将https 转发到其它端口

```shell
server {  
listen 443;  
server_name example.com;

    ssl_certificate /path/to/ssl_certificate.crt;  
    ssl_certificate_key /path/to/ssl_certificate.key;  
  
    location / {  
        proxy_pass http://localhost:8080;  
        proxy_set_header Host $host;  
        proxy_set_header X-Real-IP $remote_addr;  
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;  
    }  
}
```
