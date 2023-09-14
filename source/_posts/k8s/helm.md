---
title: helm
categories: 
- k8s
tags:
- k8s
---

# Helm
Helm 是 Kubernetes 的包管理器。包管理器类似于我们在 Ubuntu 中使用的apt、Centos中使用的yum 或者Python中的 pip 一样，能快速查找、下载和安装软件包  
官网:https://helm.sh/  
官方文档地址: https://helm.sh/zh/docs


<!--more-->

## Helm介绍

### 基本方式部署应用流程
1. 编写或者导出资源编排文件yaml  
2. deployment部署pod  
3. service暴露服务  
4. Ingress服务实现域名到Service的路由  

以上的方式存在的缺点:  
如果我们部署的是微服务项目，可能是10多20个或者更多的服务，那么每套服务都需要一套yaml文件  
这样下来，我们需要维护的yaml文件就是一大堆,而且如果我们升个级，那就非常不方便了  

### Helm作用
* 可以使用Helm把上面我们说到的微服务部署yaml作为一个整体管理
* 还可以实现yaml的高效复用
* 使用Helm应用级别的版本管理，对于版本升级和回滚操作相当简单了起来


### Helm的三个重要概念
* helm helm命令行管理工具，类似于kubectl
* Chart 打包的yaml集合
* Release 基于 Chart 的部署实体 ，一个 chart 被 helm 运行后将会生成对应的一个Release，相当于就是针对Chart的版本管理


## Helm安装

### helm安装
https://helm.sh/zh/docs/intro/install/  




























