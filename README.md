# 项目依赖 
1. npm 安装
```shell
sudo apt install npm
```

2. 通过n模块安装指定的nodejs
```shell
sudo npm install -g n
```

3. node 和 hexo 存在版本适配问题
```shell
# 安装官方12.17 不翻墙会很慢
sudo n 12.17
```


# hexo 安装
```shell

# 安装hexo
sudo npm install -g hexo-cli --registry=https://registry.npm.taobao.org

# 拉取依赖
npm install --force --registry=https://registry.npm.taobao.org
```
