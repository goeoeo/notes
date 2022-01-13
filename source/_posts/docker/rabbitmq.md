---
categories: 
- docker
tags:
- docker-install-ware
- mysql
---


# Rabbitmq

## 单机版
`
docker run -d -p 5672:5672 -p 15672:15672 --name rabbitmq  --restart=always rabbitmq:3.8.3-management
`