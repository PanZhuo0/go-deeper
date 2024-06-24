package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}

func search(nums []int, target int) int {
	// 二分查找算法
	return helper(nums, 0, len(nums)-1, target)
}

func helper(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[left] < nums[mid-1] {
		if target > nums[mid] {
			
		}
	} else {

	}
}
