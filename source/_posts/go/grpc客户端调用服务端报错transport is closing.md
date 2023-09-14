---
title: grpc客户端调用服务端报错transport is closing
categories:
- go 
tags:
- transport is closing
---

# grpc客户端调用服务端报错transport is closing

为什么要吧这个错误单独写一篇文章，因为这个问题我debug了3天才解决。

## 首先说下现象
测试环境频繁报错，transport is closing，context canceled ，最先没有引起重视，直到导致了数据不一致才决定查这个这个问题。项目中只用了一元调用

## grpc client 和 server 配置
``` go
var BackOffConfig = backoff.Config{
	BaseDelay:  1 * time.Second,
	Multiplier: 1,
	Jitter:     0.2,
	MaxDelay:   60 * time.Second,
}

var KeepaliveClientConfig = keepalive.ClientParameters{
	Time:                10 * time.Second,
	Timeout:             5 * time.Minute,
	PermitWithoutStream: true,
}


var KeepaliveServerConfig = keepalive.ServerParameters{
	MaxConnectionIdle:     15 * time.Second, // If a client is idle for 10 seconds, send a GOAWAY
	MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
	MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
	Time:    5 * time.Second, // Ping the client if it is idle for 5 seconds to ensure the connection is still active
	Timeout: 5 * time.Minute, // Wait 1 second for the ping ack before assuming the connection is dead
}

var KeepaliveEnforcementPolicy = keepalive.EnforcementPolicy{
	MinTime:             5 * time.Second,
	PermitWithoutStream: true,
}

```
配置很平常，和grpc-go example相差不远  
项目中多个客户端共用了grpc conn 由于配置 keepalive 所以是长连接的方式  
握手过程:  
grpc client (一元调用)=> k8s service (由kube-proxy通过iptables/ipvs)负载均衡到 => pod 
> 这种调用方式会存在流量不均衡，后面再解决

[官方](https://github.com/grpc/grpc-go )解释的这4中情况都不是:
```

1. mis-configured transport credentials, connection failed on handshaking
2. bytes disrupted, possibly by a proxy in between
3. server shutdown
4. Keepalive parameters caused connection shutdown, for example if you have configured your server to terminate connections
regularly to trigger DNS lookups. If this is the case, you may want to increase your MaxConnectionAgeGrace, to allow longer RPC calls to finish.

```

## 可能的原因

### 并发？
项目中实际使用是，go 并发接受rabbitmq的消息，然后再用rpc客户端（共享的grpc conn连接）去调用rpc server ,我在本地测试的是否，测过并发的情况
并发量在5000的时候没有问题，想着测试环境没有这么高的并发，变没有注意。  


### grpc debug 
我开启了grpc 调试模式 
```
$ export GODEBUG=http2debug=2
$ export GRPC_GO_LOG_VERBOSITY_LEVEL=99
$ export GRPC_GO_LOG_SEVERITY_LEVEL=info
```
观察到每隔15秒，grpc server 会发送一个goaway的数据包到客户端，这是因为上面我们配置的MaxConnectionIdle 为15秒，通知客户端关闭连接。  
假如client已经发送请求到服务端，服务端请求耗时，在ping timeout 时间以内网络抖动（断网在联网），服务端再响应给客户端是没有问题的，
但超过了这个时间，client就会报transport is closing,server 也会报context canceled 。这导致我一直认为是网络抖动的问题。


### 网络抖动？
我将grpc client和grpc server 的ping timeout 调成了5分钟，此时还是有很多请求几秒钟就报transport is closing
因此排除了网络抖动这种情况。  


## 翻阅 https://github.com/grpc/grpc-go 的issues 
直到我看到了这个issues:https://github.com/grpc/grpc-go/issues/3297  
问题主描述了，他在单连接情况下遇到的the connection is draining和transport is closing  
有人回答了他: 
```
This is caused by a race between the start of a stream on a connection and the graceful close of the connection.
The server sends a GOAWAY when number of requests on the connection reaches the limit. But the client keeps sending new streams on this connection, until it receives the GOAWAY. The steams between server sending GOAWAY and client receiving GOAWAY will fail.

The gRPC client does a retry for this kind of failures, but retry only happens once (for each RPC).

How often do the errors happen? And which of the errors ("connection is draining" vs "transport is closing") happen more often?

One thing to try is to add grpc.WaitForReady(true) to the RPC.
Another thing to think of is to handle these errors in the application code. This is what will happen when the connection (TCP) is down for whatever reason. The application should be able to handle them (retry, or fail).
```
意思是由连接上的流开始和连接的正常关闭之间的竞争引起的。 当连接上的请求数达到限制时，服务器发送 GOAWAY。但是客户端不断在这个连接上发送新的流，直到它收到 GOAWAY。服务器发送 GOAWAY 和客户端接收 GOAWAY 之间的流将失败。  
也就是说如果server发出goaway的时候，此时grpc conn 正在发送流到服务端，这时候会产生一个竞争关系。  

抱着这个想法，我重新实验了我的并发程序，我将并发数量改到10000，transport is closing 出现了。平常的并发并不会触发服务器关闭连接，而是当服务器发送GOAWAY包的时候，客户端还在
并发的发送数据,就会导致服务端和客户端的流丢失，从而触发服务端关闭连接，报transport is closing。 

于是我将grpc server的配置改成了这样:  
```
var KeepaliveServerConfig = keepalive.ServerParameters{
	Time:    5 * time.Second, // Ping the client if it is idle for 5 seconds to ensure the connection is still active
	Timeout: 5 * time.Minute, // Wait 1 second for the ping ack before assuming the connection is dead
}
```
本身我们使用的长连接，在项目中不会频繁创建连接。所以这里不用关闭它，其次即使是断网的情况下，在ping timeout 时间内，网络恢复后连接依然可以重用。最坏的情况就是断网连不上了
这个连接就交个操作系统去回收吧（本身 tcp 也有keepalive机制）改完上到测试环境，世界安静了，再没有transport is closing

## 注意
1. 使用k8s service的负载均衡是4层负载均衡，只在握手时起作用，一旦长连接建立，service的负载便没有了作用。解决这个问题需要用grpc 客户端负载均衡进行处理  
2. 在helm upgrade的时候依然会报，这是因为在关闭grpc server 没有调用 GracefulStop() 方法进行平滑停止.  
3. GOAWAY 是http2协议包，并非grpc独有,http2实现了io多路复用，就是说可以并发的向同一个tcp连接中写数据，消息也不会错乱。  
4. http2 同一个连接，发送连接数是有限制的，在服务端可以设置（官方并不建议调这个参数），超过这个显示，后面的请求会排队。  


## 参考文章
* https://github.com/grpc/grpc-go/issues/3297
* [Golang信号处理和优雅退出守护进程](https://www.jianshu.com/p/ae72ad58ecb6)
* https://pandaychen.github.io/2020/09/01/GRPC-CLIENT-CONN-LASTING/
* [gRPC的平滑关闭和在Kubernetes上的服务摘流方案总结](https://cloud.tencent.com/developer/article/1816510)
* [GRPC 性能最佳做法](https://docs.microsoft.com/zh-cn/aspnet/core/grpc/performance?view=aspnetcore-5.0#reuse-grpc-channels)



