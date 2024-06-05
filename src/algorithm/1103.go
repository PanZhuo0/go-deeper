package main

import "fmt"

func main() {
	ret := distributeCandies(7, 4)
	fmt.Println(ret)
}

func distributeCandies(candies int, num_people int) []int {
	// 循环次数=sqrt(n): 1+2+3+...+n=2**n-1  所以n颗糖需要 sqrt(n)次分配
	n := 1 //每次分的糖果
	j := 0
	ret := make([]int, num_people)
	for candies-n > 0 {
		ret[j%num_people] += n //给糖果
		candies -= n
		//下一位
		j++
		n++
	}
	// 处理最后一次分配
	if candies > 0 {
		ret[j%num_people] += candies
	}
	return ret
}
