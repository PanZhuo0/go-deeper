package main

import (
	"flow/client/flow"
	"fmt"
	"io"
	"net"
	"strings"

	"google.golang.org/grpc"
)

type server struct {
	flow.UnimplementedFlowRpcServer
}

func main() {
	// 监听端口设置
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	// grpc注册
	s := grpc.NewServer()
	flow.RegisterFlowRpcServer(s, &server{}) //register server to grpc

	// 启动服务
	err = s.Serve(l)
	if err != nil {
		panic(err)
	}
}

func (s *server) Reqs(stream flow.FlowRpc_ReqsServer) error {
	reply := "你好,"
	for {
		// 从client的流中获取数据
		req, err := stream.Recv()
		if err == io.EOF {
			// 统一回复
			return stream.SendAndClose(&flow.HelloResponse{
				Return: reply,
			})
		}

		if err != nil {
			panic(err)
		}

		// 拼接信息
		reply += req.GetName()
	}
}

func (s *server) Replies(r *flow.HelloRequest, stream flow.FlowRpc_RepliesServer) error {
	words := []string{
		"你好",
		"Hello",
		"Привет",
		"Bonjour",
		"Olá;",
		"Hola",
	}

	for _, word := range words {
		data := &flow.HelloResponse{
			Return: word + "," + r.GetName(),
		}
		if err := stream.Send(data); err != nil {
			return err
		}
	}

	return nil
}

// 双向流
func (s *server) Bi(stream flow.FlowRpc_BiServer) error {
	// receive msg from client,response magic(request.Msg) to client
	for {
		req, err := stream.Recv()
		if err != nil {
			fmt.Println(err)
			return err
		}
		reply := magic(req.GetName())
		err = stream.Send(&flow.HelloResponse{Return: reply})
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
}

// magic 一段价值连城的“人工智能”代码
func magic(s string) string {
	s = strings.ReplaceAll(s, "吗", "")
	s = strings.ReplaceAll(s, "吧", "")
	s = strings.ReplaceAll(s, "你", "我")
	s = strings.ReplaceAll(s, "？", "!")
	s = strings.ReplaceAll(s, "?", "!")
	return s
}
