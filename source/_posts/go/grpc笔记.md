---
categories: 
- go
tags:
- grpc
---

# grpc笔记

## 安装编译工具
### protoc 语言编译工具安装
1.下载protoc包 [https://github.com/protocolbuffers/protobuf/releases/tag/v3.11.4](https://github.com/protocolbuffers/protobuf/releases/tag/v3.11.4)  
2.解压
```
sudo unzip protoc-3.11.4-linux-x86_64.zip -d /usr/local/
```
### protoc-gen-go go语言插件安装
```
go get -u github.com/golang/protobuf/protoc-gen-go
```
<!--more-->


### 编译命令
Protobuf的protoc编译器是通过插件机制实现对不同语言的支持。比如protoc命令出现--xxx_out格式的参数，那么protoc将首先查询是否有内置的xxx插件，如果没有内置的xxx插件那么将继续查询当前系统中是否存在protoc-gen-xxx命名的可执行程序，最终通过查询到的插件生成代码。对于Go语言的protoc-gen-go插件来说，里面又实现了一层静态插件系统。比如protoc-gen-go内置了一个gRPC插件，用户可以通过--go_out=plugins=grpc参数来生成gRPC相关代码
```
protoc --go_out=plugins=grpc:. route_guide.proto
```


## 基本介绍
### 什么是grpc
在 gRPC 里客户端应用可以像调用本地对象一样直接调用另一台不同的机器上服务端应用的方法，使得您能够更容易地创建分布式应用和服务。与许多 RPC 系统类似，gRPC 也是基于以下理念：定义一个服务，指定其能够被远程调用的方法（包含参数和返回类型）。在服务端实现这个接口，并运行一个 gRPC 服务器来处理客户端调用。在客户端拥有一个存根能够像服务端一样的方法。
![](grpc笔记/grpc_concept_diagram_00.png)

### 使用 protocol buffers
gRPC 默认使用 protocol buffers，这是 Google 开源的一套成熟的结构数据序列化机制（当然也可以使用其他数据格式如 JSON）


### grpc服务定义
gRPC允许定义四类服务方法  
* 单项 RPC，即客户端发送一个请求给服务端，从服务端获取一个应答，就像一次普通的函数调用。
* 服务端流式 RPC，即客户端发送一个请求给服务端，可获取一个数据流用来读取一系列消息。客户端从返回的数据流里一直读取直到没有更多消息为止。
* 客户端流式 RPC，即客户端用提供的一个数据流写入并发送一系列消息给服务端。一旦客户端完成消息写入，就等待服务端读取这些消息并返回应答。
* 双向流式 RPC，即两边都可以分别通过一个读写数据流来发送一系列消息。这两个数据流操作是相互独立的，所以客户端和服务端能按其希望的任意顺序读写，例如：服务端可以在写应答前等待所有的客户端消息，或者它可以先读一个消息再写一个消息，或者是读写相结合的其他方式。每个数据流里消息的顺序会被保持。




## 安全认证
为了保障gRPC通信不被第三方监听篡改或伪造，我们可以对服务器启动TLS加密特性。  

### 用以下命令为服务器和客户端分别生成私钥和证书
```
$ openssl genrsa -out server.key 2048
$ openssl req -new -x509 -days 3650 \
    -subj "/C=GB/L=China/O=grpc-server/CN=server.grpc.io" \
    -key server.key -out server.crt


$ openssl genrsa -out client.key 2048
$ openssl req -new -x509 -days 3650 \
    -subj "/C=GB/L=China/O=grpc-client/CN=client.grpc.io" \
    -key client.key -out client.crt
```
以上命令将生成server.key、server.crt、client.key和client.crt四个文件。其中以.key为后缀名的是私钥文件，需要妥善保管。以.crt为后缀名是证书文件，也可以简单理解为公钥文件，并不需要秘密保存。在subj参数中的/CN=server.grpc.io表示服务器的名字为server.grpc.io，在验证服务器的证书时需要用到该信息。

有了证书之后，我们就可以在启动gRPC服务时传入证书选项参数：  
```
func main() {
    creds, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
    if err != nil {
        log.Fatal(err)
    }

    server := grpc.NewServer(grpc.Creds(creds))

    ...
}
```
其中credentials.NewServerTLSFromFile函数是从文件为服务器构造证书对象，然后通过grpc.Creds(creds)函数将证书包装为选项后作为参数传入grpc.NewServer函数。

客户端基于服务器的证书和服务器名字就可以对服务器进行验证：
```
func main() {
    creds, err := credentials.NewClientTLSFromFile(
        "server.crt", "server.grpc.io",
    )
    if err != nil {
        log.Fatal(err)
    }

    conn, err := grpc.Dial("localhost:5000",
        grpc.WithTransportCredentials(creds),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    ...
}
```
其中redentials.NewClientTLSFromFile是构造客户端用的证书对象，第一个参数是服务器的证书文件，第二个参数是签发证书的服务器的名字。然后通过grpc.WithTransportCredentials(creds)将证书对象转为参数选项传人grpc.Dial函数。


## 参考文献
* [grpc文档](http://doc.oschina.net/grpc?t=60133)
* [Go语言高级编程](https://chai2010.cn/advanced-go-programming-book/)