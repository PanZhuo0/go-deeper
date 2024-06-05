package main

import "fmt"

func main() {
	ret := mySqrt(10000000)
	fmt.Println(ret)
}
func mySqrt(x int) int {
	// 二分法
	min := 0
	max := x
	for max-min > 1 {
		middle := (min + max) / 2
		if middle*middle <= x && (middle+1)*(middle+1) > x {
			// 结束
			return middle
		} else if middle*middle > x {
			max = middle
		} else {
			min = middle
		}
	}
	return x
}
