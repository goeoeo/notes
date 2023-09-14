---
title: elasticsearch

categories: 
- docker

tags:
- docker-install-ware
---
# docker安装elasticsearch7.7.0


## 手动版安装
### 获取镜像
```
docker pull docker.elastic.co/elasticsearch/elasticsearch:7.7.0
```

### 运行elasticsearch
```
docker run --name=elastic770 -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" --network=network_172_19 --ip=172.19.0.101 -v /var/docker/elastic/data:/data docker.elastic.co/elasticsearch/elasticsearch:7.7.0
```

### 安装ik分词插件

```
docker exec -it elastic770 /bin/bash

./bin/elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v7.7.0/elasticsearch-analysis-ik-7.7.0.zip
```

## docker-compose安装

### Dockerfile
```
FROM docker.elastic.co/elasticsearch/elasticsearch:7.7.0
ENV VERSION=7.7.0

# https://github.com/medcl/elasticsearch-analysis-ik/releases
ADD https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v${VERSION}/elasticsearch-analysis-ik-$VERSION.zip /tmp/

RUN /usr/share/elasticsearch/bin/elasticsearch-plugin install -b file:///tmp/elasticsearch-analysis-ik-$VERSION.zip
RUN rm -rf /tmp/*
```

### docker-compose.yml
```
version: '3.5'

services:
    elastic770:
        build: ./
        container_name: elastic770
        hostname: elastic770
        volumes:
        - ./data:/usr/share/elasticsearch/data　　#这里将elasticsearch的数据文件映射本地，以保证下次如果删除了容器还有数据
        environment:
        - discovery.type=single-node
        ports:
        - "9200:9200"
```