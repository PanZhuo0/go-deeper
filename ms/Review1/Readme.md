# GRPC
## protobuf 语法
    详情见 ./protobuf_syntax
        1.protobuf中Message的数据类型 
        2.protobuf中Message中的一些特殊的关键字-oneof map repeated 
        3.protobuf中Service的定义
        4.google提供的常用的.proto文件-wrappers fieldmask 如何使用
        5.google提供的.proto文件    -Any 结构体

## protobuf 使用
    详情见 ./grpc_add_demo
        1.一个简单的grpc使用demo
        2.简单的protoc 命令
        3.用Go实现一个简单的RPC add功能

## 流式RPC
    详情见 ./grpc_flow_demo
        1.客户端流
        2.服务端流
        3.双向流

## GRPC Metadata(Header & Trailer)
    详情见 ./metadata
    客户端MD
    服务端MD
    Trailer、Header MD
    普通RPC、流RPC MD

## GRPC ErrorHandle
    详情见 ./error
    GRPC官方的错误处理
    Reference:
        https://grpc.io/docs/guides/status-codes/

    1.grpc Status status错误信息
    2.status.WithDetails 给错误增加额外的信息
    3.客户端解析status 中的错误Details信息

## TLS加密
    跳过,后期继续学习

## GRPC 拦截器(中间件)
    详情见 ./interceptor
    1.客户端一元拦截器
        三个阶段:
            预处理(调用前)
            RPC调用
            调用后
    2.客户端流式拦截器
        两个阶段：
            预处理(调用前)
            流操作拦截
    3.服务端一元拦截器
    4.服务端流式拦截器

## GRPC Gateway,映射规则

## Addtion

### 1.基于游标的分页
    跳过

### 2.同一个端口服务gRPC 和 HTTP


### gRPC 名称解析与负载均衡

# 服务注册与发现


