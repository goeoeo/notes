---
title: k8s入门文档
categories: 
- k8s
tags:
- k8s
---
## Kubernetes简介
Kubernetes（简称K8S，K和S之间有8个字母）是用于自动部署，扩展和管理容器化应用程序的开源系统。它将组成应用程序的容器组合成逻辑单元，
以便于管理和服务发现。Kubernetes 源自Google 15 年生产环境的运维经验，同时凝聚了社区的最佳创意和实践。  


## Kubernetes特性
* 服务发现与负载均衡:无需修改你的应用程序即可使用陌生的服务发现机制。
* 存储编排:自动挂载所选存储系统，包括本地存储。
* Secret和配置管理: 部署更新Secrets和应用程序的配置时不必重新构建容器镜像，且不必将软件堆栈配置中的秘密信息暴露出来。
* 批量执行: 除了服务之外，Kubernetes还可以管理你的批处理和CI工作负载，在期望时替换掉失效的容器。
* 水平扩缩:使用一个简单的命令、一个UI或基于CPU使用情况自动对应用程序进行扩缩。
* 自动化上线和回滚: Kubernetes会分步骤地将针对应用或其配置的更改上线，同时监视应用程序运行状况以确保你不会同时终止所有实例。
* 自动装箱:根据资源需求和其他约束自动放置容器，同时避免影响可用性。
* 自我修复:重新启动失败的容器，在节点死亡时替换并重新调度容器，杀死不响应用户定义的健康检查的容器。


## Minikube
Minikube是一种轻量级的Kubernetes实现，可在本地计算机上创建VM并部署仅包含一个节点的简单集群，Minikube可用于Linux、MacOS和Windows系统。
Minikube CLI提供了用于引导集群工作的多种操作，包括启动、停止、查看状态和删除。

## Kubernetes核心组件

### master组件
* apiserver 集群统一入口，以restful方式，交给etcd存储
* scheduler 节点调度，选择node节点部署应用
* controller-manager 处理集群中常规后台任务，一个资源对应一个控制器
* etcd 分布式存储数据库，保存集群相关数据

### Node
* kubelet master派到node节点代表，管理本机容器
* kube-proxy 提供网络代理、负载均衡等操作


## k8s核心概念

### Pod 
* k8s中最小的部署单元
* 一组容器的集合
* 一个Pod中的容器共享网络、存储
* 生命周期是短暂的

### controller
* 确保预期的pod的副本的数量
* 无状态的应用部署
* 有状态的应用部署
> 确保所有的node运行一个pod,一次性任务和定时任务



### service
* 定义一组pod的访问规则



## k8s 资源类型

### 名称空间级别的资源
* 工作负载型资源：Pod、ReplicaSet、Deployment、StatefulSet、DaemonSet、Job、CronJob
* 服务发现及负载均衡型资源：Service 、 Ingress
* 存储资源：Volume、CSI（容器存储接口，可以扩展各种各样的第三方存储卷）
* 特殊类型的存储卷：ConfigMap、Secret、DownwardAPI

### 集群级别
Namespace、Node、Role、ClusterRole、RoleBinding、ClusterRoleBinding

### 元数据类型
HPA、PodTemplate、LimitRange

## 必须存在的属性

|参数名|字段类型|说明|
|:---:|:---:|:---:|
|version|String|这里指k8s api的版本，目前基本上是v1,可以使用kubectl api-versions 命令查询|
|kind|String|这里指的是yaml文件定义的资源类型和角色，比如：Pod|
|metadata|Object|元数据对象，固定值就写metadata|
|metadata.name|string|元数据对象的名字，这里由我们编写，比如命名的Pod的名字|
|metadata.namespace|string|元数据对象的命名空间，由我们自身定义|
|Spec|Object|详细定义对象，固定值|
|spec.containers[]|list|这里是spec对象的容器列表定义，是个列表|
|spec.containers[].name|string|定义容器的名称|
|spec.containers[].image|string|定义要用到的镜像名称|
|spec.containers[].imagePullPolicy|string|定义镜像拉取策略，Always(总是拉取最新镜像),Never(拉取本地镜像，没有就不用),IfNotPresent(如果本地没有镜像则拉取)|
||||
||||
||||
||||
||||
||||
||||
||||
||||
||||
||||
||||
||||
||||





## 单master集群

### vmware 准备虚拟机
* Nat模式 修改子网Ip 192.168.137.0 ，nat网关 192.168.137.254
* 虚拟机内部配置ip https://blog.csdn.net/llluluyi/article/details/79041791 
>> /etc/sysconfig/network-scripts 
### 环境准备
* 一台或多台机器，操作系统 centos7
* 硬件配置，2g 2cpu 30g
* 集群之间所有机器网络互通
* 可以访问外网，需要拉取镜像
* 禁用swap分区


### 操作系统初始化工作
```
# 关闭防火墙
systemctl stop firewalld  # 临时处理
systemctl disable firewalld # 永久处理

# 关闭selinux
sed -i 's/enforcing/disabled/' /etc/selinux/config # 永久
setenforce 0 #临时

# 关闭swap 
swapoff -a # 临时
sed -ri 's/.*swap.*/#&/' /etc/fstab # 永久

# 根据规划设置主机名
hostnamectl set-hostname <hostname> 
例如：
hostnamectl set-hostname k8s-master
hostnamectl set-hostname node1

# 在master添加hosts
cat >> /etc/hosts <<EOF
192.168.137.100 k8s-master
192.168.137.101 k8s-node1
192.168.137.102 k8s-node2
192.168.137.103 k8s-node3
EOF

# 将桥接的IPv4 流量传递到iptables的链
cat > /etc/sysctl.d/k8s.conf << EOF
net.bridge.bridge-nf-call-ip6tables = 1 
net.bridge.bridge-nf-call-iptables = 1 
EOF

sysctl --system # 生效配置

# 时间同步
yum install ntpdate -y
ntpdate time.windows.com

```

### 所有节点安装docker/kubeadm/kubelet

#### 安装docker 
首先配置一下Docker的阿里yum源  
```
cat >/etc/yum.repos.d/docker.repo<<EOF
[docker-ce-edge]
name=Docker CE Edge - \$basearch
baseurl=https://mirrors.aliyun.com/docker-ce/linux/centos/7/\$basearch/edge
enabled=1
gpgcheck=1
gpgkey=https://mirrors.aliyun.com/docker-ce/linux/centos/gpg
EOF
```
yum方式安装docker  
```
# yum安装
yum -y install docker-ce

# 查看docker版本
docker --version  

# 启动docker
systemctl enable docker
systemctl start docker
```

```
cat >> /etc/docker/daemon.json << EOF
{
  "registry-mirrors": ["https://b9pmyelo.mirror.aliyuncs.com"]
}
EOF
```
#### 安装kubeadm，kubelet和kubectl

配置一下yum的k8s软件源  
```
cat > /etc/yum.repos.d/kubernetes.repo << EOF
[kubernetes]
name=Kubernetes
baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=0
repo_gpgcheck=0
gpgkey=https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg https://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
EOF
```
由于版本更新频繁，这里指定版本号部署： 
```
# 安装kubelet、kubeadm、kubectl，同时指定版本
yum install -y kubelet-1.18.0 kubeadm-1.18.0 kubectl-1.18.0
# 设置开机启动
systemctl enable kubelet
```

### 部署Kubernetes Master【master节点】
```
kubeadm init --apiserver-advertise-address=192.168.137.100 --image-repository registry.aliyuncs.com/google_containers --kubernetes-version v1.18.0 --service-cidr=10.96.0.0/12  --pod-network-cidr=10.244.0.0/16
```
使用kubectl工具 【master节点操作】    
```
  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

### 将节点加入到集群中
此命令由kubeadm init 时生产的
```
kubeadm join 192.168.137.100:6443 --token v17di1.mxjh3ryz2clh5mfk \
    --discovery-token-ca-cert-hash sha256:c33735da514c8d79e2af43b37a00bcea660fe13489566235c4a1c082e100a15b
```
### 部署CNI网络插件

```
# 添加
kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml

##①首先下载v0.13.1-rc2-amd64 镜像
##参考博客：https://www.cnblogs.com/pyxuexi/p/14288591.html
##② 导入镜像，命令，，特别提示，3个机器都需要导入，3个机器都需要导入，3个机器都需要导入，3个机器都需要导入，重要的事情说3遍。不然抱错。如果没有操作，报错后，需要删除节点，重置，在导入镜像，重新加入才行。本地就是这样操作成功的！
docker load < flanneld-v0.13.1-rc2-amd64.docker
#####下载本地，替换将image: quay.io/coreos/flannel:v0.13.1-rc2 替换为 image: quay.io/coreos/flannel:v0.13.1-rc2-amd64

# 查看状态 【kube-system是k8s中的最小单元】
kubectl get pods -n kube-system
```

raw.githubusercontent.com 如果被墙，自行下载kube-flannel.yml  
yum install  lrzsz -y # xshell 文件传输工具


### 测试k8s集群
在Kubernetes集群中创建一个pod，验证是否正常运行：  
```
# 下载nginx 【会联网拉取nginx镜像】
kubectl create deployment nginx --image=nginx
# 查看状态
kubectl get pod
```

下面我们就需要将端口暴露出去，让其它外界能够访问  
```
# 暴露端口
kubectl expose deployment nginx --port=80 --type=NodePort
# 查看一下对外的端口
kubectl get pod,svc
```

浏览器访问任意节点ip:port 
```
http://192.168.137.103:30980/
```


## K8S核心技术

### 集群命令行工具 kubectl 
#### kubectl 概述
kubectl 是k8s集群命令行工具，通过kubectl 能够对集群本身进行管理，并能够在集群上进行容器化应用的安装部署

#### 命令格式
```
kubectl [command] [type] [name] [flags]
```
参数：  
* command：指定要对资源执行的操作，例如create、get、describe、delete
* type：指定资源类型，资源类型是大小写敏感的，开发者能够以单数 、复数 和 缩略的形式
* name：指定资源的名称，名称也是大小写敏感的，如果省略名称，则会显示所有的资源
* flags：指定可选的参数，例如，可用 -s 或者 -server参数指定Kubernetes API server的地址和端口

#### 帮助命令
```
kubectl --help
```

### k8s集群yaml文件详解

#### 概述

k8s 集群中对资源管理和资源对象编排部署都可以通过声明样式（YAML）文件来解决，也就是可以把需要对资源对象操作编辑到YAML 格式文件中，
我们把这种文件叫做资源清单文件，通过kubectl 命令直接使用资源清单文件就可以实现对大量的资源对象进行编排部署了。一般在我们开发的时候，都是通过配置YAML文件来部署集群的。

YAML文件：就是资源清单文件，用于资源编排

#### YAML概述
YAML ：仍是一种标记语言。为了强调这种语言以数据做为中心，而不是以标记语言为重点。 

YAML 是一个可读性高，用来表达数据序列的格式。 

#### YAML 基本语法
* 使用空格做为缩进
* 缩进的空格数目不重要，只要相同层级的元素左侧对齐即可
* 低版本缩进时不允许使用Tab 键，只允许使用空格
* 使用#标识注释，从这个字符一直到行尾，都会被解释器忽略
* 使用 --- 表示新的yaml文件开始

#### 常用字段
|字段|含义|
|:---:|:---:|
|apiVersion|API版本|
|kind|资源类型|
|metadata|资源元数据|
|spec|资源规格|
|replicas|副本数量|
|selector|标签选择器|
|template|Pod模板|
|metadata|Pod元数据|
|spec|Pod规格|
|containers|容器配置|

#### 如何快速生成yaml文件
* 使用kubectl create命令生成yaml 文件
```
kubectl create deployment web --image=nginx -o yaml --dry-run=client > web.yaml
```
* 使用kubectl get 命令导出yaml文件
```
kubectl get deploy nginx -o=yaml > web1.yaml
```
### Pod


