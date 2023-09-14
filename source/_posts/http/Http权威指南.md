---
title: Http权威指南
categories: 
- http
---

# 第1章 HTTP概述
## 资源
### 媒体类型
* MIME(Multipurpose Internet Mail Extension,多用途英特网邮件扩展)
* Web服务器回味所有的HTTP对象数据附加一个MIME类型，以告知Web浏览器该如何处理该对象
* Content-Type:image/jpeg ,“image/jpeg”为MIME类型
<!--more-->

### URI
* URI(Uniform Resource Identifier,统一资源标识符)
* URI 有两种形式，分别称为URL和URN

### URL
* 统一资源定位符
* URL包含三个部分,协议类型,服务器地址,资源地址
* 现在,几乎所有的URI都是URL
* URL语法 : <scheme>://<user>:<password>@<host>:<port>/<path>;<params>?<query>#<frag>

### URN
* URN 统一资源名
* URN是作为特定内容的唯一名称使用的,与目前资源所在地无关
* P2P下载中使用的磁力链接是URN的一种实现，它可以持久化的标识一个BT资源，资源分布式的存储在P2P网络中，无需中心服务器用户即可找到并下载它。

## Web的结构组件
* 代理 位于客户端和服务端之间的HTTP中间实体
* 缓存 HTTP的仓库,使常用的页面或资源保存在离客户端更近的地方
* 网关 连接其他应用程序的特殊问web服务器
* 隧道 对HTTP通信报文进行盲转发的特殊代理
* Agent代理 发起自动HTTP请求的半智能Web客户端

# 第3章 HTTP报文
## 报文结构
* 起始行(start line)
* 首部(header)
* 主体(body)

# 第4章 连接管理

## 用TCP套接字编程
* 网络套接字 操作系统提供的操作TCP连接的工具

### 服务端套接字编程流程
1. 创建套接字
2. 绑定端口
3. 监听端口
4. 接收并处理消息

### 客户端套接字编程流程
1. 创建套接字
2. 连接服务端套接字
3. 发送消息

## 性能聚焦区域
###　TCP相关时延
* TCP连接建立握手
* TCP慢启动拥塞控制
* 数据聚集的Nagle算法
* 用于捎带确认的TCP延迟确认算法
* TIME_WAIT时延和端口耗尽


## 提高HTTP连接性能的方法
串行事务处理时延
* 并行连接 通过多条TCP连接发起并发的HTTP请求
* 持久连接 重用TCP连接,以消除连接及关闭时延
* 管道化连接 通过共享TCP连接发起并发的HTTP请求
* 复用连接 交替传送请求和响应报文


### 并行连接
打开大量连接会消耗很多内存资源
浏览器确实使用了并行连接,但他们会将并行连接的总数限制为一个较小的值(通常是4个). 服务器可以随意关闭来自特定客户端的超量连接
### 持久连接
在HTTP事务处理结束后仍然保持在打开状态的TCP连接被称为持久连接;持久连接会在不同事务之间保持打开状态,直到客户端或服务器决定将其关闭为止

### 不能被代理转发或缓存响应使用的首部
* Connection
* Proxy-Authenticate
* Proxy-Connection
* Transfer-Encoding
* Upgrade


## 持久连接
### HTTP/1.0
HTTP/1.0 通过Connection:keep-alive 头部发送持久连接信号 ,服务端响应Connection:keep-alive表示支持持久连接,但是会存在一个盲中继问题
盲中继(哑代理)会导致,客户端保持TCP连接,服务端保持TCP连接,他们的TCP连接都是连接在代理上的,代理却什么都不知道
> 只是将一个连接转发到另一个连接去,不对Connection首部进行特殊处理
### HTTP/1.1
HTTP/1.1使用persistent connection持久连接,改进了HTTP/1.0 中Connection:keep-alive的缺陷,HTTP/1.1中持久连接是默认激活的,要关闭则需要在报文首部显示添加Connection:close关闭持久连接

### 持久连接的限制
* 只有当连接上的所有报文都是正确的,自定义报文长度时,也就是说实体部分的长度和相应的Content-length一致,或则是用分块传输编码方式编码的,连接才能持久保持
* HTTP/1.1的代理必须能够分别管理与客户端和服务端的持久连接,每个连接都只适用于一跳传输
* 一个用户客户端对任何服务器或代理最多维持两条持久连接,以防止服务器过载.


# 第5章 Web服务器