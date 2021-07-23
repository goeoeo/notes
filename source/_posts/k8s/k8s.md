---
categories:
- k8s
  tags:
- harbor
---

# k8s 备忘录

<!--more-->

## k8s 中Service Account 
k8s创建两套独立的账号系统，原因如下：  
1. User账号给用户用，Service Account是给Pod里的进程使用的，面向的对象不同  
2. User账号是全局性的，Service Account则属于某个具体的Namespace  
3. User账号是与后端的用户数据库同步的，创建一个新用户通常要走一套复杂的业务流程才能实现，Service Account的创建则需要极轻量级的实现方式，
   集群管理员可以很容易地为某些特定任务创建一个Service Account  
4. 对于一个复杂的系统来说，多个组件通常拥有各种账号的配置信息，Service Account是Namespace隔离的，可以针对组件进行一对一的定义，同时具备很好的“便携性”


Controller Manager创建了ServiceAccount Controller和Token Controller这两个安全相关的控制器。其中ServiceAccount Controller一直
监听Service Account和Namespace的事件，如果在一个Namespace中没有default Service Account，那么Service Account会给Namespace创建一个默认（default）的Service Account

默认的service account 仅能获取当前Pod自身的相关属性，无法观察到其他名称空间Pod的相关属性信息，如果想要扩展pod,假设有一个Pod 需要用于管理
其他Pod或者其他资源对象，是无法通过吱声的名称空间的serviceAccount 进行其他Pod相关属性信息的获取的，此时就需要进行手动创建一个serviceAccount，并在创建Pod时进行定义。实际上serviceAccount也属于k8s个资源

