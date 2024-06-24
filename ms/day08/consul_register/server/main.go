package main

import (
	"c/client/c"
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type service struct {
	c.UnimplementedConsulServer
}

func main() {
	// 1.listen
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net.Listen failed,err:", err)
		return
	}
	/* consul Append */
	// new consul client
	// config Health check
	cc, err := api.NewClient(&api.Config{
		Address: `localhost:8500`,
		Scheme:  `http`,
	})

	if err != nil {
		fmt.Println("api.NewClient failed,err:", err)
		return
	}

	// registe config
	const serviceName = "hello"
	serviceID := fmt.Sprintf("%s-%s-%d", serviceName, `192.168.18.136`, 8972) //服务唯一ID
	srv := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceName,
		Tags:    []string{"zhangsan"},
		Address: `192.168.18.136`,
		Port:    8972,
		Check: &api.AgentServiceCheck{
			GRPC:                           fmt.Sprint(GetOutboundIP()), //外网地址
			Timeout:                        `5s`,                        //超时时间
			Interval:                       `5s`,                        //发送请求的间隔
			DeregisterCriticalServiceAfter: `1m`,                       //1m后注销不健康的节点
		},
	}
	// register server to consul
	err = cc.Agent().ServiceRegister(srv)
	if err != nil {
		fmt.Println(err)
	}
	/* End */

	// 2.new grpc server
	s := grpc.NewServer()
	hc := health.NewServer()
	healthpb.RegisterHealthServer(s, hc) //s:grpc 服务,hc 对应grpc的HC服务
	// 3.register server to pb
	c.RegisterConsulServer(s, &service{})
	// 4.run
	go func() {
		err = s.Serve(l)
		if err != nil {
			fmt.Println("failed to server,err:", err)
			return
		}
	}()
	// 5.优雅注销服务
	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGTERM, syscall.SIGINT)
	fmt.Println("wait quit signal...")
	<-quitCh //没收到停止信息,就会一直阻塞
	fmt.Println(`service out...`)
	err = cc.Agent().ServiceDeregister(serviceID)
	if err != nil {
		fmt.Println("ServiceDeregister failed,err:", err)
		return
	}

}

/* Consul Append
Consul 注册的使用需要使用外网可访问的IP
*/

func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("dup", "8.8.8.8")
	if err != nil {
		fmt.Println("GetOutboundIP failed,err:", err)
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}

func (s *service) SayHello(ctx context.Context, r *c.Request) (res *c.Response, err error) {
	return &c.Response{Reply: "你好," + r.GetName()}, nil
}
