package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {

}

/*
1.将控制权交给函数调用者
2.使用超时控制context，控制被调用的函数
*/
func process(term string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	ch := make(chan result) //result:是一个结构体

	go func() {
		record, err := search(term)
		ch <- result{record, err}
	}()

	// Block waiting to either receive from the goroutine's
	// channel or for the context to be canceled
	select {
	case <-ctx.Done():
		// 如果超时,返回搜索取消
		return errors.New("search canceled")
	case result := <-ch:
		if result.err != nil {
			return result.err
		}
		fmt.Println("Received:", result.record)
	}
	return nil
}
