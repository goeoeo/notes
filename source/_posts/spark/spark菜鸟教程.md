# 基础
## Spark解决的问题
海量数据的计算，可以进行离线以及流处理

## Spark模块
* SparkCore
* SparkSQL 
* SparkStreaming 流计算
* Graphx 图计算
* MLlib 机器学习

## Spark 特点
速度快，使用简单，通用性强，多种运行模式

## spark运行模式
* 本地模式 一个独立的进程，通过内部多线程来模拟整个spark运行时环境
* Standalone(集群) Spark中的各个角色以独立进程的形式存在，并组成Spark集群环境
* Hadoop YARN(集群) Spark各个角色运行在YARN容器内部，组成Spark集群环境
* K8s模式 各个角色运行在k8s容器中，组成spark集群环境
* 云服务模式

## spark架构
* Master,管理整个集群资源
* Worker,管理单个服务器的资源
* Dirver,管理单个Spark在运行的时候的工作
* Executor,工作者

