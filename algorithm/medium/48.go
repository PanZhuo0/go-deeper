package main

import "fmt"

func main() {
	a := make([][]int, 3)
	a[0] = []int{1, 2, 3}
	a[1] = []int{4, 5, 6}
	a[2] = []int{7, 8, 9}
	rotate(a)
}

func rotate(matrix [][]int) {
	// 两次线性变化矩阵
	// 上下变换,m[i]=m[len(m)-1-i]
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	fmt.Println(matrix)
	// 主对角线交换
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	fmt.Println(matrix)
}
