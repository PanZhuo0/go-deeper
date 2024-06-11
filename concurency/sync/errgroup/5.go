package main

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	var a, b, c []int
	// eg:调用广告服务
	g.Go(func() error {
		a = []int{0}
		return nil
	})

	// 调用AI
	g.Go(func() error {
		b = []int{1}
		return errors.New("ai")
	})

	// 调用运营平台数据
	g.Go(func() error {
		c = []int{2}
		return nil
	})

	err := g.Wait() //wait for 3 go routine end ,err will return first Err

	// a + b + c
	fmt.Println(a, b, c)
	fmt.Println(err)
	fmt.Println(ctx.Err())
}
