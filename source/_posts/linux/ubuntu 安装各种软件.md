---
categories: 
- linux
tags:
- linux软件安装
---

### chrome

#### deb包
http://www.ubuntuchrome.com/

#### 脚本安装
```bash
sudo wget http://www.linuxidc.com/files/repo/google-chrome.list -P /etc/apt/sources.list.d/
wget -q -O - https://dl.google.com/linux/linux_signing_key.pub  | sudo apt-key add -
sudo apt-get update
sudo apt-get install -y google-chrome-stable
```
<!--more-->

### yarn
```bash
curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -
echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list
sudo apt-get update && sudo apt-get install yarn
```
### npm
```bash
sudo apt install npm

sudo npm install -g n #通过n模块安装指定的nodejs
sudo n stable #安装官方稳定版本

```

### pavucontrol(解决机箱前置耳机没声音)
```bash
 sudo apt-get install pavucontrol
# 选择配置-》模拟立体声双工
# 输出设备-》模拟耳机
```

### composer 
```bash
sudo apt-get install composer 

composer config -g repo.packagist composer https://packagist.laravel-china.org #laravel-china 社区镜像
```

### lnmp
```bash
sudo apt-get install libcurl4-gnutls-dev #ubuntu18.04下先安这个curl扩展,不然php 中curl 请求https会报502

# https://lnmp.org/auto.html 这里可以生成无人值守安装命令
wget http://soft.vpser.net/lnmp/lnmp1.5.tar.gz -cO lnmp1.5.tar.gz && tar zxf lnmp1.5.tar.gz && cd lnmp1.5 && LNMP_Auto="y" DBSelect="4" DB_Root_Password="root" InstallInnodb="y" PHPSelect="8" SelectMalloc="1" ./install.sh lnmp
```
### xx_net
```bash
cd /usr/local && https://github.com/XX-net/XX-Net.git
sudo apt-get install miredo #开启ipv6
```
> [需要的crx](https://crxdl.com/)
#### ubuntu终端使用xx_net
```
sudo  apt install privoxy

#XX-Net默认的端口是8087，对https代理  ,XX-Net默认的端口是8087，对http代理
export https_proxy=127.0.0.1:8087
export http_proxy=127.0.0.1:8087 

# 测试
curl https://www.google.com
```


### Tim
```bash
cd /usr/local/ && sudo git clone https://gitee.com/wszqkzqk/deepin-wine-for-ubuntu.git && cd deepin-wine-for-ubuntu && ./install.sh
sudo apt-get install -f # 如果提示依赖,执行这个
wget https://mirrors.aliyun.com/deepin/pool/non-free/d/deepin.com.qq.office/deepin.com.qq.office_2.0.0deepin4_i386.deb   #tim 
```

### supervisor(进程管理工具)  
Supervisor是用Python开发的一套通用的进程管理程序，能将一个普通的命令行进程变为后台daemon，并监控进程状态，异常退出时能自动重启。
	
```bash
sudo apt-get install supervisor
```

### java 

```bash
java -version 	#输出
Command 'java' not found, but can be installed with:
sudo apt install default-jre            
sudo apt install openjdk-11-jre-headless
sudo apt install openjdk-8-jre-headless 

sudo apt install default-jre  #安装默认版本的java,ubuntu 默认安装的是java10.0.2
```

### jmeter (apache压测工具)
```bash
sudo apt install jmeter
```
### 系统托盘
安装 Gnome Shell 插件：TopIcons Plus

### golang 环境安装
1.到[https://golang.org/dl/](https://golang.org/dl/)下载最新的golang包，需要翻墙
2. 国内下载地址[https://studygolang.com/dl](https://studygolang.com/dl)

2.将包解压到/usr/local下
```bash
sudo tar -C /usr/local -xzf go1.11.4.linux-amd64.tar.gz
```
3.将go执行命令放到环境变量中
修改系统环境变量 /etc/environment  
PATH=$PATH:/usr/local/go/bin:/home/yu/go/bin  
通过go 命令更改相关设置
```bash
go env -w GOROOT=/usr/local/go 
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn
```


4.go version 测试是否安装成功

### 安装多版本go 
```shell
go install golang.org/dl/go1.19.3@latest
go1.19.3 download
go1.19.3 version
```

### protoc 安装
1. 下载protoc包
```
https://github.com/protocolbuffers/protobuf/releases/tag/v3.17.1
```
2. 解压
```
sudo unzip protoc-3.11.4-linux-x86_64.zip -d /usr/local/
```
3.  protoc-gen-go 编译插件
```
 go get -u github.com/golang/protobuf/protoc-gen-go
```

> protoc --go_out=plugins=grpc:. route_guide.proto


### Wireshark 抓包工具
```
sudo add-apt-repository ppa:wireshark-dev/stable 
sudo apt update
sudo apt -y install wireshark
```

### Docker

#### docker
```
sudo apt install docker.io
sudo gpasswd -a $USER docker 
newgrp docker
```
#### 更换docker远程镜像
```bash
cat >> /etc/docker/daemon.json <<EOF

{
  "registry-mirrors": ["https://9lrfffi7.mirror.aliyuncs.com"]
}
EOF
sudo systemctl daemon-reload &&  sudo systemctl restart docker
docker info
```

#### docker-compose
```
# 方式1
sudo curl -L --fail https://github.com/docker/compose/releases/download/1.29.2/run.sh -o /usr/local/bin/docker-compose

# 方式1不好使，直接去https://github.com/docker/compose/releases/上下载 
mv docker-compose-Linux-x86_64   /usr/local/bin/docker-compose

sudo chmod +x /usr/local/bin/docker-compose
docker-compose --version
```

#### docker-machine
[官方网站](https://docs.docker.com/machine/install-machine/)
```
base=https://github.com/docker/machine/releases/download/v0.16.0 &&
  curl -L $base/docker-machine-$(uname -s)-$(uname -m) >/tmp/docker-machine &&
  sudo mv /tmp/docker-machine /usr/local/bin/docker-machine &&
  chmod +x /usr/local/bin/docker-machine

```

#### mysql-workbench
```
# ubuntu 20.04以下
sudo apt-get install mysql-workbench

# ubuntu 20.04以上
https://linuxhint.com/installing_mysql_workbench_ubuntu/


sudo apt install ./mysql-apt-config_0.8.15-1_all.deb
ok
ok
sudo apt update
sudo apt install mysql-workbench-community

```


## 安装nps 内网穿透
1. 安装内网穿透服务端
```
mkdir /usr/local/nps && /usr/local/nps  && wget https://github.com/ehang-io/nps/releases/download/v0.26.8/linux_amd64_server.tar.gz
tar -zxf linux_amd64_server.tar.gz
```

## helm 
```
https://github.com/helm/helm/releases 

tar -zxvf helm-v3.0.0-linux-amd64.tar.gz
mv linux-amd64/helm /usr/local/bin/helm

```

## easyconnect 
Failed to load module "canberra-gtk-module"  
```
sudo apt-get install libcanberra-gtk-module
```
提示Pango-ERROR **: 10:24:00.000: Harfbuzz version too old (1.3.1)  
将data/easyconnect/lib.zip 拷贝到 /usr/share/sangfor/EasyConnect 即可  


## ubuntu 安装markdown编辑器
地址: https://typora.io/#linux  

```
# or run:
# sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys BA300B7755AFCFAE
wget -qO - https://typora.io/linux/public-key.asc | sudo apt-key add -
# add Typora's repository
sudo add-apt-repository 'deb https://typora.io/linux ./'
sudo apt-get update
# install typora
sudo apt-get install typora

```

## ubuntu 安装DBEaver
https://dbeaver.io/download/

## 安装企业微信
### 安装deepin-wine
https://gitee.com/wszqkzqk/deepin-wine-for-ubuntu
https://github.com/zq1997/deepin-wine

### 安装企业微信
https://github.com/zq1997/deepin-wine


```shell
wget -O- https://deepin-wine.i-m.dev/setup.sh | sh
sudo apt install com.qq.weixin.work.deepin
# 注释 /opt/apps/com.qq.weixin.work.deepin/files/run.sh中的第24和26行代码
```

#### 企业微信无法显示登录二维码
```
# 执行下面命令，重新打开企业微信就有二维码了，实测
$ sudo apt install libjpeg62:i386 fonts-wqy-zenhei xfonts-wqy
```

## ubuntu22.04 安装sogou
https://blog.csdn.net/weixin_64684404/article/details/125461860

安装依赖
```
sudo apt install libqt5qml5 libqt5quick5 libqt5quickwidgets5 qml-module-qtquick2

sudo apt install libgsettings-qt1
```

## 终端管理工具electerm
https://github.com/electerm/electerm

```shell
sudo snap install electerm --classic
```

## ipfs
https://github.com/ipfs/ipfs-desktop/releases/download/v0.23.0/ipfs-desktop-0.23.0-linux-amd64.deb