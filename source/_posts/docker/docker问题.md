---
title: docker问题

categories: 
- docker

tags:
- docker-install-ware
---

# docker本地空间足够，但容器日志提示空间不足

## 情况描述
项目中，es报空间不足，es容器 已使用30G左右的空间
默认情况下docker 对每个容器使用空间有限制  
```
Elasticsearch：high disk watermark [90%] exceeded
```
<!--more-->

## 处理
修改配置   
vim /etc/docker/daemon.json   
```
{
    "storage-opt": [ "dm.basesize=80G" ]
} 
```
重启 
```
sudo systemctl daemon-reload &&  sudo systemctl restart docker
```

# docker 清理磁盘
```shell
docker system df -v # -v 可以显示详细信息
```
该命令列出了 docker 使用磁盘的 4 种类型：  
* Images: 所有镜像占用的空间，包括拉取的镜像、本地构建的镜像
* Containers: 运行中的容器所占用的空间（没运行就不占空间），其实就是每个容器读写层的空间
* Local Volumes: 本地数据卷的空间
* Build Cache: 镜像构建过程中，产生的缓存数据

```shell
docker system prune -f
## 可以配置定时任务每天凌晨1点清理
0 0 1 * * docker system prune -f
```
该命令会删除暂停中的容器、没有关联容器的镜像、没有 tag 的镜像、没有被使用的数据卷，简单而言，没有在 run 或被使用的东西都被清理掉，
注意，如果你有一些暂时暂停的容器，这个命令也会将其清理掉。


# docker镜像操作
```
# 打包镜像
docker save xx.tar $imageName

# 载入镜像
docker load - xx.tar

# docker打包镜像
docker buildx build --platform linux/amd64,linux/arm64 -t martindai/wechat-robot:1.0 --load .
```


# docker buildx 使用

## 安装buildx
```
# 下载
 wget https://github.com/docker/buildx/releases/download/v0.8.2/buildx-v0.8.2.linux-amd64
 
# copy
mv buildx-v0.8.2.linux-amd64 ~/.docker/cli-plugins/docker-buildx
chmod +x  ~/.docker/cli-plugins/docker-buildx

# 查看版本
docker buildx 
```

## 安装模拟器
安装模拟器的主要作用是让 buildx 支持跨 CPU 架构编译。
```
docker run --privileged --rm tonistiigi/binfmt --install all

docker buildx ls
```
## 创建构建器
docker buildx create --name multi-platform --use --platform linux/amd64,linux/arm64 --driver docker-container

## 启动构建器
docker buildx inspect multi-platform  --bootstrap

## 编译
docker buildx build --push --platform linux/amd64,linux/arm64/v8 -t dockerhub.betacloud.com/service/pitrix-flyway:v0.0.2  .


# docker 报错 x509: certificate signed by unknown authority
1. 打开daemon.json，加入你的私库地址 vi /etc/docker/daemon.json
```
{
"insecure-registries":["私库地址"]
}
```
2. 重启dokcer
```
systemctl daemon-reload && systemctl restart docker
```
# 参考
* [收藏！24 个 Docker 疑难杂症处理技巧](https://www.bilibili.com/read/cv16472262)




# docker restrt 失败卡住一直处于activating状态
```
# 关键命令
systemctl daemon-reexec

systemctl daemon-reload

systemctl restart docker
```