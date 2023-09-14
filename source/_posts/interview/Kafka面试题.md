---
title: Kafka面试题
categories: 
- interview
tags:
- kafka面试题
---
## Kafka的用途有哪些？使用场景如何？

1. 日志收集：一个公司可以用Kafka可以收集各种服务的log，通过kafka以统一接口服务的方式开放给各种consumer，例如Hadoop、Hbase、Solr等  
2. 消息系统：解耦和生产者和消费者、缓存消息等
3. 用户活动跟踪：Kafka经常被用来记录web用户或者app用户的各种活动，如浏览网页、搜索、点击等活动，这些活动信息被各个服务器发布到kafka的topic中，然后订阅者通过订阅这些topic来做实时的监控分析，或者装载到Hadoop、数据仓库中做离线分析和挖掘  
4. 运营指标：Kafka也经常用来记录运营监控数据。包括收集各种分布式应用的数据，生产各种操作的集中反馈，比如报警和报告  
5. 流式处理：比如spark streaming和storm
6. 事件源

<!--more-->

## Kafka中的ISR、AR又代表什么？ISR的伸缩又指什么
1. 分区中的所有副本统称为AR（Assigned Repllicas）。所有与leader副本保持一定程度同步的副本（包括Leader）组成ISR（In-Sync Replicas），ISR集合是AR集合中的一个子集。与leader副本同步滞后过多的副本（不包括leader）副本，组成OSR(Out-Sync Relipcas),由此可见：AR=ISR+OSR。
2. ISR集合的副本必须满足:   副本所在节点必须维持着与zookeeper的连接；副本最后一条消息的offset与leader副本最后一条消息的offset之间的差值不能超出指定的阈值  
3. 每个分区的leader副本都会维护此分区的ISR集合，写请求首先由leader副本处理，之后follower副本会从leader副本上拉取写入的消息，这个过程会有一定的延迟，导致follower副本中保存的消息略少于leader副本，只要未超出阈值都是可以容忍的  
4. ISR的伸缩指的是Kafka在启动的时候会开启两个与ISR相关的定时任务，名称分别为“isr-expiration"和”isr-change-propagation".。isr-expiration任务会周期性的检测每个分区是否需要缩减其ISR集合。  

## Kafka中的HW、LEO、LSO、LW等分别代表什么？
1. HW是High Watermak的缩写，俗称高水位，它表示了一个特定消息的偏移量（offset），消费之只能拉取到这个offset之前的消息。
2. LEO是Log End Offset的缩写，它表示了当前日志文件中下一条待写入消息的offset。  
3. LSO特指LastStableOffset。它具体与kafka的事物有关。消费端参数——isolation.level,这个参数用来配置消费者事务的隔离级别。字符串类型，“read_uncommitted”和“read_committed”。  
4. LW是Low Watermark的缩写，俗称“低水位”，代表AR集合中最小的logStartOffset值，副本的拉取请求（FetchRequest）和删除请求（DeleteRecordRequest）都可能促使LW的增长。   

## Kafka中是怎么体现消息顺序性的？具体的可参考这篇博文
[https://www.cnblogs.com/windpoplar/p/10747696.html](https://www.cnblogs.com/windpoplar/p/10747696.html) 
 1. 一个 topic，一个 partition，一个 consumer，内部单线程消费，单线程吞吐量太低，一般不会用这个。
 2. 写 N 个内存 queue，具有相同 key 的数据都到同一个内存 queue；然后对于 N 个线程，每个线程分别消费一个内存 queue 即可，这样就能保证顺序性。  

## Kafka中的分区器、序列化器、拦截器是否了解？它们之间的处理顺序是什么？
 拦截器 -> 序列化器 -> 分区器 
 
## Kafka生产者客户端的整体结构是什么样子的？
[https://zhuanlan.zhihu.com/p/94412266](https://zhuanlan.zhihu.com/p/94412266)

## 消费组中的消费者个数如果超过topic的分区，那么就会有消费者消费不到数据”这句话是否正确？如果不正确，那么有没有什么hack的手段？
不正确，通过自定义分区分配策略，可以将一个consumer指定消费所有partition。  

## 消费者提交消费位移时提交的是当前消费到的最新消息的offset还是offset+1? 
offset+1 

## 有哪些情形会造成重复消费？
消费者消费后没有commit offset(程序崩溃/强行kill/消费耗时/自动提交偏移情况下unscrible)    

## 那些情景下会造成消息漏消费？
先提交offset，后消费，有可能造成数据的重复   

## KafkaConsumer是非线程安全的，那么怎么样实现多线程消费？参考： 
[https://blog.csdn.net/clypm/article/details/80618036](https://blog.csdn.net/clypm/article/details/80618036)
1.每个线程维护一个KafkaConsumer  
2.维护一个或多个KafkaConsumer，同时维护多个事件处理线程(worker thread)  

## 简述消费者与消费组之间的关系 
消费者从属与消费组，消费偏移以消费组为单位。每个消费组可以独立消费主题的所有数据，同一消费组内消费者共同消费主题数据，每个分区只能被同一消费组内一个消费者消费。  

## 当你使用kafka-topics.sh创建（删除）了一个topic之后，Kafka背后会执行什么逻辑？  
1. 会在zookeeper中的/brokers/topics节点下创建一个新的topic节点，如：/brokers/topics/first
2. 触发Controller的监听程序
3. kafka Controller 负责topic的创建工作，并更新metadata cache

## topic的分区数可不可以增加？如果可以怎么增加？如果不可以，那又是为什么？
可以增加 
bin/kafka-topics.sh --zookeeper localhost:2181/kafka --alter --topic topic-config --partitions 3  

## topic的分区数可不可以减少？如果可以怎么减少？如果不可以，那又是为什么？
不可以减少，被删除的分区数据难以处理。 

## 创建topic时如何选择合适的分区数？
[https://blog.csdn.net/weixin_42641909/article/details/89294698](https://blog.csdn.net/weixin_42641909/article/details/89294698) 

1.创建一个只有1个分区的topic
2.测试这个topic的producer吞吐量和consumer吞吐量。
3.假设他们的值分别是Tp和Tc，单位可以是MB/s。
4.然后假设总的目标吞吐量是Tt，那么分区数=Tt / max（Tp，Tc）
例如：producer吞吐量=5m/s；consumer吞吐量=50m/s，期望吞吐量100m/s；  
分区数=100 / 50 =2分区  
分区数一般设置为：3-10个   

## Kafka目前有那些内部topic，它们都有什么特征？各自的作用又是什么？ 
consumer_offsets 以下划线开头，保存消费组的偏移  

## 优先副本是什么？它有什么特殊的作用？
优先副本 会是默认的leader副本 发生leader变化时重选举会优先选择优先副本作为leader 

## Kafka有哪几处地方有分区分配的概念？简述大致的过程及原理。参考：
[https://blog.csdn.net/weixin_42641909/article/details/89294698](https://blog.csdn.net/weixin_42641909/article/details/89294698)

Kafka提供的两种分配策略： range和roundrobin，由参数partition.assignment.strategy指定，默认是range策略。  
当以下事件发生时，Kafka 将会进行一次分区分配：  
同一个 Consumer Group 内新增消费者  
消费者离开当前所属的Consumer Group，包括shuts down 或 crashes  
订阅的主题新增分区  

### Range strategy
Range策略是对每个主题而言的，首先对同一个主题里面的分区按照序号进行排序，并对消费者按照字母顺序进行排序。
### RoundRobin strategy
使用RoundRobin策略有两个前提条件必须满足：
1. 同一个Consumer Group里面的所有消费者的num.streams必须相等；
2. 每个消费者订阅的主题必须相同。
RoundRobin策略的工作原理：将所有主题的分区组成 TopicAndPartition 列表，然后对 TopicAndPartition 列表按照 hashCode 进行排序   

## 简述Kafka的日志目录结构。Kafka中有那些索引文件？ 参考：
[https://zhuanlan.zhihu.com/p/94412266](https://zhuanlan.zhihu.com/p/94412266)

每个分区对应一个文件夹，文件夹的命名为topic-0，topic-1，内部为.log和.index文件 以及 .timeindex leader-epoch-checkpoint

## 如果我指定了一个offset，Kafka怎么查找到对应的消息？参考
[https://blog.csdn.net/moakun/article/details/101855315](https://blog.csdn.net/moakun/article/details/101855315)

1. 通过文件名前缀数字x找到该绝对offset 对应消息所在文件  
2. offset-x为在文件中的相对偏移
3. 通过index文件中记录的索引找到最近的消息的位置
4. 从最近位置开始逐条寻找

##  如果我指定了一个timestamp，Kafka怎么查找到对应的消息？
[https://blog.csdn.net/moakun/article/details/101855315](https://blog.csdn.net/moakun/article/details/101855315)
原理同上 但是时间的因为消息体中不带有时间戳 所以不精确

## kafka过期数据清理
日志清理保存的策略只有delete和compact两种  
log.cleanup.policy=delete启用删除策略  
log.cleanup.policy=compact启用压缩策略  

## Kafka中的幂等是怎么实现的
Producer的幂等性指的是当发送同一条消息时，数据在Server端只会被持久化一次，数据不丟不重，但是这里的幂等性是有条件的：  
1. 只能保证Producer在单个会话内不丟不重，如果Producer出现意外挂掉再重启是无法保证的（幂等性情况下，是无法获取之前的状态信息，因此是无法做到跨会话级别的不丢不重）。
2. 幂等性不能跨多个Topic-Partition，只能保证单个Partition内的幂等性，当涉及多个 Topic-Partition时，这中间的状态并没有同步。

## kafka事务。分享一篇大佬讲kafka事务的博客，这一篇讲的更深入：
[http://matt33.com/2018/11/04/kafka-transaction/](http://matt33.com/2018/11/04/kafka-transaction/)

Kafka从0.11版本开始引入了事务支持。事务可以保证Kafka在Exactly Once语义的基础上，生产和消费可以跨分区和会话，要么全部成功，要么全部失败。  

### Producer事务
为了实现跨分区跨会话的事务，需要引入一个全局唯一的Transaction ID，并将Producer获得的PID和Transaction ID绑定。这样当Producer重启后就可以通过正在进行的Transaction ID获得原来的PID。  

为了管理Transaction，Kafka引入了一个新的组件Transaction Coordinator。Producer就是通过和Transaction Coordinator交互获得Transaction ID对应的任务状态。Transaction Coordinator还负责将事务所有写入Kafka的一个内部Topic，这样即使整个服务重启，由于事务状态得到保存，进行中的事务状态可以得到恢复，从而继续进行。

### Consumer事务 
上述事务机制主要是从Producer方面考虑，对于Consumer而言，事务的保证就会相对较弱，尤其时无法保证Commit的信息被精确消费。这是由于Consumer可以通过offset访问任意信息，而且不同的Segment File生命周期不同，同一事务的消息可能会出现重启后被删除的情况。   

##  Kafka中有那些地方需要选举？这些地方的选举策略又有哪些？参考：
[https://blog.csdn.net/u013256816/article/details/89369160](https://blog.csdn.net/u013256816/article/details/89369160)
### 控制器的选举  
Kafka Controller的选举是依赖Zookeeper来实现的，在Kafka集群中哪个broker能够成功创建/controller这个临时（EPHEMERAL）节点他就可以成为Kafka Controller。

### 分区leader的选举
https://www.jianshu.com/p/1f02328a4f2e

### 消费者相关的选举
组协调器GroupCoordinator需要为消费组内的消费者选举出一个消费组的leader，这个选举的算法也很简单，分两种情况分析。如果消费组内还没有leader，那么第一个加入消费组的消费者即为消费组的leader。如果某一时刻leader消费者由于某些原因退出了消费组，那么会重新选举一个新的leader。

## Kafka中的延迟队列怎么实现？参考：
[https://blog.csdn.net/u013256816/article/details/80697456](https://blog.csdn.net/u013256816/article/details/80697456)

Kafka中存在大量的延迟操作，比如延迟生产、延迟拉取以及延迟删除等。Kafka并没有使用JDK自带的Timer或者DelayQueue来实现延迟的功能，而是基于时间轮自定义了一个用于实现延迟功能的定时器（SystemTimer）。JDK的Timer和DelayQueue插入和删除操作的平均时间复杂度为O(nlog(n))，并不能满足Kafka的高性能要求，而基于时间轮可以将插入和删除操作的时间复杂度都降为O(1)。Kafka中的时间轮（TimingWheel）是一个存储定时任务的环形队列，底层采用数组实现，数组中的每个元素可以存放一个定时任务列表（TimerTaskList）。TimerTaskList是一个环形的双向链表，链表中的每一项表示的都是定时任务项（TimerTaskEntry），其中封装了真正的定时任务TimerTask。时间轮由多个时间格组成，每个时间格代表当前时间轮的基本时间跨度（tickMs）。时间轮的时间格个数是固定的，可用wheelSize来表示，那么整个时间轮的总体时间跨度（interval）可以通过公式 tickMs × wheelSize计算得出。  

## kafka实现高吞吐
[https://blog.csdn.net/qq_39429714/article/details/84879543](https://blog.csdn.net/qq_39429714/article/details/84879543)
