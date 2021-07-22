#!/usr/bin/env bash

hexo generate
hexo deploy
hexo clean


# 安装node环境
# sudo apt install npm

# 通过n模块安装指定的nodejs
# sudo npm install -g n


# node 和 hexo 存在版本适配问题
# 安装官方12.17 不翻墙会很慢
# sudo n 12.17


# 安装hexo
# sudo npm install -g hexo-cli --registry=https://registry.npm.taobao.org

# 拉取依赖
# npm install --force --registry=https://registry.npm.taobao.org
