package main

import (
	"cf/cf"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

/*
	客户端流

server端实现
*/
type server struct {
	cf.UnimplementedCFServer
}

func main() {
	// 1.监听
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
	// 2.new grpc server
	s := grpc.NewServer()
	// 3. register
	cf.RegisterCFServer(s, &server{})
	// 4.开始监听
	err = s.Serve(l)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func (s *server) Gretter(stream cf.CF_GretterServer) error {
	// 从流中读取客户端的信息
	res := `你好,`
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// client主动关闭了stream,表明客户端已经发送完想发送的数据了
			// server端需要把对应结果通过流一次性响应给client
			err := stream.SendAndClose(&cf.Response{Reply: res})
			if err != nil {
				log.Fatal("服务端写入客户流过程中出错", err)
				return err
			}
			break
		}
		// 传输过程中出错了
		if err != nil {
			return err
		}
		// 每次client发送一次请求，就记录一个结果,增加到response数组中
		res += req.GetName() + " "
	}
	return nil
}
