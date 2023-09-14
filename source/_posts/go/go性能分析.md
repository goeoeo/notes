---
title: go性能分析
categories:
- go

tags:
- 性能分析
---


在 Go 语言中，PProf 是用于可视化和分析性能分析数据的工具，PProf 以 profile.proto 读取分析样本的集合，并生成报告以可视化并帮助分析数据（支持文本和图形报告）。

<!--more-->


# 有哪几种采样方式
* runtime/pprof：采集程序（非 Server）的指定区块的运行数据进行分析。
* net/http/pprof：基于 HTTP Server 运行，并且可以采集运行时数据进行分析。
* go test：通过运行测试用例，并指定所需标识来进行采集。
* gops: 针对非HTTP Server的其它Server 比如GRPC Server 持续采集


# 支持什么使用模式
* Report generation：报告生成。
* Interactive terminal use：交互式终端使用。
* Web interface：Web 界面。

# 可以做什么
* CPU Profiling：CPU 分析，按照一定的频率采集所监听的应用程序 CPU（含寄存器）的使用情况，可确定应用程序在主动消耗 CPU 周期时花费时间的位置。
* Memory Profiling：内存分析，在应用程序进行堆分配时记录堆栈跟踪，用于监视当前和历史内存使用情况，以及检查内存泄漏。
* Block Profiling：阻塞分析，记录 Goroutine 阻塞等待同步（包括定时器通道）的位置，默认不开启，需要调用 runtime.SetBlockProfileRate 进行设置。
* Mutex Profiling：互斥锁分析，报告互斥锁的竞争情况，默认不开启，需要调用 runtime.SetMutexProfileFraction 进行设置。
* Goroutine Profiling： Goroutine 分析，可以对当前应用程序正在运行的 Goroutine 进行堆栈跟踪和分析。这项功能在实际排查中会经常用到，
因为很多问题出现时的表象就是 Goroutine 暴增，而这时候我们要做的事情之一就是查看应用程序中的 Goroutine 正在做什么事情，因为什么阻塞了，然后再进行下一步。



# 服务型持续采集

## net/http/pprof
以下代码会在http路由中加入debug/pprof

```
import (
    _ "net/http/pprof"
    ...
)
```
假如web服务器地址为 http://127.0.0.1:6060

web访问 http://127.0.0.1:6060/debug/pprof  
![](go性能分析/img_2.png)  
* allocs：查看过去所有内存分配的样本，访问路径为 $HOST/debug/pprof/allocs。
* block：查看导致阻塞同步的堆栈跟踪，访问路径为 $HOST/debug/pprof/block。
* cmdline： 当前程序的命令行的完整调用路径。
* goroutine：查看当前所有运行的 goroutines 堆栈跟踪，访问路径为 $HOST/debug/pprof/goroutine。
* heap：查看活动对象的内存分配情况， 访问路径为 $HOST/debug/pprof/heap。
* mutex：查看导致互斥锁的竞争持有者的堆栈跟踪，访问路径为 $HOST/debug/pprof/mutex。
* profile： 默认进行 30s 的 CPU Profiling，得到一个分析用的 profile 文件，访问路径为 $HOST/debug/pprof/profile。
* threadcreate：查看创建新 OS 线程的堆栈跟踪，访问路径为 $HOST/debug/pprof/threadcreate。


## gops

### gops 安装
go install github.com/google/gops@latest

### 注入采集代理
```
if err := agent.Listen(agent.Options{Addr: "0.0.0.0:6060"}); err != nil {
    log.Fatal(err)
}
```

### gops 查看pid
带星号的可使用gops命令
![](go性能分析/img.png)

### 采集cpu
* 程序在本地: gops pprof-cpu 54143
* 远程采集: gops pprof-cpu 127.0.0.1:6060

### 查看内存情况  gops memstats
gops memstats 127.0.0.1:6060
```
//已分配的对象的字节数 和HeapAlloc相同
alloc: 337.66MB (354064872 bytes)  

//分配的字节数累积之和,所以对象释放的时候这个值不会减少
total-alloc: 2.70TB (2970376940944 bytes) 

//从操作系统获得的内存总数  Sys是下面的XXXSys字段的数值的和, 是为堆、栈、其它内部数据保留的虚拟内存空间 注意虚拟内存空间和物理内存的区别
//sys 不是实际占用的物理内存大小， 是向操作系统申请的所有虚拟内存空间大小之和，等同于 XXX_sys + XXX_sys。 虚拟内存只保留了和物理内存的页表映射，不代表实际内存，go 也不会去做这个映射的解除，所以这个值不会掉。
sys: 1.85GB (1983248840 bytes) 

//运行时地址查找的次数，主要用在运行时内部调试上.
lookups: 0 

//堆对象分配的次数累积和 活动对象的数量等于Mallocs - Frees
mallocs: 72348366613 

//释放的对象数.
frees: 72345733547 

//分配的堆对象的字节数 包括所有可访问的对象以及还未被垃圾回收的不可访问的对象. 所以这个值是变化的，分配对象时会增加，垃圾回收对象时会减少.
heap-alloc: 337.66MB (354064872 bytes) 

//从操作系统获得的堆内存大小. 虚拟内存空间为堆保留的大小，包括还没有被使用的. HeapSys 可被估算为堆已有的最大尺寸.
heap-sys: 1.76GB (1885208576 bytes) 

//HeapIdle是idle(未被使用的) span中的字节数. Idle span是指没有任何对象的span,这些span **可以**返还给操作系统，或者它们可以被重用 或者它们可以用做栈内存.
//HeapIdle 减去 HeapReleased 的值可以当作"可以返回到操作系统但由运行时保留的内存量". 以便在不向操作系统请求更多内存的情况下增加堆，也就是运行时的"小金库".
//如果这个差值明显比堆的大小大很多，说明最近在活动堆的上有一次尖峰.
heap-idle: 1.34GB (1438941184 bytes) 

//正在使用的span的字节大小.
//正在使用的span是值它至少包含一个对象在其中.
//HeapInuse 减去 HeapAlloc的值是为特殊大小保留的内存，但是当前还没有被使用
heap-in-use: 425.59MB (446267392 bytes)

//HeapReleased 是返还给操作系统的物理内存的字节数.
//它统计了从idle span中返还给操作系统，没有被重新获取的内存大小.
heap-released: 1.21GB (1297219584 bytes)

//HeapObjects 实时统计的分配的堆对象的数量,类似HeapAlloc.
heap-objects: 2633066

//栈span使用的字节数。
//正在使用的栈span是指至少有一个栈在其中.
//注意并没有idle的栈span,因为未使用的栈span会被返还给堆(HeapIdle).
stack-in-use: 2.12MB (2228224 bytes)

//从操作系统取得的栈内存大小.
//等于StackInuse 再加上为操作系统线程栈获得的内存.
stack-sys: 2.12MB (2228224 bytes)

//分配的mspan数据结构的字节数.
stack-mspan-inuse: 3.47MB (3637184 bytes)

//从操作系统为mspan获取的内存字节数
stack-mspan-sys: 10.54MB (11048640 bytes)

//分配的mcache数据结构的字节数.
stack-mcache-inuse: 4.69KB (4800 bytes)

//从操作系统为mcache获取的内存字节数.
stack-mcache-sys: 15.23KB (15600 bytes)

//off-heap的杂项内存字节数.
other-sys: 2.11MB (2215889 bytes)

//垃圾回收元数据使用的内存字节数.
gc-sys: 74.55MB (78175424 bytes)

//下一次垃圾回收的目标大小，保证 HeapAlloc ≤ NextGC.
//基于当前可访问的数据和GOGC的值计算而得.
next-gc: when heap-alloc >= 528.37MB (554037040 bytes)

//上一次垃圾回收的时间.
last-gc: 2022-11-28 10:56:14.274883811 +0800 CST

//自程序开始 STW 暂停的累积纳秒数.
//STW的时候除了垃圾回收器之外所有的goroutine都会暂停.
gc-pause-total: 9.017849591s

//一个循环buffer，用来记录最近的256个GC STW的暂停时间.
gc-pause: 467466

//最近256个GC暂停截止的时间.
gc-pause-end: 1669604174274883811

//GC的总次数.
num-gc: 12970
// 强制GC的次数
num-forced-gc: 0
//自程序启动后由GC占用的CPU可用时间，数值在 0 到 1 之间.
//0代表GC没有消耗程序的CPU. GOMAXPROCS * 程序运行时间等于程序的CPU可用时间.
gc-cpu-fraction: 0.0011345879785226885
//是否允许GC.
enable-gc: true
debug-gc: false
```

### gops 功能
* gops 查看当前的运行的go程序，含有星号即可使用下面的命令
* gops +pid  简单查看当前状态
* gops trace + pid  
  * view trace 查看跟踪
  * goroutine analysis go协程分析，目前开着哪些协程
  * Network blocking profile 查看网络阻塞情况，看到网络耗时在哪部分比较多
  * Synchronization blocking profile 同步阻塞配置文件，查看哪个程序调用线路耗时较多
  * Syscall blocking profile 系统调用阻塞配置文件，系统调用的耗时显示
  * Scheduler latency profile 调度程序延迟配置文件
  * User-defined tasks 用户定义的任务
  * User-defined regions 用户定义区域
  * Minimum mutator utilization 最小mutator利用率
* gops tree + pid 显示当前调用的进程树
* gops stack +pid 显示当前栈使用情况
* gops memstats + pid 打印当前内存统计信息
* gops gc + pid 显示gc使用情况
* gops setgc +pid + 数字 将垃圾回收目标设置为特定百分比
* gops version +pid 报告构建目标程序的 Go 版本
* gops stats +pid 打印运行时统计信息
* gops pprof-cpu +pid
* gops pprof-heap +pid

> 远程查看，将pid替换为远端代理 例如: gops pprof-cpu 172.31.60.114:6060

# 非服务型 一次采集

## runtime/pprof
采集CPU
```go
package main

import (
  "flag"
  "log"
  "os"
  "runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
  flag.Parse()
  // 如果命令行设置了 cpuprofile
  if *cpuprofile != "" {
    // 根据命令行指定文件名创建 profile 文件
    f, err := os.Create(*cpuprofile)
    if err != nil {
      log.Fatal(err)
    }
    // 开启 CPU profiling
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()
  }
}

  
```

## go test 
```
go test . -v  -test.run Test_RunTestCase -cpuprofile cpu.pprof
```


# 图形化采集结果

## 安装graphviz
```
sudo apt install graphviz
```

## 其它图形化工具
* go get -u github.com/google/pprof
* go get github.com/uber/go-torch


## 本地文件
* go tool pprof -http=:8080 ./A文件
* go tool pprof ./A文件

## http
* go tool pprof -inuse_space http://localhost:6060/debug/pprof/heap
* go tool pprof http://localhost:6060/debug/pprof/goroutine
* go tool pprof http://localhost:6060/debug/pprof/profile\?seconds\=60
> 60s 表示采集时间，默认为30s

