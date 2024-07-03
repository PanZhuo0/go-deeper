package main

import "fmt"

/*
动态规划
	dp[i][j]:代表从左上方到[i,j]这个节点的最小路径和
*/
func main() {
	fmt.Println(minPathSum([][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
	}))
}

func minPathSum(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			// DP[i][j]:从(0,0)到达点(i,j)的最短路径
			if i != 0 || j != 0 {
				// 第一行的元素
				if i == 0 {
					dp[i][j] = dp[i][j-1] + grid[i][j]
					continue
				}
				// 第一列的元素
				if j == 0 {
					dp[i][j] = dp[i-1][j] + grid[i][j]
					continue
				}
				// 其他位置的元素的DP[i][j]
				if dp[i-1][j] < dp[i][j-1] {
					dp[i][j] = grid[i][j] + dp[i-1][j]
				} else {
					dp[i][j] = grid[i][j] + dp[i][j-1]
				}
			}
		}
	}
	return dp[row-1][col-1]
}
