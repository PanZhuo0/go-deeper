package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// arguments of this server
type Args struct {
	X, Y int
}

// ServiceA 自定义一个结构体类型
type ServiceA struct{}

// Add 为ServiceA 类型增加一个可导出的Add方法
func (s *ServiceA) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

func main() {
	s := new(ServiceA)
	// register RPC server
	rpc.Register(s)
	rpc.HandleHTTP()

	// listen and serve
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("listen err:", err)
	}
	http.Serve(l, nil)
}
