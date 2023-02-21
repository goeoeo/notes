---
categories: 
- go
tags:
- gomod
---

## 环境变量GO111MODULE
* GO111MODULE=off 
> go命令从不使用新模块支持。使用GOPATH模式(查找vendor目录和GOPATH路径下的依赖)

* GO111MODULE=on 
> go命令开启模块支持,此时执行go get 命令,所有依赖都会下到$GOPATH/pkg/mod ,go依赖查找也只会到$GOPATH/pkg/mod进行查找

* GO111MODULE=auto 
> 默认值,go命令根据当前目录启用或禁用模块支持。仅当当前目录位于$GOPATH/src之外并且其本身包含go.mod文件或位于包含go.mod文件的目录下时，才启用模块支持。简而言之,就是项目包含go.mod文件时会开启模块支持

<!--more-->
注意：

Go 1.13中，module mode优先级提升，GO111MODULE的默认值依然为auto，但在这个auto下，无论是在GOPATH/src下还是GOPATH之外的repo中，只要目录下有go.mod，go编译器都会使用go module来管理依赖。

## go mod命令
```
download    下载依赖的module到本地cache
edit        编辑go.mod文件
graph       打印模块依赖图
init        在当前文件夹下初始化一个新的module, 创建go.mod文件
tidy        增加丢失的module，去掉未用的module
vendor      将依赖复制到vendor下,注意依赖需要在import 中声明后才能进行导入
verify      校验依赖
why         解释为什么需要依赖
```
## go get ./... 
查找出当前项目的依赖，并进行下载

## 代理
GOPROXY 只有在GO111MODULE开启的时候生效，并且使用 go ./... 命令才能用到代理
* http://mirrors.aliyun.com/goproxy #阿里云代理
* https://goproxy.cn #七牛云代理

```
export GOPROXY=http://mirrors.aliyun.com/goproxy/
```

## 使用本地包
go.mod 文件中：
```
require modtest v0.0.0
replace modtest v0.0.0 => ../modtest
```

## 打包命令
1.使用GOPATH模式进行打包
```
export GO111MODULE=off
go build  -a -v -o app main.go
```
2.使用vendor目录下包来进行打包
```
export GO111MODULE=on
go build -mod=vendor -a -v -o app main.go
```



## go1.18 go workspaces
go work init  
会生成  go.work 文件    
```
go 1.18

use (
	../benchmark
)
```