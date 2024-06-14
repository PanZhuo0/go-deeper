package main

import (
	"context"
	"net"
	"tls/server/tls"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	tls.UnimplementedTServer
}

func main() {
	// listen
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		// new listener failed
		panic(err)
	}
	// add credentials from file
	creds, err := credentials.NewServerTLSFromFile("../server.crt", "../server.key")
	if err != nil {
		// load credentials failed
		panic(err)
	}
	s := grpc.NewServer(grpc.Creds(creds)) //add credentials
	tls.RegisterTServer(s, &server{})
	err = s.Serve(l)
	if err != nil {
		// server failed
		panic(err)
	}
}

func (s *server) THello(ctx context.Context, req *tls.TReq) (resp *tls.TResp, err error) {
	name := req.GetName()
	resp = &tls.TResp{
		Reply: "你好,nihao," + name,
	}
	return resp, nil
}
