---
title: git相关问题
categories: 
- git


---

# ssh: connect to host github.com port 22: Connection refused
检查
```
ssh -vT git@github.com
```
发现连接的是本地 

## 解决
```
nslookup github.com 8.8.8.8

# 加DNS 域名解析
sudo vim /etc/hosts 
```

