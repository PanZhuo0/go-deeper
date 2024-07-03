package main

import (
	"context"
	"fmt"
	"interceptor/client/interceptor"
	"log"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/credentials/oauth"
)

/*拦截器Demo客户端实现*/
func main() {
	conn, err := grpc.NewClient(`localhost:3306`,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		/* 需要在NewClient时增加上一元拦截器选项 */
		grpc.WithUnaryInterceptor(unaryInterceptor))
	if err != nil {
		log.Fatal("err", err)
		return
	}
	c := interceptor.NewGreeterClient(conn)
	resp, err := c.SayHello(context.Background(), &interceptor.Request{Name: "张三"})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(resp.Reply)
}

/*定义一个服务端一元拦截器,该拦截器实现了如果请求没有增加认证证书，则给该请求增加认证证书*/
func unaryInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	var credsConfigurated bool
	for _, o := range opts {
		_, ok := o.(grpc.PerRPCCredsCallOption) //类型断言，查看是否配置了认证证书
		if ok {                                 //配置了
			credsConfigurated = true
			break
		}
	}
	if !credsConfigurated {
		// 如果没设置配置证书
		opts = append(opts, grpc.PerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: "some-secret-token",
		})))
	}
	// 同时计算一下程序的耗时
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	end := time.Now()
	fmt.Printf("RPC: %s ,Start Time:%s, End Time:%s,err:%v\n", method, start.Format("Basic"), end.Format(time.RFC3339), err)
	return err
}
