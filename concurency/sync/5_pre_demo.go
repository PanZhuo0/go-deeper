package main

import "fmt"

func main() {
	fmt.Println("Go")
	var a, b int
	var err1, err2 error
	go func() {
		// call rpc1
		a, err1 = xxx()
		// if err can not return

	}()

	go func() {
		// call rpc2
		b, err2 = xx()
	}()
}
