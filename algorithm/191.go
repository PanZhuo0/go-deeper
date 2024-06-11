package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(255))
}

func hammingWeight(n int) int {
	count := 0
	for n != 0 {
		n &= n - 1
		count++
	}
	return count
}
