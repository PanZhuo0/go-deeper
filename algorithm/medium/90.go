package main

import "sort"

/* 回溯算法

 */
func main() {

}

var (
	path []int
	res  [][]int
)

func subsetsWithDup(nums []int) [][]int {
	path, res = make([]int, 0, len(nums)), make([][]int, 0)
	sort.Ints(nums)
	dfs(nums, 0)
	return res
}

func dfs(nums []int, start int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	res = append(res, tmp)

	for i := start; i < len(nums); i++ {
		if i != start && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		dfs(nums, i+1)
		path = path[:len(path)-1]
	}
}
