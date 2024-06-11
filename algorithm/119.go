package main

import "fmt"

func main() {
	fmt.Println(getRow(5))
}

func getRow(rowIndex int) []int {
	// 杨辉三角第n行第k个数= C(n-1,k-1)  从n-1个元素中取出k-1个元素的总方案数
	ans := make([]int, rowIndex+1)
	ans[0] = 1
	ans[rowIndex] = 1
	for i := 1; i < rowIndex; i++ {
		ans[i] = combinations(rowIndex, i)
	}
	return ans
}

// C(n-1,k-1)
func combinations(n, m int) int {
	ret := 1
	for i := 0; i < m; i++ {
		ret *= (n - i)
		ret /= (i + 1)
	}
	return ret
}
