package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}

// @lc code=start
/*
	1.对数据进行排序
	2.对结果进行合理的剪枝
*/
func permuteUnique(nums []int) [][]int {
	// 1.排序
	sort.Ints(nums)
	// 2.回溯+剪枝
	ans := [][]int{}          //结果集
	n := len(nums)            //n:数组的长度
	path := make([]int, n)    //走过的结果路径
	onPath := make([]bool, n) //某方案中数组中某个元素是否被使用

	var dfs func(int)
	dfs = func(i int) {
		// 如果是最深层
		if i == n {
			ans = append(ans, append([]int(nil), path...))
		}
		// 如果不是最深层
		for j, on := range onPath {
			if j > 0 {
				/* 去重的逻辑
				1.这个数已经在上一次的方案中使用过，
				2.并且这个数在方案的上层中未被使用
					如果不加上 &&onPaht[j-1]==false 会导致丢失一些可能的结果(出现了意料之外的叶子剪枝/树层剪枝)
				*/
				if nums[j] == nums[j-1] && onPath[j-1] == false {
					// 这个方案由于两者相同将会是重复的，剪枝
					continue
				}
			}
			if !on {
				path[i] = nums[j] //该方案:试着把第i层的元素设置为nums[j]
				onPath[j] = true  //将第j元素这个数据设置为使用过
				dfs(i + 1)        //考虑下一层
				onPath[j] = false //第i层为nums[j]的方案结束,将onPath[j]重置
			}
		}
	}
	dfs(0)
	return ans
}
