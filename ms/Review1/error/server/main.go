package main

import (
	"context"
	"error/client/e"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*
	在serve中定义一个visted map

key:用户名
value:bool类型 表示是否访问过

对于Map需要注意的问题:

	1.线程安全?
	2.map是否初始化
*/
type server struct {
	e.UnimplementedGreeterServer
	// 用来判断某个用户是否已经访问过了
	mux     sync.Mutex
	visited map[string]bool
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
	s := grpc.NewServer()
	e.RegisterGreeterServer(s, &server{visited: make(map[string]bool)}) //记得初始化
	err = s.Serve(l)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func (s *server) SayHello(ctx context.Context, req *e.Request) (resp *e.Response, err error) {
	// 加锁,确保map线程安全
	s.mux.Lock()
	defer s.mux.Unlock()

	name := req.GetName()
	ok := s.visited[name]
	if ok {
		// status 基本错误信息
		st := status.New(codes.ResourceExhausted, "request limit")
		// 往 status 中增加错误详细信息
		ds, err := st.WithDetails(
			&errdetails.QuotaFailure{
				Violations: []*errdetails.QuotaFailure_Violation{
					{
						Subject:     fmt.Sprintf("name:%s", req.GetName()),
						Description: "每个Name 只能调用一次",
					},
				},
			},
		)
		if err != nil {
			fmt.Println(err)
			return nil, st.Err() //不携带额外信息的error
		} else {
			fmt.Println(ds.Details()...)
			return nil, ds.Err() //详细的error
		}
	} else {
		resp = &e.Response{
			Reply: "你好," + name,
		}
		s.visited[name] = true //置为已访问
		return resp, nil
	}
}
