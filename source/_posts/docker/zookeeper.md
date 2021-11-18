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

