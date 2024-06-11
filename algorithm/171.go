package main

import "fmt"

func main() {
	fmt.Println(titleToNumber("ZY"))
}

func titleToNumber(columnTitle string) int {
	sum := 0
	weight := 1
	for i := len(columnTitle) - 1; i >= 0; i-- {
		// 从尾部开始计算
		sum += weight * int(columnTitle[i]-64)
		weight *= 26
	}
	return sum
}
