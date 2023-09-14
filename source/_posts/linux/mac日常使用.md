---
title: mac日常使用
categories:
- linux

tags:
- mac
---
# docker安装
docker-desktop 太消耗资源且无法商用，这里使用 colima +dockercli 的方案   

1. 安装docker client   
```shell
brew install docker docker-compose
```
2. 安装colima (很漫长)  
```shell
brew install colima
```

3. 将docker-compose 作为docker的插件  
```shell
$ mkdir -p ~/.docker/cli-plugins
$ ln -sfn $(brew --prefix)/opt/docker-compose/bin/docker-compose ~/.docker/cli-plugins/docker-compose
```
> docker info 可以看到 不再报找不到插件的错误

4. 启动colima 同时配置镜像
```shell
# 找到docker 配置  "registry-mirrors": ["https://9lrfffi7.mirror.aliyuncs.com"]
colima start -e 
```




# mac宿主机和docker容器网络不通 (测试无效)
安装docker-connector服务  
1. 使用brew安装docker-connector  
```shell
brew install wenjunxiao/brew/docker-connector
```
2. 执行下面命令将docker所有 bridge 网络都添加到docker-connector路由
```shell
docker network ls --filter driver=bridge --format "{{.ID}}" | xargs docker network inspect --format "route {{range .IPAM.Config}}{{.Subnet}}{{end}}" >> "$(brew --prefix)/etc/docker-connector.conf"
```
> /usr/local/etc/docker-connector.conf是安装docker-connector后生成的配置文件

3. 使用sudo启动docker-connector服务  
```shell
sudo brew services start docker-connector
```

4. 使用下面命令创建wenjunxiao/mac-docker-connector容器，要求使用 host 网络并且允许 NET_ADMIN
```shell
docker run -it -d --restart always --net host --cap-add NET_ADMIN --name connector wenjunxiao/mac-docker-connector
```
注意服务启动后宿主机和容器网络还是没有生效 执行下面这个，可以不以守护进程的模式启动
```shell
sudo docker-connector -config /opt/homebrew/etc/docker-connector.conf
```
docker-connector容器启动成功后，macOS宿主机即可访问其它容器网络


# 解决brew install 缓慢问题
## 首次安装
```shell
curl https://raw.githubusercontent.com/Homebrew/install/master/install.sh  > install-brew.sh
```
然后,在下载的文件中, 修改BREW_REPO为:  
```shell
BREW_REPO="https://mirrors.ustc.edu.cn/brew.git"
```

最后, 运行:  
```shell
HOMEBREW_CORE_GIT_REMOTE=https://mirrors.ustc.edu.cn/homebrew-core.git bash install-brew.sh
```

## 已经安装  
```shell
cd "$(brew --repo)"
git remote set-url origin https://mirrors.ustc.edu.cn/brew.git

echo 'export HOMEBREW_BOTTLE_DOMAIN=https://mirrors.ustc.edu.cn/homebrew-bottles' >> ~/.bash_profile
source ~/.bash_profile

cd "$(brew --repo)/Library/Taps/homebrew/homebrew-core"
git remote set-url origin https://mirrors.ustc.edu.cn/homebrew-core.git
```


# 命令自动补全 
```shell
git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
```
vim .zshrc 写入
```
plugins=(git zsh-autosuggestions)
```

## git命令 自动补全
```shell
brew install bash-completion
mkdir .zsh_fpath
curl https://raw.githubusercontent.com/git/git/master/contrib/completion/git-completion.zsh \
-o ~/.zsh_fpath/.git-completion.zsh
```
vim .zshrc 内容如下  
```
zstyle ':completion:*:*:git:*' script ~/.zsh_fpath/.git-completion.zsh
fpath=(~/.zsh_fpath $fpath)
autoload -Uz compinit && compinit
```


# 安装ab压测工具
## 先安装三个工具 
```shell
brew install apr
brew install apr-util
brew install prce
```

> 注意根据输出更新 ~/.zshrc 以及对应环境变量

## 下载httpd
https://httpd.apache.org/download.cgi#apache24  
解压进行 httpd目录   
```
./configure
make  
sudo make install 

# 测试
ab -v
```

# mac 添加 ssh-agent
现象：当我用golang代码通过sshtunnel 通过跳板机连接数据库的时候，报错 认证失败  
原因：我本地mac 机器没有添加ssh client    
解决： 
```
eval 'ssh-agent'
ssh-add ~/.ssh/id_rsa
```
