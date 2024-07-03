package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2}, 1))
}

/*
	Method: 回溯法

	非常典型的使用回溯法的题目
*/

func combinationSum(candidates []int, target int) [][]int {
	sum := 0
	path := []int{}
	ans := [][]int{}
	var dfs func(startIdx int)
	dfs = func(startIdx int) {
		// 符合条件
		if sum == target {
			// append 2 ans
			ans = append(ans, append([]int(nil), path...))
			return
		}
		// 结束条件
		if startIdx == len(candidates) {
			return
		}
		if sum > target {
			return
		}
		// 尝试当前startIdx
		path = append(path, candidates[startIdx]) //push
		sum += candidates[startIdx]
		dfs(startIdx)
		path = path[:len(path)-1] //弹出
		sum -= candidates[startIdx]
		// 尝试下一个StartIdx
		dfs(startIdx + 1)
	}
	dfs(0)
	return ans
}
