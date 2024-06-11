package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{-1, -1, -1, 1, 1, 1, 1}))
}

// 枚举法
func pivotIndex1(nums []int) int {
	for i := 0; i < len(nums); i++ {
		left := 0
		right := 0
		// 左侧求和 0->i-1
		for j := 0; j < i; j++ {
			left += nums[j]
		}
		// 右侧求和
		for j := i + 1; j < len(nums); j++ {
			right += nums[j]
		}
		if left == right {
			return i
		}
	}
	return -1
}

func pivotIndex(nums []int) int {
	// 指针法
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	total := 0
	for i := 0; i < len(nums); i++ {
		sum -= nums[i]
		if total == sum {
			return i
		}
		total += nums[i]
	}
	// 找不到
	return -1
}
