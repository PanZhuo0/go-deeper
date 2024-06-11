package main

import "fmt"

func main() {
	a := []int{1, 2, 2, 2, 2, 2, 3, 33, 1, 1, 22, 3, 3, 3, 123, 12, 3, 123, 1, 23, 12, 312, 3, 12, 3}
	fmt.Println(Bubble(a))
}

func Bubble(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}
