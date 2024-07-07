> Reference https://go-kratos.dev/en/docs/getting-started/start/

安装
>  go install github.com/go-kratos/kratos/cmd/kratos/v2@latest

创建项目
```bash
    kratos new helloworld
    cd helloworld
    go mod download
    go mod tidy
```

kratos 运行项目
```bash

    kratos run

```

手动启动,这些命令在kratos中的Makefile有定义
```bash

mkdir ./bin
go build -o ./bin  ./...
./bin/helloworld -conf ./configs/config.yaml

```


##### 测试
postman
curl -x get localhost:8000/helloworld/zhangsan


### 项目的目录结构
> 下面是文件夹
api
    各种.proto文件以及protobuf生成的代码

bin
    可执行的命令

cmd 
    main文件所处的位置,以及一些依赖注入文件

configs
    放置配置文件的部分

internal(不会对外暴露,外部无法导入使用)
    存放业务、逻辑、具体实现的文件夹 (从底向上结构)
    (配置)conf:存放配置文件结构体的文件夹(基于protobuf生成,KRATOS的特点:全部用.proto来生成结构体)
    (dao)data:存放dao层的代码,包括model
    (logic)biz:存放业务逻辑的代码文件夹 
    (server)service:逻辑调用层，这一层实现服务
    (运输层/handler)server:服务的运输层:HTTP/GRPC/其他,这一层实现将服务结果运输出去

third_party:
    用来放第三方库的数据,比如google的.proto一些支持文件


> 下面是文件
Dockerfile
    Kratos使用Docker的分层构建、提供一个可以在Docker运行的Dockerfile，便于我们在容器中创建微服务

go.mod
go.sum

Makefile
    用来简化开发者写cmd命令的，用起来很方便，在C/C++中非常常用,可以提高开发效率,可以理解为命令行的脚本,批处理


