---
title: k8s踩坑日志
categories:
- k8s

tags:
- k8s
---

# pvc挂载不上pv
* annotations 不一致会影响pvc的挂载

# 通过coredns 访问mysql.beatflow-data.svc和mysql.beatflow-data.svc.cluster.local ，有些时候在不同命名空间下可能访问不到

# coredns debug
```
kubectl run dig --rm -it --image=docker.io/azukiapp/dig /bin/sh
/ # nslookup kubernetes.default
Server:		10.96.0.10
Address:	10.96.0.10#53

Name:	kubernetes.default.svc.cluster.local
Address: 10.96.0.1
```

# 在kubeSphere 管理的k8s集群上搭建使用Prometheus Operator helm安装prometheus 起不来
* kubeSphere 已经安装了prometheus ,并且每个节点已经装了node-exporter所以导致起不来
* 目前我们只需要安装grafana就可以了
* Prometheus Operator 通过service monitor 添加监控节点。

# k8s服务中headless 和 Cluster IP 区别，使用headless会导致pod访问延迟
headless 模式通过CoreDNS 解析到pod上面不走service的负载均衡  
因为通过coredns ，pod起来后，并不能立即访问到会有延迟  
clusterIp 模式，访问service 的clusterIp 再通过iptables 转发到pod上面，pod起来后可以直接访问，不会有延迟。  

# minikube 在使用ntf 报错  does not support NFS export 
挂载 /nfs 报错 does not support NFS export  
解决: 挂载 /data/nfs
minikube 在driver 为docker的模式下面 只能挂载到/data目录下面去，其他目录都会报错


# helm 安装报错
```
Error: Kubernetes cluster unreachable: Get "http://localhost:8080/version?timeout=32s": dial tcp 127.0.0.1:8080: connect: connection refused
```
报错原因: helm v3版本不再需要Tiller，而是直接访问ApiServer来与k8s交互，通过环境变量KUBECONFIG来读取存有ApiServre的地址与token的配置文件地址，默认地址为~/.kube/config  
export KUBECONFIG=~/.kube/config  


# minikube 服务端口映射到主机
场景: 需要外网访问minikube集群中的服务  
```
kubectl port-forward svc/mysql 30001:3306 -n beatflow-data --address 0.0.0.0
```


# OCI runtime exec failed: exec failed: unable to start container process: open /dev/pts/0: operation not permitted: unknown
调整contained的配置后，调整之前的pod进不去，需要重启pod