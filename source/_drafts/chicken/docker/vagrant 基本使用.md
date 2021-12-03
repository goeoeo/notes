### 1.介绍
​　　Vagrant是一个基于Ruby的工具，用于创建和部署虚拟化开发环境。它 使用Oracle的开源[VirtualBox](https://baike.baidu.com/item/VirtualBox)虚拟化系统，使用 Chef创建自动化虚拟环境。我们可以使用它来干如下这些事：

*   建立和删除虚拟机
*   配置虚拟机运行参数
*   管理虚拟机运行状态
*   自动配置和安装开发环境
*   打包和分发虚拟机运行环境

​　　Vagrant的运行，需要**依赖**某项具体的**虚拟化技术**，最常见的有VirtualBox以及VMWare两款，早期，Vagrant只支持VirtualBox，后来才加入了VMWare的支持。

### 2.安装VirtualBox和Vagrant
#### 1.安装VirtualBox
下载地址：[https://www.virtualbox.org/wiki/Linux_Downloads](https://www.virtualbox.org/wiki/Linux_Downloads)
```bash
sudo dpkg -i virtualbox-6.0_6.0.4-128413_Ubuntu_bionic_amd64.deb
```
#### 2.安装 vagrant
下载地址：[https://www.vagrantup.com/downloads.html](https://www.vagrantup.com/downloads.html)
```bash
sudo dpkg -i vagrant_2.2.4_x86_64.deb
```
#### 3.vagrant box基本命令
* 列出本地环境中所有的box
```bash
vagrant box list
```
* 添加box到本地vagrant环境
```bash
vagrant box add box-name(box-url)
```
* 更新本地环境中指定的box
```bash
vagrant box update box-name
```
* 删除本地环境中指定的box
```bash
vagrant box remove box-name
```
* 重新打包本地环境中指定的box
```bash
vagrant box repackage box-name
```

* 在线查找需要的box
官方网址：[https://app.vagrantup.com/boxes/search](https://app.vagrantup.com/boxes/search)

#### 4. vagrant基本命令
* 在空文件夹初始化虚拟机
```bash
vagrant init [box-name]
```
* 在初始化完的文件夹内启动虚拟机
```bash
vagrant up
```
* ssh登录启动的虚拟机
```bash
vagrant ssh
```
* 挂起启动的虚拟机
```bash
vagrant suspend
```

* 重启虚拟机
```bash
vagrant reload
```
* 关闭虚拟机
```bash
vagrant halt
```
* 查找虚拟机的运行状态
```bash
vagrant status
```
* 销毁当前虚拟机
```bash
vagrant destroy
```
