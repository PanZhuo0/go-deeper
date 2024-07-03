package main

import (
	"log"
	"net"
	"sf/client/sf"

	"google.golang.org/grpc"
)

/*
	服务流

server端实现
*/
type server struct {
	sf.UnimplementedGreeterServer
}

func main() {
	// 1. 设置监听端口
	l, err := net.Listen("tcp", ":8080")
	// 2. new grpc server
	s := grpc.NewServer()
	if err != nil {
		log.Fatal(err)
		return
	}
	// 3.册服务到grpc server
	sf.RegisterGreeterServer(s, &server{})
	// 4.开始监听
	err = s.Serve(l)
	if err != nil {
		log.Fatal(err)
		return
	}
}

/*
由于这里是服务端是流式的，所以响应给client的结果只有错误,其他server结果通过流(入参中有)传递给client
*/
func (s *server) LotsOfReplies(req *sf.Request, stream sf.Greeter_LotsOfRepliesServer) (err error) {
	words := []string{
		"你好",
		"hello",
	}

	for _, word := range words {
		// 通过stream响应给client
		data := &sf.Response{
			Response: word + " " + req.GetName(),
		}
		// 使用stream的Send 方法返回给client
		if err := stream.Send(data); err != nil {
			return err //如果发送过程中出错,结束服务,响应err 给client
		}
	}
	return nil //如果没错,响应nil给client(代表返回没出错)
}
