package main

import "fmt"

func main() {
	ret := distributeCandies([]int{1, 1, 2, 2, 3, 3})
	fmt.Println(ret)
}
func distributeCandies(candyType []int) int {
	// 使用map 列出种类
	m := make(map[int]bool, len(candyType))
	for i := 0; i < len(candyType); i++ {
		m[candyType[i]] = true
	}
	// 比较x 与 2/len(candyType) 的长度、取较小值
	if len(m) <= len(candyType)/2 {
		return len(m)
	} else {
		return len(candyType) / 2
	}
}
