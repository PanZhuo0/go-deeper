package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	// fmt.Println(threeSum([]int{0, 1, 1}))
}

/*
三数之和
1.三数之和可以变为sum-其中一个数退化为两数之和
2.如何确保相同的序列只出现一次(去重)

	由于结果不要求顺序，只要求不能重复,因此，通过对数组进行一次排序后再进行三数之和的操作即可
*/
func threeSum(nums []int) [][]int {
	// 1.排序数组
	sort.Ints(nums)
	// 2.获取结果
	for i := 0; i < len(nums)-2; i++ {

	}

}
