# 项目依赖 
1. npm 安装
```shell
sudo apt install npm
```


## 安装指定版本node

### linux
1. 通过n模块安装指定的nodejs
```shell
sudo npm install -g n
```

2. node 和 hexo 存在版本适配问题
```shell
# 安装官方12.17 不翻墙会很慢
sudo n 12.17
```

### mac
1. 通过nvm安装指定版本nodejs
```shell
brew install nvm 
nvm install v12.17
nvm use v12.17
```

# hexo 安装
```shell

# 安装hexo
sudo npm install -g hexo-cli --registry=https://registry.npm.taobao.org

# 拉取依赖
npm install --force --registry=https://registry.npm.taobao.org
```


# hexo升级
```
//以下指令均在Hexo目录下操作，先定位到Hexo目录
//查看当前版本，判断是否需要升级
> hexo version

//全局升级hexo-cli
> npm i hexo-cli -g

//再次查看版本，看hexo-cli是否升级成功
> hexo version

//安装npm-check，若已安装可以跳过
> npm install -g npm-check

//检查系统插件是否需要升级
> npm-check

//安装npm-upgrade，若已安装可以跳过
> npm install -g npm-upgrade

//更新package.json
> npm-upgrade

//更新全局插件
> npm update -g

//更新系统插件
> npm update --save

//再次查看版本，判断是否升级成功
> hexo version
```