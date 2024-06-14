# pyenv 多环境安装
macos 环境下 直接使用brew 安装 2.7 会提示找不到包    
virtualenv 是py 提供的一个隔离环境包的工具，不会管理多个解释器  
pyenv 可以用来管理多个解释器，以及利用virtualenv 来管理多个虚拟环境  

```
brew update
brew install pyenv
brew install pyenv-virtualenv

```

## 常用命令如下
1. 可安装版本列表 pyenv install -l
2. 安装python版本 pyenv install x.x.x
3. 创建虚拟环境【创建一个名字为test的虚拟环境，使用python版本为3.6.8】 pyenv virtualenv 3.6.8 test：
4. 激活虚拟环境【激活并使用名字为test的虚拟环境】 pyenv activate test
5. 退出虚拟环境 source deactivate
6. 删除虚拟环境【删除test虚拟环境】 pyenv uninstall test
7. 查看所有安装的虚拟环境 pyenv virtualenvs
8. 查看所有安装的python版本 pyenv versions
9. 显示当前版本 pyenv version
10. 设置Python解释器的默认版本 pyenv global 3.4.1

通过python install 2.7.13的时候在MAC M1下报错，换成安装2.7.18解决。