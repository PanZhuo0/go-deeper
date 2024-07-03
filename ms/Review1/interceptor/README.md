### 拦截器(类似中间件)
    存在的问题:
        有时我们希望在发送请求前增加一部分信息/进行某些操作,或者发送请求后增加一部分信息/进行某些操作

    拦截器可以用于:
        1.进行日志记录
        2.身份验证/授权
        3.指标收集
        4.其他可以跨RPC共享的功能

### 拦截器种类
    客户端一元拦截器
        调用RPC前
        调用RPC时
        调用RPC后
    客户端流拦截器
        调用RPC前
        流式调用RPC时

    服务端一元拦截器
    服务端流拦截器

### 拦截器如何实现?
    一元拦截器

    流式拦截器
        底层实现:外层包装原本的stream,实现Stream接口中的Recev(),SendAndClose(),以及其他方法,再把这个wrappedStream当成原本的stream使用
                从而，当用户调用Recv和SendAndClose()以及其他方法时,实现类似中间件的功能
            先执行自定义的操作，再实现原本stream应该执行的功能


# Demo
    该Demo实现单元拦截器以及流式拦截器,实现Token认证
    自己实现具体的拦截器代码
        1.单元拦截器方法
        2.流拦截器方法


> Reference: https://liwenzhou.com/posts/Go/gRPC/#c-0-9-2




### 社区开源的GRPC拦截器工具库
https://github.com/grpc-ecosystem/go-grpc-middleware

