package main

import (
	"fmt"
)

func main() {

	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}

func search(nums []int, target int) int {
	// 二分查找算法
	return helper(nums, 0, len(nums)-1, target)
}

func helper(nums []int, left, right, target int) int {
	mid := (left + right) / 2
	if nums[right] == target {
		return right
	}
	if nums[mid] == target {
		return mid
	}
	if mid == left || mid == right {
		return -1
	}
	if nums[left] <= nums[mid-1] {
		// 左侧有序
		if nums[left] <= target && target <= nums[mid-1] {
			return helper(nums, left, mid-1, target)
		} else {
			return helper(nums, mid, right, target)
		}
	} else {
		// 右侧有序
		if nums[mid+1] <= target && target <= nums[right] {
			return helper(nums, mid+1, right, target)
		}
		return helper(nums, left, mid-1, target)
	}
}
