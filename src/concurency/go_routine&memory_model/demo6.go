package main

import "fmt"

type A struct {
}

func main() {
	a := 1
	fmt.Println(a)

	// 在goroutine1中重新赋值了一个
	b := new(A) //一个指针默认是64bit 需要小心是否是机器字长
	b = new(A)

	// 问题:一个32bit的系统，获取一个64bit的地址，可不可能读取到一半地址？
	// 原子安全问题:如果交换一个指针对应的地址(64bit)但是机器字长=32bit，怎么解决？

	go func() {
		fmt.Println(b) //data rece ,atomic   --- go build -race
	}()
}
