---
title: python使用ZeroMQ
categories:
- python

 
tags:
- zmq
---

# 什么是 ZMQ
ZeroMQ 是一种通信框架，可以与多种编程语言和操作系统集成使用，但在具体实现中需要注意相应的语言和平台特性。   
ZeroMQ 是一种高性能、异步的消息传递库，可以在多个进程,线程和计算机之间传递消息，支持多种协议和通信模式。   
通信协议包含：进程内(线程)、进程间、机器间、广播，格式分别为：inproc://、ipc://、tcp://、pgm://。  
模式包含：独占对模式（PAIR），请求/应答(REQ/REP)，请求/应答代理(Router/Dealer)，发布/订阅(PUB/SUB)，管道(PUSH/PULL)



# ZMQ 基本概念
* Socket（套接字）：zmq的核心概念之一，它表示一个进程中用于发送或接收消息的一端。zmq中有多种不同类型的socket，每种类型都有不同的行为和使用场景。
* Message（消息）：zmq中的基本数据单位，它可以是任何形式的数据，例如字符串、字节序列、JSON对象等。zmq的消息传递是基于异步的消息队列实现的，发送方发送消息到队列中，接收方从队列中获取消息。
* Context（上下文）：zmq中所有socket的创建都需要一个context对象，它负责管理所有socket的生命周期和线程安全性。
* Pattern（模式）：zmq中有多种不同类型的socket，每种类型都有不同的通信模式，例如：REQ/REP、PUB/SUB、PUSH/PULL等。

## Socket（套接字）
Socket 是 ZeroMQ 的核心概念，它表示一种通信模式和协议，用于在进程或计算机之间传递消息

## Context
Context 是 ZeroMQ 中用于创建和管理 Socket 和线程的对象，它包含全局状态和运行环境。在一个应用中，通常只需要创建一个 Context 对象，然后通过它创建和管理多个 Socket 和线程。

## Message
Message 是 ZeroMQ 中用于传递数据的基本单位，它可以包含多个帧（Frame），每个帧都是一个二进制数据块，最大长度为 2^63 字节。消息是 ZeroMQ 的基本数据类型，可以在 Socket 之间传递和交换。

## Endpoint
Endpoint 是 ZeroMQ 中表示网络地址的概念，它由协议、IP 地址和端口号组成，用于标识 Socket 的网络位置。Endpoint 格式根据协议的不同而不同，例如 TCP 协议的 Endpoint 格式为 "tcp://IP:Port"，in-process 协议的 Endpoint 格式为 "inproc://Name"。

## 基本通信流程
1. 创建 Context 对象，用于管理 Socket 和线程。
2. 创建 Socket 对象，选择合适的通信模式和协议。
3. 绑定或连接 Socket，指定网络地址和端口号。
4. 发送消息，使用 send() 方法将消息发送给远程 Socket。
5. 接收消息，使用 recv() 方法从本地 Socket 接收消息。
6. 关闭 Socket 和 Context 对象，释放资源。


# 通信模式
## REQ/REP
REQ/REP 模式是一种简单的请求-响应模式，用于一对一的通信。在该模式下，客户端发送请求消息给服务端，服务端接收并处理请求，然后发送响应消息给客户端。  
该模式的优点：简单、易于理解和使用，缺点是单向流量、同步阻塞和无法扩展。在高并发和高可用性的场景下，不适合使用该模式。  

## PUB/SUB
PUB/SUB 模式是一种发布-订阅模式，用于一对多的通信。在该模式下，发布者（PUB）将消息广播给多个订阅者（SUB），订阅者接收并处理消息。  
该模式的优点是灵活、可扩展和异步非阻塞，缺点是无法保证消息的可靠性和顺序性。在需要高可靠性和消息顺序的场景下，不适合使用该模式。  

## PUSH/PULL
PUSH/PULL 模式是一种多对多的通信模式，用于推-拉模式的数据传输。在该模式下，Pusher（PUSH）将消息推送给多个 Puller（PULL），Puller 按顺序接收消息并处理。  
该模式的优点是高效、可扩展和异步非阻塞，缺点是无法保证消息的可靠性和顺序性。在需要高可靠性和消息顺序的场景下，不适合使用该模式。  

## PAIR
PAIR 模式是一种简单的对等通信模式，用于一对一的通信。在该模式下，两个 Socket 直接互相发送消息，没有中间的调度器。  
该模式的优点是简单、高效、可靠，缺点是无法扩展。在需要多对多的通信或需要扩展的场景下，不适合使用该模式。  

## DEALER/ROUTER
EALER/ROUTER 模式是一种分布式哈希表模式，通常用于多个进程之间的负载均衡和故障转移。DEALER Socket 充当工作线程，ROUTER Socket 充当调度器，将工作任务平均地分配给每个工作线程，并在某个工作线程出现故障时自动重新分配任务。

该模式的优点是可扩展、高效、可靠和动态故障转移，缺点是复杂度高、消息顺序性难以保证。在需要动态负载均衡和高可用性的场景下，适合使用该模式。  


# ZeroMQ 的高级特性
* 消息队列 ZeroMQ 支持基于 Socket 的消息队列，可以存储和缓冲消息，实现异步处理和流量控制。消息队列可以通过设置 Socket 的高水位线和低水位线来控制。
* 消息过滤 ZeroMQ 支持基于消息内容的订阅和过滤，可以根据消息的标签、类型和内容进行订阅和过滤。通过设置 Socket 的过滤规则，可以实现消息的精确过滤和路由。
* 多路复用 ZeroMQ 支持多路复用（Multiplexing），可以在单个 Socket 上实现多个通信模式和协议的混合使用。通过使用不同的消息头和消息体，可以实现多种通信模式的无缝切换和共存。
* 安全性 ZeroMQ 支持基于 SSL/TLS 的加密和认证，可以保证通信的安全性和机密性。通过使用 CurveZMQ 或 GSSAPI 机制，可以实现更高级的身份验证和加密算法。
* 多线程支持 支持多线程并发操作，可以实现多个线程同时处理不同的消息，并避免线程间的竞争和阻塞。
* 异步处理 支持异步消息处理，可以实现消息的快速发送和接收，避免线程的阻塞和等待。
* 数据序列化 ZeroMQ 支持多种数据序列化方式，如 JSON、MessagePack、Protobuf 等，可以实现数据的高效传输和存储。


# ZMQ的API
ZMQ的API主要分为两类：一类是基于socket的API，它提供了socket的创建、配置、发送和接收等功能；另一类是基于context的API，它提供了context的创建和销毁等功能。

## Socket API  
* zmq_socket：创建一个socket，需要指定socket类型和context对象。
* zmq_bind：将socket绑定到一个地址上，使得其他进程可以通过该地址与此socket进行通信。
* zmq_connect：将socket连接到一个地址上，使得此socket可以与该地址对应的socket进行通信。
* zmq_send：将一个消息发送到指定的socket上。
* zmq_recv：从指定的socket上接收一个消息。
* zmq_setsockopt：设置socket的选项，例如设置超时时间、设置是否启用心跳检测等。
* zmq_getsockopt：获取socket的选项值。
* zmq_close：关闭socket。

## Context API
* zmq_ctx_new：创建一个新的context对象。
* zmq_ctx_term：销毁一个context对象。
* zmq_ctx_set：设置context的选项，例如设置最大线程数、设置IO线程数等。
* zmq_ctx_get：获取context的选项值


