# pyenv mac安装
macos 环境下 直接使用brew 安装 2.7 会提示找不到包    
virtualenv 是py 提供的一个隔离环境包的工具，不会管理多个解释器  
pyenv 可以用来管理多个解释器，以及利用virtualenv 来管理多个虚拟环境  

```
brew update
brew install pyenv
brew install pyenv-virtualenv

```

# pyenv ubuntu安装
## 安装依赖
```
sudo apt install -y make build-essential libssl-dev zlib1g-dev libbz2-dev libreadline-dev libsqlite3-dev wget curl llvm libncurses-dev xz-utils tk-dev libffi-dev liblzma-dev  git
```
## 安装pyenv
```
curl https://pyenv.run | bash
```

## 更改环境变量
```
vim ~/.bashrc

export PATH="$HOME/.pyenv/bin:$PATH"
eval "$(pyenv init --path)"
eval "$(pyenv init -)"
eval "$(pyenv virtualenv-init -)"

source ~/.bashrc
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



# python模块路径搜索
Python解释器会搜索当前目录、所有已安装的内置模块和第三方模块(site-packages)，搜索路径存放在sys模块的path变量中

添加自定义搜索目录，两种方法：
1. 直接修改sys.path （临时生效，运行结束后失效）
    ```shell
    >>> import sys
    >>> sys.path.append('/Users/michael/my_py_scripts')
    ```
2. 设置环境变量PYTHONPATH，该环境变量的内容会被自动添加到模块搜索路径中。设置方式与设置Path环境变量类似。注意只需要添加你自己的搜索路径，Python自己本身的搜索路径不受影响。


# pip install 加速
```
pip install package_name -i https://pypi.tuna.tsinghua.edu.cn/simple
```