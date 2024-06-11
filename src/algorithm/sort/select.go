package main

import "fmt"

func main() {
	a := []int{1, 2, 312, 3, 12, 31, 23, 123, 1, 3, 12, 31, 23, 1, 23, 12, 31, 23, 12, 31, 23, 12, 31, 23, 12, 3, 123, 12, 3, 12, 31, 23, 12, 31, 23, 123}
	fmt.Println(Select(a))
}

func Select(a []int) []int {
	// a[0]=4
	// a[1]=5
	// a[2]=3
	for i := 1; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[j] > a[i] { //a[0]=3>a[1]=5 //交换位置
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}
