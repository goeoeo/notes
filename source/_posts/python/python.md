
# 使用国内镜像 

## 临时使用
```
pip3.8 install -i https://pypi.tuna.tsinghua.edu.cn/simple some-package
```
## 设为默认
```
pip3.8 config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple
```

# 升级pip 
```
/usr/bin/python3.8 -m pip install --upgrade pip

# 默认安装到 ~/.local/bin 下面， 需要将此目录加入PATH中
```