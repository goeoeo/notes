---
title: go高级编程笔记

categories: 
- go
---


[go语言高级编程](https://chai2010.cn/advanced-go-programming-book/)


# 语言基础
## 函数、方法和接口
* init不是普通函数，可以定义有多个，所以也不能被其它函数调用
* 如果某个init函数内部用go关键字启动了新的goroutine的话，新的goroutine只有在进入main.main函数之后才可能被执行到。
* 每种类型对应的方法必须和类型的定义在同一个包中
<!--more-->


## 面向并发的内存模型
* go基于CSP(通信顺序进程)模型实现并发,并且Goroutine之间可共享内存
* Goroutine和系统线程不是等价的,一般系统线程占用一个固定大小的栈(2MB),Goroutine以一个大约2kb的栈启动,并且可以动态伸缩,最大可达1G的栈空间
* Goroutine采用的是半抢占式的协作调度,只有在当前Goroutine发生阻塞时才会发生调度
* 初始化顺序 同包内(const > var > init),如果一个包内有多个init函数,会根据文件名顺序进行调用
* 可通过带缓存的channel,控制Goroutine的并发数量
* 顺序一致性内存模型 Go语言中，同一个Goroutine线程内部，顺序一致性内存模型是得到保证的。但是不同的Goroutine之间，并不满足顺序一致性内存模型
  > 如果在一个Goroutine中顺序执行a = 1; b = 2;两个语句，虽然在当前的Goroutine中可以认为a = 1;语句先于b = 2;语句执行，但是在另一个Goroutine中b = 2;语句可能会先于a = 1;语句执行，甚至在另一个Goroutine中无法看到它们的变化（可能始终在寄存器中）

## 常见并发模式
不要通过共享内存来通信,而应通过通信来共享内存
* 等待N个线程完成后再进行下一步同步操作,可以使用sync.WaitGroup来实现
* 生产者消费者模型
* 发布订阅模型
* 控制并发数
* 赢者为王 通过适当开启一些冗余的线程，尝试用不同途径去解决同样的问题，最终以赢者为王的方式提升了程序的相应性能。
* 素数筛
* 并发的安全退出
* 通过Context包,实现Goroutine安全退出


# 第5章 Go和Web
## 哪些事情适合在中间件中做
```
compress.go
  => 对http的响应体进行压缩处理
heartbeat.go
  => 设置一个特殊的路由，例如/ping，/healthcheck，用来给负载均衡一类的前置服务进行探活
logger.go
  => 打印请求处理处理日志，例如请求处理时间，请求路由
profiler.go
  => 挂载pprof需要的路由，如`/pprof`、`/pprof/trace`到系统中
realip.go
  => 从请求头中读取X-Forwarded-For和X-Real-IP，将http.Request中的RemoteAddr修改为得到的RealIP
requestid.go
  => 为本次请求生成单独的requestid，可一路透传，用来生成分布式调用链路，也可用于在日志中串连单次请求的所有逻辑
timeout.go
  => 用context.Timeout设置超时时间，并将其通过http.Request一路透传下去
throttler.go
  => 通过定长大小的channel存储token，并通过这些token对接口进行限流
```

##  常见的流量限制手段
1.漏桶是指我们有一个一直装满了水的桶，每过固定的一段时间即向外漏一滴水。如果你接到了这滴水，那么你就可以继续服务请求，如果没有接到，那么就需要等待下一滴水。  

2.令牌桶则是指匀速向桶中添加令牌，服务请求时需要从桶中获取令牌，令牌的数目可以按照需要消耗的资源进行相应的调整。如果没有令牌，可以选择等待，或者放弃。  

实际应用中令牌桶应用较为广泛，开源界流行的限流器大多数都是基于令牌桶思想的。并且在此基础上进行了一定程度的扩充，比如github.com/juju/ratelimit  
