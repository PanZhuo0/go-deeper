package main

import "fmt"

func main() {
	a := new(int)
	*a = 10
	fmt.Println(a)
}
