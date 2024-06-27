package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
func permute(nums []int) [][]int {
	ans := [][]int{}
	n := len(nums)            //结果的长度
	path := make([]int, n)    //记录数据结果
	onPath := make([]bool, n) //用于记录某个特定方案中第index的数是否已被使用

	/*
		该函数用于遍历所有可能并将结果加入到结果集中
		DFS:deepth first search 深度优先搜索
		i代表深度
	*/
	var dfs func(int)
	dfs = func(i int) {
		// 是最深层
		if i == n {
			ans = append(ans, append([]int(nil), path...)) //将这个结果增加到结果集中
			return
		}
		// 不是最深层
		for j, on := range onPath {
			/*
				j:onPath的元素下标标识符
				on:onPath元素是否在某一阶段方案中已使用,T:已使用，F:未使用
			*/
			if !on {
				path[i] = nums[j] //试着在第i层放入nums[j]这个元素
				onPath[j] = true
				dfs(i + 1)        //放入完成，考虑i+1层应该放那个元素
				onPath[j] = false //将第i层放入nums[j]这个元素返回，接下来去考虑第i层放入nums[!j]元素的可能
			}
		}
	}
	dfs(0)
	return ans
}
