package main

import (
	"context"
	"e/server/e"
	"fmt"
	"net"
	"sync"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	e.UnimplementedEServer
	// use map to store invite state
	mu sync.Mutex //ensure concurrent secure
	m  map[string]bool
}

func main() {
	// listen
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		// new Listener failed
		panic(err)
	}
	// new grpc client
	s := grpc.NewServer()
	// initilize map
	sv := &server{
		m: make(map[string]bool),
	}

	e.RegisterEServer(s, sv)
	err = s.Serve(l)
	if err != nil {
		// server failed
		panic(err)
	}
}

func (s *server) Err(ctx context.Context, req *e.ERequest) (resp *e.EResponse, err error) {
	// requirement:Do not allow twice request from same name
	// need a global map[string]bool
	name := req.GetName()
	s.mu.Lock()
	defer s.mu.Unlock()
	// not invite before
	if !s.m[name] {
		s.m[name] = true
		resp = &e.EResponse{
			Reply: "你好，" + name,
		}
		return resp, nil
	}
	// invited before
	st := status.New(codes.ResourceExhausted, "request limit")
	ds, err := st.WithDetails(
		&errdetails.QuotaFailure{
			Violations: []*errdetails.QuotaFailure_Violation{
				{
					Subject:     fmt.Sprintf("name:%s", req.Name),
					Description: "每个name只能调用一次该方法",
				},
			},
		},
	)
	// origin error
	if err != nil {
		fmt.Println("with detail failed,err:", err)
		return nil, st.Err()
	}
	// ds error
	return nil, ds.Err()
}
