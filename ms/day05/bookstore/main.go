package main

import (
	"bookstore/pb"
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 连接数据库
	db, err := NewDB("bookstore.db")
	if err != nil {
		fmt.Println("connect to DB failed,err:", err)
		return
	}

	// 创建server
	src := server{
		bs: &bookstore{db: db},
	}
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("failed to listen,err:", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterBookstoreServer(s, &src)
	go func() {
		// grpc server
		err = s.Serve(l)
		if err != nil {
			fmt.Println("start server failed,err:", err)
			return
		}
	}()
	// grpc-gateway server
	conn, err := grpc.DialContext(
		context.Background(),
		"localhost:8080", //target grpc server port
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Println("grpc conn failed,err:", err)
		return
	}

	gwmux := runtime.NewServeMux()
	pb.RegisterBookstoreHandler(context.Background(), gwmux, conn)

	gwServer := http.Server{
		Addr:    "localhost:8081",
		Handler: gwmux,
	}

	fmt.Println("grpc-Gateway serve on:8081")
	gwServer.ListenAndServe()
}
