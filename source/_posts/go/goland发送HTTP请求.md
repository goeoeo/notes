---
categories: 
- go
tags:
- goland
---

# 简介
在工作中我们需要测试api，不免需要用到发送HTTP 请求的工具，比如postman ,这里工具有几个缺点： 
1. 共享数据比较麻烦
2. 不可纳入版本管理

这里介绍使用goland的 HTTP 客户端来发送请求

<!--more-->
先建立一个目录 http_test,以下所有的文件都在这个目录下面

# 变量设置

## 自定义变量
创建文件 vim http-client.env.json 内容为：
```json
{
  "dev": {
    "host": "127.0.0.1:9800",
    "token": "xx"
  }
}
```
在后续的 *.http文件中即可以选择 dev环境，并使用其中的变量

## 动态变量
* $uuid：生成一个通用唯一标识符（UUID-v4）
* $timestamp: 生成当前的 UNIX 时间戳
* $randomInt: 生成 0 到 1000 之间的随机整数。


# 发送请求
创建文件 test.http 内容如下：
```
### 创建
// @no-log
POST {{host}}/exp-distribution-level
Accept: */*
Cache-Control: no-cache
Content-Type: application/json
X-Requested-With:XMLHttpRequest
Authorization: Bearer {{token}}

{
  "name": "test",
  "discount_order": 1
}


### 更新
PATCH {{host}}/exp-distribution-level
Accept: */*
Cache-Control: no-cache
Content-Type: application/json
X-Requested-With:XMLHttpRequest
Authorization: Bearer {{token}}

{
  "id": 3,
  "name": "test1",
  "discount_order": 2
}


### 删除
DELETE {{host}}/exp-distribution-level
Accept: */*
Cache-Control: no-cache
Content-Type: application/json
X-Requested-With:XMLHttpRequest
Authorization: Bearer {{token}}

{
  "id": 2
}


### 获取数据
GET {{host}}/exp-distribution-level/2
Accept: */*
Cache-Control: no-cache
Content-Type: application/json
X-Requested-With:XMLHttpRequest
Authorization: Bearer {{token}}

```

# 上传文件

## 单文件上传
```
### 单文件上传
POST {{host}}//public/ocr
Accept: */*
Cache-Control: no-cache
Content-Type: multipart/form-data;boundary=boundary
Authorization: Bearer {{token}}

--boundary
Content-Disposition: form-data; name="file"; filename="img.png"

< img.png
```

## 多文件上传

```
### 单文件上传
POST {{host}}//public/ocr
Accept: */*
Cache-Control: no-cache
Content-Type: multipart/form-data;boundary=boundary
Authorization: Bearer {{token}}

--boundary
Content-Disposition: form-data; name="first"; filename="input.txt"

< ./input.txt

--boundary
Content-Disposition: form-data; name="second"; filename="input-second.txt"

< ./input-second.txt

--boundary
Content-Disposition: form-data; name="third";

< ./input.txt --boundary--
```

# 请求设置
* 禁用重定向  @no-redirect
* 禁用将请求保存到请求历史记录  @no-log
* 禁用将收到的 cookie 保存到 cookie jar @no-cookie-jar

示例：
```
// @no-redirect @no-log @no-cookie-jar
POST {{host}}//public/ocr
```


# 处理响应
您可以使用 JavaScript 处理响应。键入>请求后的字符并指定 JavaScript 文件的路径和名称或将响应处理程序脚本代码包裹在{% ... %}.

```
GET {{host}}/get

> /path/to/responseHandler.js
```

```
GET {{host}}/get

> {%
    client.global.set("my_cookie", response.headers.valuesOf("Set-Cookie")[0]);
%}
```

# 重定向响应
可以将响应重定向到文件。如果文件已存在，则用于>>创建带有后缀的新文件，如果文件存在>>!则重写该文件。您可以指定绝对路径或相对于当前 HTTP 请求文件
的相对路径。您还可以在路径中使用变量，包括环境变量和以下预定义变量：

* {{$projectRoot}}指向项目根目录：.idea
* {{$historyFolder}}指向.idea /httpRequests/

以下示例 HTTP 请求在.idea /httpRequests/中创建myFile.json。如果文件已经存在，它会覆盖该文件。它还使用位于项目根目录中的handler.js脚本处理响应。

```
POST https://httpbin.org/post
Content-Type: application/json

{
  "id": 999,
  "value": "content"
}

> {{$projectRoot}}/handler.js

>>! {{$historyFolder}}/myFile.json
```

# 参考
[GoLand HTTP请求语法](https://www.javatiku.cn/goland/2675.html)