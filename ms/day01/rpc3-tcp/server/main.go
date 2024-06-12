package main

import (
	"log"
	"net"
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
	s := new(ServiceA)
	rpc.Register(s)
	// rpc.HandleHTTP()  based on TCP
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("listen tcp failed,err:", err)
	}
	// based on RPC process RPC connection
	for {
		conn, _ := l.Accept()
		rpc.ServeConn(conn)
	}
}
