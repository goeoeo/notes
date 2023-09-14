---
title: 一站式学习redis从入门到高可用分布式实践

categories: 
- redis

tags:
- redis 
---



# 第1章到第3章
## redis 特性
* 速度快
* 持久化
* 多种数据结构
* 支持多种编程语言
* 功能丰富
* 简单
* 主从复制
* 高可用,分布式
<!--more-->

## 通用命令
* keys *: 列出所有的key ,不建议线上使用
* dbsize :统计所有的key数量
* exists :检查某个key是否存在
* del: 删除某个key
* expire key seconds :key 在seconds秒后过期
* ttl key :查询key剩余的过期时间
* persist key:去掉key的过期时间
* type key :返回key的类型
* info memory: 占用内存
* config get : 获取当前生效的配置
* flushall :清空数据库


## 单线程
* 一次只运行长慢的命令
* 拒绝长（慢）的命令


# 第4章 Redis持久化的取舍与选择
## 持久化的作用
### 什么是持久化
redis所有数据保持在内存中，对数据的更新将异步的保存到磁盘上  

### 持久化的实现方式
#### 快照
在某个时间将数据进行完全的拷贝  
* Mysql Dump
* Redis RDB

#### 写日志
记录命令日志  
* Mysql Binlog
* Hbase Hlog
* Redis AOF

## RDB
### 什么是RDB
* 二进制文件
* 复制媒介

### 触发机制
#### 主要三种触发方式
##### save (同步)　
* 命令执行过程中redis处于阻塞状态
* 文件策略，如存在老的RDB文件，新替换老
* O(N)

##### bgsave (异步)
* bgsave 命令执行的时候redis 会fork出一个子进程，由这个子进程去执行RDB文件的生成

|命令|save|bgsave|
|:---|:---|:---|
|IO类型|同步|异步|
|阻塞|是|是（阻塞发生在fork）,fork的过程很快|
|复杂度|O(n)|O(n)|
|优点|不会消耗额外内存|不阻塞客户端命令|
|缺点|阻塞客户端命令|需要fork,消耗内存|


##### 自动

|配置|seconds|changes|说明|
|:---|:---|:---|:---|
|save|900|1|900秒内发生1次数据变动，执行bgsave|
|save|300|10|300秒内发生10次数据变动，执行bgsave|
|save|60|10000||
以上条件满足任一条件都会执行bgsave  

1.redis默认配置文件 
```
save 900 1
save 300 10
save 60 10000
dbfilename dump.rdb
dir ./
stop-writes-on-bgsave-error yes
rdbcompression yes
rdbchecksum yes

```
2.redis推荐配置  

```
dbfilename dump-${port}.rdb
dir /bigdiskpath 
stop-writes-on-bgsave-error yes
rdbcompression yes
rdbchecksum yes
```

#### 不容忽视的方式
* 全量复制
* debug reload
* shutdown 

### RDB 总结
* RDB是Redis内存到硬盘的快照，用于持久化
* save通常会阻塞redis
* bgsave不会阻塞redis,但会fork新进程
* save自动配置满足任一就会被执行
* 有些触发机制不容忽视


## AOF
### RDB的问题
* 耗时、耗性能
* 不可控、容易丢失数据

### 什么是AOF
日志原理，每执行一条命令，都会写入到aof文件中
具体过程：  
1. 写命令刷新到缓冲区
2. 缓冲区的命令会根据一定的策略fsync到硬盘的AOF文件中    

### AOF三种策略
* always : 每执行一条命令就会执行fsync
* everysec （默认值）: 每秒写入到硬盘中
* no : 根据操作系统的策略决定的
|命令|always| everysec|no|
|:---|:---|:---|:---|
|优点|不丢失数据|每秒一次fsync丢1秒数据|不用管|
|缺点|IO开销大一般的sata盘只有几百TPS|丢1秒数据|不可控|

### AOF重写

#### 为什么要进行AOF重写
把过期的、没有用的、重复的、 以及一些可以优化的命令进行化简
* 减少磁盘的占用量
* 加速恢复速度

#### AOF重写实现的两种方式
##### bgrewriteaof
##### AOF重写配置
|配置名|含义|
|:---|:---|
|auto-aof-rewrite-min-size|AOF文件重写需要的尺寸|
|auto-aof-rewrite-percentage|AOF文件增长率（(aof_current_size-aof_base_size)/aof_base_size）|

以上两个条件同时满足才会触发aof自动重写

|统计名|含义|
|:---|:---|
|aof_current_size|AOF当前尺寸（单位：字节）|
|aof_base_size|AOF上次启动和重写的尺寸（单位：字节）|


### AOF 配置
```
appendonly yes
appendfilename "appendonly-${port}.aof"
appendfsync everysec
dir /bigdiskpath
no-appendfsync-on-rewrite yes  # 在进行aof重写的时候，不进行aof操作
auto-aof-rewrite-min-size 100
auto-aof-rewrite-percentage 64mb


```

## RDB和AOF的选择

### RDB与AOF比较
|命令|RDB|AOF|
|:---|:---|:---|
|启动优先级|低|高|
|体积|小|大|
|恢复速度|块|慢|
|数据安全|丢数据|根据策略决定|
|轻重|重|轻|

### RDB最佳策略
* "关"   
在主从进行全量复制的时候会生成RDB文件，所以关是关不绝对的

* 集中管理  
在固定的时间点进行备份的时候可以用

* 主、从
一定要开的情况下建议主关，从开

### AOF最佳策略
* 开  
redis作为存储的时候建议开启  
redis作为缓存的时候可以不开启

* AOF重写集中管理  
避免单机多部署的时候aof重写集中发生，占满内存

* aof重写策略推荐everysec 
只会丢失1s的数据，不会频繁将buffer中的数据fsync到硬盘

### 最佳策略
* 小分片
* 缓存或存储
* 监控（硬盘、内存、负载、网络）
* 足够的内存

# 第5章 常见持久化开发运维问题
## fork操作
* 同步操作
* 与内存量息息相关：内存越大、耗时越长（与机器类型有关）
* info:last_fork_usec

### 改善fork 
* 优先使用物理机或者高效支持fork操作的虚拟化技术
* 控制Redis实例最大可用内存：maxmemory
* 合理配置Linux内存分配策略：vm.overcommit_memory=1
* 降低fork频率：例如放宽AOF重写自动触发机制，不必要的全量复制

## 进程外开销
### CPU
* 开销 RDB和AOF文件生成，属于CPU密集型
* 优化 不做CPU绑定，不和CPU密集型部署

### 内存
* 开销 fork内存开销，copy-on-write
* 优化 echo never > /sys/kernel/mm/transparent_hugepage/enabled

### 硬盘
* 开销： AOF和RDB文件写入，可以结合iostat,iotop分析
* 优化   
不要和高硬盘负载服务器部署一起，存储服务、消息队列等  
no-appendfsync-on-rewrite=yes  
根据写入量决定磁盘类型 例如ssd  
单机多实例持久化文件目录可以考虑分盘

## AOF追加阻塞

## 单机多部署实例


# 第6章 Redis主从复制原理以及优化

## 什么是主从复制
### 单机有什么问题？
* 机器故障
* 容量瓶颈
* QPS瓶颈 

### 主从复制的作用
* 数据副本
* 扩展读性能

### 简单总结
* 一个master可以有多个slave
* 一个slave只能有一个master
* 数据流向是单向的，master到slave

## 复制的配置
### replicaof 命令
* replicaof 127.0.0.1 6379 
* replicaof no one
> 任何一个节点成为从节点的时候，都会将自身节点的数据进行清除

### 配置
```
replicaof ip port
replica-read-only yes
```
### 比较
|方式|命令|配置|
|:---|:---|:---|
|优点|无需重启|统一配置|
|缺点|不便于管理|需要重启|


## 全量复制和部分复制
### 全量复制过程
1. slave: psync ? -1
2. master: FULLRESYNC {runId} {offset}
3. slave: save masterInfo
4. master: 执行bgsave 生成rdb文件,send RDB to slave
5. master: send buffer 
6. slave: flush old data
7. slave: load RDB

### 全量复制开销 
1. bgsave时间
2. RDB文件网络传输
3. 从节点清空数据时间
4. 从节点加载RDB时间
5. 可能的AOF重写时间

### 部分复制
部分复制主要是Redis针对全量复制的过高开销做出的一种优化措施；  
如果出现网络闪断或者命令丢失等异常情况时，从节点会向主节点要求补发丢失的命令数据，如果主节点的复制积压缓冲区内存在这部分数据则直接发送给从节点


## 故障处理
### slave 故障
### master 故障

## 开发运维常见问题

### 读写分离
1. 读流量分摊到节点
2. 可能遇到的问题  
* 复制数据延迟
* 读到过期的数据
* 从节点故障

### 配置不一致
1. 例如maxmemory不一致：丢失数据
### 主从配置不一致
### 规避全量复制
1. 第一次全量复制  
* 第一次不可避免
* 小主节点、低峰
2. 节点运行ID不匹配
* 主节点重启（运行ID变化）
* 故障转移，列如哨兵或集群
3. 复制积压缓冲区不足
* 网络中断，部分复制无法满足
* 增大复制缓冲区配置rel_backlog_size, 网络“增强”，默认为1M,可以配置成10M

### 规避复制风暴
1. 单节点复制风暴  
* 问题： 主节点重启，多从节点复制
* 解决： 更换复制拓扑

2. 单机器复制风暴
* 机器宕机后，大量全量复制
* 主节点分散多机器



# 第7章 Redis sentine (哨兵)

## 简介
Redis Sentinel为Redis提供了很高的可用性，在实践中，这意味着你可以部署一个可以解决非人为干预导致节点故障的Redis集群系统。Redis Sentinel还提供了其他的功能：如监控，通知和客户端配置服务的提供方。下面列出来了Redis Sentinel的功能列表：  
* 监控：Sentinel能够监控master节点或slave节点是否处于按照预期工作的状态。
* 通知：Sentinel能够通过api通知系统管理原，其他的计算机程序，Redis实例运行过程中发生了错误。
* 自动故障转移：如果Redis的master节点出现问题，Sentinel能够启动一个故障转移处理，该处理会将一个slave节点提升为master节点，其他的slave节点则会自动配置成新的master节点的slave节点，如果原来的master重新正常启动后，也会成为该新Master的slave节点。
* 客户端配置提供者：Sentinel可作为客户端服务发现的一个权威来源，客户端通过连接到Sentinel来请求当前的Redis Master节点，如果Master节点发生故障，Sentinel将会提供新的master地址。

Redis Sentinel 是一个分布式系统， 你可以在架构中运行多个 Sentinel 进程，这些进程通过相互通讯来判断一个主服务器是否断线，以及是否应该执行故障转移。
在配置Redis Sentinel时，至少需要有1个Master和1个Slave。当Master失效后,Redis Sentinel会报出失效警告，并使用投票协议（agreement protocols）
来决定是否执行自动故障迁移， 以及选择哪个从服务器作为新的主服务器，并提供读写服务;当失效的Master恢复后，Redis Sentinel会自动识别，
将Master自动转换为Slave并完成数据同步。通过Redis Sentinel可以实现Redis零手工干预并且短时间内进行M-S切换，减少业务影响时间。  

虽然 Redis Sentinel 释出为一个单独的可执行文件 redis-sentinel ， 但实际上它只是一个运行在特殊模式下的 Redis 服务器， 你可以在启动一个普通 Redis 服务器时通过给定 --sentinel 选项来启动 Redis Sentinel。

## 主从复制高可用的问题
* 手动故障转移
* 写能力和存储能力受限

## 架构说明
sentinel 主要作用是监控、故障转移、通知
一个sentinel可以监控多套master,slave

### Redis Sentinel故障转移
1. 多个sentinel发现并确认master有问题
2. 选举出一个sentinel作为领导
3. 选出一个slave作为master
4. 通知其余slave成为新的master的slave
5. 通知客户端主从变化

## 安装配置
sentinel默认监听节点是26379
1. 配置开启主从节点
2. 配置开启sentinel监控节点（sentinel是特殊的redis）
3. 实际应该是多台机器
4. 详细配置节点

### redis master节点配置
```
port 7000
pidfile /var/run/redis-7000.pid
logfile "7000.log"
dir "/data"
```
### redis slave1节点配置
```
port 7001
pidfile /var/run/redis-7001.pid
logfile "7001.log"
dir "/data"
replicaof 127.0.0.1 7000

```

### redis sentinel主要配置
```
port ${port}
dir "/data"
logfile "${port}.log"
# 设置判定mymaster 失效的monitor 个数 
sentinel monitor mymaster 127.0.0.1 7000 2 
sentinel down-after-millseconds mymaster 30000
sentinel parallel-syncs mymaster 1
sentinel failover-timeout mymaster 180000
```
## sentinel中的3个定时任务

### 每10秒每个sentinel对master和slave 执行info
* 发现slave节点
* 确认主从关系

### 每两秒每个sentinel通过master节点的channel交换信息（pub/sub）
* 通过 _sentinel_:hello 频道交换信息
* 交换对节点的“看法”和自身信息

### 每1秒每个sentinel对其他sentinel和redis进ping
* 心跳检测，失败判定依据


## 主观下线和客观下线

### 主观下线
在默认情况下，Sentinel会以每秒一次的频率向所有与它创建了命令连接的实例（主服务器、从服务器、其它Sentinel）发送ping命令，并通过相应的回复判断实例是否在线。  
如果实例的相应时间超过down-after-milliseconds设置的时间（毫秒），或者响应无效，那么Sentinel会将此实例标记为 主观下线 状态

### 客观下线
当Sentinel将一个Master判断为主观下线之后，为了确定这个Master是否真的下线了，它会向同样监视这个Master的其它Sentinel进行询问，看它们是否也认为Master已经进入下线状态（可以是主观下线也可以是客观下线）。
当Sentinel从其它Sentinel那里接收到足够数量的已下线判断之后，Sentinel就会将从服务器判断为客观下线，并对主服务器执行故障转移。

## 客户端连接

## 实现原理

## 常见开发运维问题

##  总结
* Redis Sentinel 的Sentinel节点数应该大于等于3且最好为奇数
* Redis Sentinel 中的数据节点与普通数据节点没有区别
* 客户端初始化时连接的是Sentinel节点集合，不再是具体的Redis节点，但Sentinel只是配置中心而不是代理
* Redis Sentinel 通过三个定时任务实现了Sentinel节点对于主节点、从节点、其余Sentinel节点的监控
* Redis Sentinel 在对节点做失败判定时分为主观下线和客观下线
* 看懂Redis Sentinel 故障转移日志对于Redis Sentinel以及排查问题非常有帮助
* Redis Sentinel 实现读写分离高可用可以依赖Sentinel节点的消息通知，获取Redis数据节点的状态变化


# 第8章 Redis Cluster

## 数据分布概论

### 数据分布对比
|分布方式|特点|典型产品|
|:---|:---|:---|
|哈希分布|数据分散度高键值分布业务无关，无法顺序访问，支持批量操作|一致性哈希Memcache,Redis cluster,其他缓存产品|
|顺序分布|数据分散度易倾斜，键值业务相关，可顺序访问，支持批量操作|BigTable HBase|

### 哈希分布方式
#### 节点取余分区
* 客户端分片：哈希+取余
* 节点伸缩： 数据节点变化，导致数据迁移
* 迁移数量和添加节点数量有关：建议翻倍扩容
#### 一致性哈希分区
* 客户端分片：哈希+顺时针（优化取余）
* 节点伸缩： 只影响临近节点，但是还是有数据迁移
* 翻倍伸缩： 保证最小迁移数据和负载均衡

#### 虚拟槽分区
* 预设虚拟槽：每个槽映射一个数据子集，一般比节点数大
* 良好的哈希函数：例如CRC16
* 服务端管理节点、槽、数据 例如：Redis Cluster

## Cluster节点主要配置
```
cluster-enabled yes
cluster-node-timeout 15000
cluster-config-file "nodes.conf"
cluster-require-full-coverage yes
```