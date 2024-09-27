# docker 安装memcached
## 单机版
```shell

docker run --name my-memcache -d   --restart=always --network=compose_redis-sentinel --ip=172.15.12.21 memcached 
```