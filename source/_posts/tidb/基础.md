---
categories: 
- tidb


tags:
- tidb
---



# TiDB基础入门

TiDB 是一个分布式 NewSQL [SQL 、 NoSQL 和 NewSQL 的优缺点比较 ](https://www.jianshu.com/p/ed55f20e736a)数据库。它支持水平弹性扩展、ACID 事务、标准 SQL、MySQL 语法和 MySQL 协议，具有数据强一致的高可用特性，是一个不仅适合 OLTP 场景还适合 OLAP 场景的混合数据库。

<!--more-->



## 简介

*  高度兼容 MySQL 

  大多数情况下，无需修改代码即可从 MySQL 轻松迁移至 TiDB，分库分表后的 MySQL 集群亦可通过 TiDB 工具进行实时迁移。

* 水平弹性扩展 

  通过简单地增加新节点即可实现 TiDB 的水平扩展，按需扩展吞吐或存储，轻松应对高并发、海量数据场景。

* 分布式事务 TiDB 

  100% 支持标准的 ACID 事务。

* 真正金融级高可用  

  相比于传统主从 (M-S) 复制方案，基于 Raft 的多数派选举协议可以提供金融级的 100% 数据强一致性保证，且在不丢失大多数副本的前提下，可以实现故障的自动恢复 (auto-failover)，无需人工介入

* 一站式 HTAP 解决方案 

   TiDB 作为典型的 OLTP 行存数据库，同时兼具强大的 OLAP 性能，配合 TiSpark，可提供一站式 HTAP解决方案，一份存储同时处理OLTP & OLAP[OLAP、OLTP的介绍和比较](https://www.jianshu.com/p/b1d7ca178691)无需传统繁琐的 ETL 过程。

* 云原生 SQL 数据库  

  TiDB 是为云而设计的数据库 

  同 Kubernetes （[Kubernetes核心概念](https://www.jianshu.com/p/6326c7b4bc63) ）深度耦合，支持公有云、私有云和混合云，使部署、配置和维护变得十分简单。   

  TiDB 的设计目标是 100% 的 OLTP 场景和 80% 的 OLAP 场景，更复杂的 OLAP 分析可以通过 TiSpark 项目来完成。 TiDB 对业务没有任何侵入性，能优雅的替换传统的数据库中间件、数据库分库分表等 Sharding 方案。同时它也让开发运维人员不用关注数据库 Scale 的细节问题，专注于业务开发，极大的提升研发的生产力.    

## TiDB整体架构

![](基础/1766027-20190909160918250-1390331381.png)

### 三大组件

#### TIDB Server

TiDB Server 负责接收 SQL 请求，处理 SQL 相关的逻辑，并通过 PD 找到存储计算所需数据的 TiKV 地址，与 TiKV 交互获取数据，最终返回结果。 TiDB Server 是无状态的，其本身并不存储数据，只负责计算，可以无限水平扩展，可以通过负载均衡组件（如LVS、HAProxy 或 F5）对外提供统一的接入地址。

### PD Server

Placement Driver (简称 PD) 是整个集群的管理模块，其主要工作有三个： 一是存储集群的元信息（某个 Key 存储在哪个 TiKV 节点）；二是对 TiKV 集群进行调度和负载均衡（如数据的迁移、Raft group leader 的迁移等）；三是分配全局唯一且递增的事务 ID。

PD 是一个集群，需要部署奇数个节点，一般线上推荐至少部署 3 个节点。

#### TIKV Server

TiKV Server 负责存储数据，从外部看 TiKV 是一个分布式的提供事务的 Key-Value 存储引擎。存储数据的基本单位是 Region（区域），每个 Region 负责存储一个 Key Range （从 StartKey 到 EndKey 的左闭右开区间）的数据，每个 TiKV 节点会负责多个 Region 。TiKV 使用 Raft 协议做复制，保持数据的一致性和容灾。副本以 Region 为单位进行管理，不同节点上的多个 Region 构成一个 Raft Group，互为副本。数据在多个 TiKV 之间的负载均衡由 PD 调度，这里也是以 Region 为单位进行调度。



## 核心特性

* 水平扩展

  无限水平扩展是 TiDB 的一大特点，这里说的水平扩展包括两方面：计算能力和存储能力。TiDB Server 负责处理 SQL 请求，随着业务的增长，可以简单的添加 TiDB Server 节点，提高整体的处理能力，提供更高的吞吐。TiKV 负责存储数据，随着数据量的增长，可以部署更多的 TiKV Server 节点解决数据 Scale 的问题。PD 会在 TiKV 节点之间以 Region 为单位做调度，将部分数据迁移到新加的节点上。所以在业务的早期，可以只部署少量的服务实例，随着业务量的增长，按照需求添加 TiKV 或者 TiDB 实例。

* 高可用

  高可用是 TiDB 的另一大特点，TiDB/TiKV/PD 这三个组件都能容忍部分实例失效，不影响整个集群的可用性。

  

## TiDB原理与实现

TiDB 架构是 SQL 层和 KV 存储层分离，相当于 InnoDB 插件存储引擎与 MySQL 的关系。从下图可以看出整个系统是高度分层的，最底层选用了当前比较流行的存储引擎 RocksDB，RockDB 性能很好但是是单机的，为了保证高可用所以写多份，上层使用 Raft 协议来保证单机失效后数据不丢失不出错。保证有了比较安全的 KV 存储的基础上再去构建多版本，再去构建分布式事务，这样就构成了存储层 TiKV。有了TiKV，TiDB 层只需要实现 SQL 层，再加上 MySQL 协议的支持，应用程序就能像访问 MySQL 那样去访问 TiDB 了。



## TIDB vs Mysql

项目中使用mysql 目前出现的问题。单表数据量持续扩大目前某些表已达到500w级别,带来查询性能的下降，有些聚合的时间达到30s+ 。   

最先考虑的方案是对mysql 进行拆表。  拆表带来的问题 ： 

1. 修改业务代码。

2. 拆表后，总查询的需求需要借助新的组件，例如es、clickhouse等。

   

tidb优势: 完全兼容msyql 协议，不需要修改业务代码。  

分布式数据库，弹性伸缩，能自动完成类似于分表的操作，应用层不用关心。  

针对于 OLTP 场景，其tiflash存储引擎实现了列式存储。我本地测试,原来mysql中，20s的聚合sql 利用tiflash 可以将查询时间缩短至2s以内，数据量越大其优势越明显  

丰富的周边工具，官方提供的dm工具实现了全量同步+增量同步的方式，使得数据迁移平滑进行。



## 总结  

tidb 作为国内最流行的NewSql 数据库之一 能解决我们目前项目的遇到的问题，值得探索。 

