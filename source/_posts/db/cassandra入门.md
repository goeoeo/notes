---
title: cassandra入门

categories: 
- db
tags:
- cassandra
- 大数据
---

# 简介

## cassandra 端口

| 端口 | 协议 | 功能 |
| --- | --- | --- |
| 7000 | thrift | 节点之间通信 |
| 7001 | ssl thrift | 节点之间通信 |
| 9042 | cql | 客户端通信 |
| 9160 | thrift | 客户端通信 |

## cassandra.yaml 内容介绍
* cluster_name: 集群的名称  
* seed_provider: 种子节点  
* listen_address: 节点的监听地址  
* rpc_address: 节点的rpc地址  
* endpoint_snitch: 节点的snitch  
* num_tokens: 节点的token数量  
* disk_failure_policy: 磁盘失败策略  
* commitlog_sync: 提交日志同步策略  
* commitlog_sync_batch_window_in_ms: 提交日志同步批量窗口时间  
* commitlog_segment_size_in_mb: 提交日志分段大小  
* commitlog_total_space_in_mb: 提交日志总大小  
* commitlog_roll_when_full: 提交日志满了是否滚动  
* commitlog_roll_when_full_in_mb: 提交日志满了滚动的大小  
* commitlog_roll_when_full_in_ms: 提交日志满了滚动的时间  
* commitlog_roll_when_full_in_bytes: 提交日志满了滚动的字节数  



# Casssandra基本概念

## 数据模型

### 列
列是cassandra的基本单位，具有三个值：名称，值，时间戳   
在Casssandra中不需要预先定义列，只需要在KeySpace中定义列族，然后就可以开始写数据了。 

### 列族 ColumnFamily
列族相当于关系型数据库的表，包含了多个列的容器

列族的2种类型：
* 静态列族 字段名是固定的
* 动态列族 字段名是动态的

Row key:列族中每一行都是一个key，key是唯一的，用来标识这一行数据。 

主键：
* 主键是用来唯一标识一行数据的
* 主键可以由一列组成，也可以由多列组成

### 键空间 KeySpace
键空间是cassandra中最外层的容器，相当于关系型数据库中的数据库。  
键空间中包含了多个列族，每个列族中包含了多行数据。  
每个键空间中都有一个默认的列族，列族的名称为default。  

键空间创建的时候可以指定一些属性：副本因子，副本策略，Durable_write(是否启用commitLog机制)  

副本因子：
* 副本因子是用来指定每个键空间中副本的数量的
* 副本因子的默认值是1
* 副本因子可以在键空间创建的时候指定，也可以在键空间创建之后修改

副本策略：
* 副本策略是用来指定每个键空间中副本的分布策略的
* 副本策略的默认值是SimpleStrategy
* 副本策略可以在键空间创建的时候指定，也可以在键空间创建之后修改

Durable_write：
* Durable_write是用来指定每个键空间中是否启用commitLog机制的
* Durable_write的默认值是true
* Durable_write可以在键空间创建的时候指定，也可以在键空间创建之后修改


### 副本
副本就是把数据存储到多个节点上，来提高容错性

### 节点
存储数据的机器

### 数据中心
多台机器组成一个数据中心

### 集群
Cassandra数据库是为跨越多条主机共同工作，对用户呈现为一个整体的分布式系统设计的。Cassandra最外层容器被称为集群。
Cassandra将集群中的节点组成一个环（ring）,然后把数据分配到集群中的节点上

## 数据类型
* 数值类型
* 文本类型
* 时间类型
* 标识符类型 例如UUID
* 集合类型 例如set,list,map
* 其它基本类型 例如boolean,blob,inet,counter
* 用户自定义类型

## CQL Shell客户端
CQL Shell是Cassandra提供的一个命令行客户端，用于执行CQL语句。
CQL Shell可以在Cassandra的安装目录下的bin目录中找到。
CQL Shell的使用方法如下：
1. 进入Cassandra的安装目录下的bin目录
2. 执行cqlsh命令
3. 输入用户名和密码
4. 执行CQL语句
5. 退出CQL Shell

```
cqlsh dev01-zoocassa0 9042
```


