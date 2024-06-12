package main

import "fmt"

// Local call demo
func main() {
	// call add function
	sum := add(1, 2)
	fmt.Println(sum)
}

func add(x, y int) int {
	return x + y
}
