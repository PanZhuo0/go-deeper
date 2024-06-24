package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{1, 1, 1, 1}, 0))
}

func threeSumClosest(nums []int, target int) int {
	// 1.排序
	sort.Ints(nums)
	mindiff := 100000
	sum := 10000
	// 2.找出结果
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[i]+nums[left]+nums[right] == target {
				return target
			} else if nums[i]+nums[left]+nums[right] > target {
				d := diff(nums[i]+nums[left]+nums[right], target)
				if d < mindiff {
					mindiff = d
					sum = nums[i] + nums[left] + nums[right]
				}
				right--
			} else if nums[i]+nums[left]+nums[right] < target {
				d := diff(nums[i]+nums[left]+nums[right], target)
				if d < mindiff {
					mindiff = d
					sum = nums[i] + nums[left] + nums[right]
				}
				left++
			}
		}
	}
	return sum
}

func diff(x, y int) int {
	d := x - y
	if d < 0 {
		return -d
	}
	return d
}
