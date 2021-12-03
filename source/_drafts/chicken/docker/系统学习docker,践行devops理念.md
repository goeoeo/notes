#系统学习docker,践行devops理念

## docker环境搭建
###1.安装docker
```bash
sudo apt install docker.io 

sudo docker version # 测试是否安装成功
```

###2.取消sudo，添加用户到docker用户组
```bash
sudo gpasswd -a ${USER} docker #添加当前用户到docker用户组
sudo restart #重启电脑生效
docker version # 测试取消sudo是否生效
```

###3.更换镜像

由于科学上网的原因，官方的docker镜像基本是用不了的。这里[参考阿里的docker镜像](https://help.aliyun.com/document_detail/60750.html)
```base
sudo vim /etc/docker/daemon.json #编辑docker配置文件 
# 内容为：
{
  "registry-mirrors": ["https://9lrfffi7.mirror.aliyuncs.com"]
}
sudo systemctl daemon-reload #重启docekr服务端，使得配置文件生效   
sudo systemctl restart docker #重启docker
docker info # 查看配置的镜像是否生效（Registry Mirrors）  
```
###4.安装docker-compose
```bash
sudo apt install docker-compose
```
## docker-machine 
[docker-machine介绍和使用](https://www.jianshu.com/p/cc3bb8797d3b)

docker-machine用于快速创建docker容器环境
Docker Machine常用于以下方面：
* 用于演示学习用途：例如在你的windows主机或者MAC、linux主机上，Docker Machine可以通过相关驱动（Hyper-V，VirtualBox），快速的创建多台docker虚拟机。
* 可以通过互联网云服务商的相关API接口，在Azure，AWS或Digital Ocean等云上创建你的docker云主机。

###1.安装docker-machine
[官方网站](https://docs.docker.com/machine/install-machine/)
```bash
base=https://github.com/docker/machine/releases/download/v0.16.0 &&
  curl -L $base/docker-machine-$(uname -s)-$(uname -m) >/tmp/docker-machine &&
  sudo mv /tmp/docker-machine /usr/local/bin/docker-machine &&
  chmod +x /usr/local/bin/docker-machine
```
