---
layout: post
title: mac下使用colima运行容器
categories: linux
tags: [docker, colima]
---
# colima 原理
## colima 原理

colima 通过以下方式在 macOS 上提供容器运行时：

1. 基于 Lima（Linux on macOS）启动一个轻量级 Linux 虚拟机  
2. 在该虚拟机内部运行 containerd 与 nerdctl，实现与 Docker CLI 兼容的接口  
3. 通过 Unix socket 将虚拟机内的 containerd 暴露给 macOS 主机，用户可像使用原生 Docker 一样执行 `docker`/`nerdctl` 命令  
4. 默认使用 QEMU 虚拟化，可选 VZ（Apple 原生虚拟化框架）以获得更好的性能与 Apple Silicon 兼容性  
5. 网络层通过 gvproxy 实现端口转发与 DNS 解析，使容器服务可直接通过 `localhost` 访问  
6. 卷挂载利用 virtiofs（或 sshfs）实现主机目录与虚拟机内目录的高效共享  
7. 整体架构无需安装 Docker Desktop，即可在 macOS 上获得接近原生的容器体验，并避免 Docker Desktop 的授权限制


# 安装&&配置
## 安装 colima
```bash
brew install colima docker docker-compose
```

## 配置 colima

启动 colima    
* 4核8G内存100G磁盘空间的虚拟机,
* 并配置docker registry,
```bash
colima start --edit 

# 修改配置
cpu: 4
memory: 8
disk: 100

# 修改docker配置，103.61.37.134是我搭的镜像代理
docker:
  registry-mirrors:
    - http://103.61.37.134:5000
  insecure-registries:
    - "103.61.37.134:5000"

```

运行nginx 测试
```bash
docker run -d -p 8080:80 nginx 
```

# 关于宿主机能够直接ping通容器的问题
orbstack 能够直接ping通容器的问题，而colima 不能,这是因为:  
Colima 使用的是传统的、嵌套的 NAT-on-NAT 模型。  
OrbStack 使用的是先进的、扁平的 无缝桥接 模型。   

## Colima 的网络模型：NAT-on-NAT
Colima 的核心是启动一个完整的 Linux 虚拟机（通过 QEMU），然后在这个 VM 内部运行 Docker。这就造成了双重网络隔离：

第一层 NAT（虚拟机层）：你的 Mac 宿主机和 Colima VM 处于不同的网络。Colima VM 有一个虚拟网卡（如 eth0），它从 Hypervisor（QEMU）获取一个私有 IP（比如 192.168.64.X）。QEMU 负责在宿主机和 VM 之间进行 NAT。这就是为什么你从 Mac 上无法直接 ping 通 VM 的 172.17.0.2，因为这个网段在 VM 内部，没有被映射到宿主机。
第二层 NAT（容器层）：在 Colima VM 内部，Docker 又创建了自己的私有网络（docker0 网桥，如 172.17.0.1/16）。Docker 引擎通过 iptables NAT规则 将容器端口映射到 VM 的 IP 地址上。   

## OrbStack 的网络模型：无缝桥接（Virtio-networking）
OrbStack 从头开始设计，其核心目标就是实现 macOS 和容器/虚拟机之间的高性能和无缝集成。它没有使用传统的 QEMU NAT 模型。

扁平化网络：OrbStack 使用自定义的 virtio-networking 驱动，在宿主机 macOS 和轻量级 Linux VM 之间建立一个共享的、扁平的虚拟网络。这意味着：

OrbStack 的 VM 和你的容器直接位于你 Mac 的局域网中。
它们可以从你的 Mac 获得一个真实的、可路由的 IP 地址（通常是 192.168.64.X 网段），而不是一个双层 NAT 后的地址。
网络栈共享：OrbStack 的 VM 和容器共享同一个网络栈。从网络的角度看，它们更像是一台设备上的多个进程，而不是嵌套的隔离环境。Docker 的 bridge 网络（172.17.0.0/16）被巧妙地桥接并暴露到了 macOS 的网络环境中。 

但orbstack是商业付费产品，而colima 是开源免费的。


# 搭建docker registry代理
搭建一个docker registry 代理，用于加速镜像拉取。


## 创建配置文件
创建一个名为 config.yml 的文件：
```yaml
version: 0.1
proxy:
  remoteurl: https://registry-1.docker.io # 上游官方仓库地址
  #username: [] # 可选，用于访问上游私有库
  #password: [] # 可选
storage:
  filesystem:
    rootdirectory: /var/lib/registry # 容器内缓存存储路径
http:
  addr: 0.0.0.0:5000 # 代理服务监听地址

```
## 启动代理仓库容器
```bash
# 创建本地目录存储数据 注意挂载到一块大容量的硬盘上
mkdir -p /opt/registry-data

# 运行容器
docker run -d \
  --name docker-registry-proxy \
  -p 5000:5000 \ # 将容器的5000端口映射到宿主机
  -v $(pwd)/config.yml:/etc/docker/registry/config.yml \ # 挂载配置文件
  -v /opt/registry-data:/var/lib/registry \ # 挂载数据目录
  --restart=always \
  registry:2


```



# 安装kubectl，配置本地连接容器化测试环境
```bash
brew install kubectl

# 将config配置为测试环境的config文件
vim ~/.kube/config
```


