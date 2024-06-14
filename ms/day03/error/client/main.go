package main

import (
	"context"
	"e/client/e"
	"fmt"
	"time"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		// dial grpc failed
		panic(err)
	}
	c := e.NewEClient(conn)
	// context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// prarmeter
	name := "张三"
	resp, err := c.Err(ctx, &e.ERequest{Name: name})
	if err != nil {
		// try to unwarp err detail
		s := status.Convert(err)
		for _, d := range s.Details() {
			switch info := d.(type) {
			case *errdetails.QuotaFailure:
				fmt.Println("QutotaFailure:", info.Violations)
			default:
				fmt.Println("unexpected type:", info)
			}
		}
	} else {
		fmt.Println(resp.Reply)
	}
}
