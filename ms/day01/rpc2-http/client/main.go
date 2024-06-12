package main

import (
	"fmt"
	"log"
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
	// build connection with server
	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("dialing err:", err)
	}

	// sync call
	args := &Args{10, 20}
	var reply int
	err = client.Call("ServiceA.Add", args, &reply)
	if err != nil {
		log.Fatal("SeverA.Add error:", err)
	}
	fmt.Println("RPC result:", reply)

	// async call
	var reply2 int
	divCall := client.Go("ServiceA.Add", args, &reply2, nil) //return a chan，if done mean call got response(could be fail)
	replyCall := <-divCall.Done                              //receive call result
	fmt.Println(replyCall.Error)
	fmt.Println(reply2)
}
