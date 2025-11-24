---
title: kafak
categories: 
- docker
tags:
- docker-install-ware
- kafak
---

# docker 安装kafak

## 单机版
```
docker run -p 2181:2181 -p 9092:9092 --env ADVERTISED_HOST=192.168.11.36 --env ADVERTISED_PORT=9092 -d spotify/kafka
```



# 测试
## 创建名为 test-topic 的主题，1个分区，1个副本
```
/opt/kafka_2.12-2.3.0/bin/kafka-topics.sh --bootstrap-server localhost:9092 \
    --create \
    --topic test-topic \
    --partitions 2 \
    --replication-factor 1

/opt/kafka_2.12-2.3.0/bin/kafka-topics.sh --bootstrap-server localhost:9092 --list
```



## 终端1：启动消费者（等待消息）
```
/opt/kafka_2.12-2.3.0/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 \
    --topic test-topic \
    --from-beginning
```

## 终端2：启动生产者
```
/opt/kafka_2.12-2.3.0/bin/kafka-console-producer.sh --bootstrap-server localhost:9092 \
    --topic test-topic
```

## 在生产者终端输入消息，在消费者终端查看接收情况