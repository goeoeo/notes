# docker 安装cassandra
## 单机版
```shell

docker run --name my-cassandra -d   --restart=always --network=compose_redis-sentinel --ip=172.15.12.22 cassandra:2.0.14
```