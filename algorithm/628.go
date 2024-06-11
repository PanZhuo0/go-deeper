package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumProduct([]int{-1, -2, -3}))
}

func maximumProduct2(nums []int) int {
	sort.Ints(nums)
	// 三个正数 / 一个正数 两个负数
	n1 := nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]
	n2 := nums[len(nums)-1] * nums[0] * nums[1]
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}

func maximumProduct(nums []int) int {
	if len(nums) == 3 {
		return nums[0] * nums[1] * nums[2]
	}
	// 只关心最大的三个值和最小的两个值
	min1, min2 := nums[0], nums[1]
	max1, max2, max3 := nums[0], nums[1], nums[2]

	for i := 0; i < len(nums); i++ {
		if nums[i] < min1 {
			// 小于最小值
			min1, min2 = nums[i], min1
		} else if nums[i] < min2 {
			// 小于第二小值，大于最小值
			min2 = nums[i]
		}

		if nums[i] > max1 {
			// 大于最大值
			max1, max2, max3 = nums[i], max1, max2
		} else if nums[i] > max2 {
			// 小于最大值，大于次大值
			max2, max3 = nums[i], max2
		} else if nums[i] > max3 {
			// 小于最大，次大值，大于次次大值
			max3 = nums[i]
		}
	}

	fmt.Println(max1, max2, max3)
	fmt.Println(min1, min2)
	n1 := max1 * max2 * max3
	n2 := max1 * min1 * min2
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}
