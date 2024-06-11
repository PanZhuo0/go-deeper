package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{0, 1, 2, 2, 2, 2, 2, 3, 3, 3}))
}
func removeDuplicates(nums []int) int {
	// 双指针法
	i, j := 0, 1 //i慢指针，j快指针
	flag := false
	for j < len(nums) {
		// 第一种情况，两者不同,i++,j++
		// 第二种情况,两种相同,j++,一直j++直到nums[i]!=nums[j]
		// 同时,i++,nums[i]=nums[j]
		if nums[i] != nums[j] && flag {
			i++
			nums[i] = nums[j]
		} else if nums[i] != nums[j] {
			flag = false
			i++
		} else {
			flag = true //需要调整
		}
		j++
	}
	return i + 1
}
