package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	X, Y int
}

type ServiceA struct {
}

func (s *ServiceA) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

func main() {
	t, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("dial tcp failed,err:", err)
	}

	// sync call
	a := &Args{10, 30}
	var reply int
	err = t.Call("ServiceA.Add", a, &reply)
	if err != nil {
		log.Fatal("ServiceA.Add error:", err)
	}
	fmt.Println("RPC-TCP result:", reply)

	// async call
	var reply2 int
	divCall := t.Go("ServiceA.Add", a, &reply2, nil)
	replyCall := <-divCall.Done
	fmt.Println(replyCall.Error)
	fmt.Println(reply2)
}
