package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(nums, 8))
}

/*
	利用Go中内置的searchInts方法

// SearchInts searches for x in a sorted slice of ints and returns the index
// as specified by [Search]. The return value is the index to insert x if x is
// not present (it could be len(a)).
// The slice must be sorted in ascending order.

	func SearchInts(a []int, x int) int {
		return Search(len(a), func(i int) bool { return a[i] >= x })
	}
*/

func searchRange(nums []int, target int) []int {
	leftMost := sort.SearchInts(nums, target)
	if leftMost == len(nums) || nums[leftMost] != target {
		// 找不到
		return []int{-1, -1}
	}
	rightMost := sort.SearchInts(nums, target+1) - 1
	return []int{leftMost, rightMost}
}
