package main

import (
	"context"
	hello "d9/client/server"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

var port = 8976

type server struct {
	hello.UnimplementedGreeterServer
}

func main() {
	// 从命令行获取端口
	flag.IntVar(&port, "port", 8976, "端口")
	flag.Parse()
	// 监听该端口
	l, err := net.Listen(`tcp`, fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("net listen failed,err:", err)
		return
	}

	// 注册服务到Consul
	cc, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Println(`get consul Client failed,err:`, err)
		return
	}
	err = cc.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      `hello-1`,
		Name:    `hello`,
		Tags:    []string{`hello`, `test`},
		Address: `localhost`,
		Port:    port,
		Check: &api.AgentServiceCheck{
			GRPC:                           `localhost`,
			Timeout:                        `5s`,
			Interval:                       `5s`,
			DeregisterCriticalServiceAfter: `1m`,
		},
	})
	if err != nil {
		fmt.Println(`服务注册,注册服务到Consul中出错,err:`, err)
		return
	}
	// 注册服务到grpc
	s := grpc.NewServer()
	hello.RegisterGreeterServer(s, &server{})
	// 启动服务
	go func() {
		err = s.Serve(l)
		if err != nil {
			fmt.Println(`start server failed,err:`, err)
			return
		}
	}()
	// 优雅注销服务
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	fmt.Println(`wait quit signal...`)
	<-sigCh //等待信号
	fmt.Println(`service out...`)

}

func (s server) SayHello(ctx context.Context, req *hello.HelloRequest) (resp *hello.HelloResponse, err error) {
	return &hello.HelloResponse{Reply: "你好," + req.GetName()}, nil
}
