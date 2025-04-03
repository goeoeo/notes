---
title: k8s中对grpc服务进行健康检查
categories: 
- k8s
tags:
- k8s
- grpc
---

# 健康检查
健康检查（Health Check）是一种用于探测应用或服务是否处于正常运行状态的机制，通常通过预设规则主动检测系统组件（如容器、微服务、数据库等）的可用性，是保障应用高可用性和自愈能力的核心手段。

健康检查的目的：
1. 故障自动恢复：快速发现不可用实例并触发重启（如 Kubernetes 的 livenessProbe）。
2. 流量控制：避免将请求转发到未就绪或异常的服务（如 readinessProbe 控制 Pod 是否加入 Service Endpoints）。
3. 状态可视化：为监控系统（Prometheus、Zabbix）提供健康状态指标。
4. 优雅启停：确保应用完成初始化后再接收流量，关闭前完成未处理请求。
<!--more-->

# k8s的健康检查
k8s健康检查有3类：
* 存活探针（Liveness Probe）：检测容器是否崩溃（触发重启）。
* 就绪探针（Readiness Probe）：检测容器是否准备好接收流量（控制 Service Endpoints）。
* 启动探针（Startup Probe）：延迟其他探针直到应用完成启动（适用于慢启动应用）。

## livenessProbe 与 readinessProbe 的区别

| 特性  | livenessProbe  |  readinessProbe |
|---|---|---|
| 目的  |  检测容器是否存活（是否需要重启） | 检测容器是否准备好接收流量（是否加入 Service 负载均衡）  |
|  失败后果 | 重启容器（触发容器重建）  |  将 Pod 从 Service 的 Endpoints 列表中移除 |
|  适用阶段 | 容器整个生命周期内持续监控  |  容器启动后是否初始化完成、是否临时不可用 |
| 典型检查逻辑  |  判断应用是否完全崩溃（如死锁、内存泄漏） |  判断应用是否依赖外部资源（如数据库连接是否就绪） |


# grpc服务的健康检查
gRPC 健康检查协议是 gRPC 官方定义的标准健康检查机制，用于让客户端（如 Kubernetes、服务网格）探测服务端的状态是否健康。其核心是基于 grpc.health.v1.Health 服务的接口实现。

## 协议定义
健康检查协议通过 Protocol Buffers 定义，需在服务端实现以下接口
```protobuf 
// health.proto
syntax = "proto3";

package grpc.health.v1;

service Health {
  // 检查单个服务的健康状态
  rpc Check(HealthCheckRequest) returns (HealthCheckResponse);
  // 流式监控服务状态变更（可选）
  rpc Watch(HealthCheckRequest) returns (stream HealthCheckResponse);
}

message HealthCheckRequest {
  string service = 1;  // 要检查的服务名称（空字符串表示整体服务）
}

message HealthCheckResponse {
  enum ServingStatus {
    UNKNOWN = 0;
    SERVING = 1;
    NOT_SERVING = 2;
    SERVICE_UNKNOWN = 3;  // 服务不存在
  }
  ServingStatus status = 1;
}

```
## 服务端代码
```
import (
  "google.golang.org/grpc"
  "google.golang.org/grpc/health"
  "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
  // 创建 gRPC 服务器
  srv := grpc.NewServer()

  // 初始化健康检查服务
  healthSrv := health.NewServer()
  healthSrv.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING) // 默认服务状态
  healthSrv.SetServingStatus("my-service", grpc_health_v1.HealthCheckResponse_SERVING) // 自定义服务状态

  // 注册健康检查服务
  grpc_health_v1.RegisterHealthServer(srv, healthSrv)

  // 启动服务器
  lis, _ := net.Listen("tcp", ":50051")
  srv.Serve(lis)
}

```

## 客户端检查
使用官方工具 grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/
```
# 下载
wget https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.24/grpc_health_probe-linux-amd64 -O /usr/local/bin/grpc_health_probe

# 检查整体服务状态
grpc_health_probe -addr=localhost:50051

# 检查特定服务状态
grpc_health_probe -addr=localhost:50051 -service=my-service

```

## 与k8s集成
Kubernetes 从 1.23 版本 开始原生支持通过 grpc 类型的探针（livenessProbe 和 readinessProbe）进行 gRPC 服务健康检查。
```
    spec:
      containers:
        - name: server
          image: my-image
          ports:
            - containerPort: 50051
              protocol: TCP

          # 配置 livenessProbe
          livenessProbe:
            grpc:
              port: 50051         # 监听的 gRPC 端口
              service: "my-service"  # 可选，检查指定服务名的健康状态
            initialDelaySeconds: 5  # 容器启动后等待 5 秒开始探测
            periodSeconds: 10       # 每 10 秒探测一次
            timeoutSeconds: 2       # 超时时间 2 秒
            failureThreshold: 3     # 连续失败 3 次后标记为不健康
```

如果k8s版本小于1.23，需要将grpc_health_probe 命令封装到image中，然后在livenessProbe中定义command去实现检查


# k8s中其它的检查类型
## HTTP(S) 检查
原理：向应用的 HTTP 端点发送 GET 请求，根据状态码（如 200 OK）判断健康状态。  
适用场景：Web 服务、RESTful API。  
配置示例（Kubernetes）：  
```
readinessProbe:
  httpGet:
    path: /health
    port: 8080
  initialDelaySeconds: 10
```

## TCP 检查
原理：尝试建立 TCP 连接，判断端口是否可访问。  
适用场景：数据库、消息队列（如 MySQL、Redis）。  
示例：
```
livenessProbe:
  tcpSocket:
    port: 3306
```


