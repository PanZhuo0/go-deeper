package error

import (
	"fmt"
)

/*
错误处理---微服务中goroutine的使用
需要处理好服务中goroutine可能出现的panic
*/

// func main() {
// 	// go func() {
// 	// 	fmt.Println("Hello")
// 	// 	panic("panic happened")
// 	// }()
// 	Go(func() {
// 		panic("panic happened")
// 	})
// 	time.Sleep(time.Second)
// }

func Go(x func()) {
	go func() {
		// if panic
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		// call x
		x()
	}()
}
