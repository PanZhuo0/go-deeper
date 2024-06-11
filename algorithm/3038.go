package main

import "fmt"

func main() {
	fmt.Println(maxOperations([]int{3, 2, 1, 4, 5}))
}

func maxOperations(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	score := nums[0] + nums[1]
	nums = nums[2:]
	count := 1
	for len(nums) >= 2 {
		if nums[0]+nums[1] == score {
			nums = nums[2:]
			count++
		} else {
			return count
		}
	}
	return count
}
