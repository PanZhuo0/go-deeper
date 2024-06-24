package main

import "fmt"

func main() {
	n := []int{1, 3, 2}
	nextPermutation(n)
	fmt.Println(n)
}

func nextPermutation(nums []int) {
	// 1.找到left和right
	// 2.交换left、right
	// 3.将left后的下标的所有数据倒序
	if len(nums) == 1 {
		return
	}
	left := len(nums) - 1
	right := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		// 完全倒序
		if i == 0 && nums[i] > nums[left] {
			for i := 0; i < len(nums)/2; i++ {
				nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
			}
			return
		}
		// nums[i]<nums[left]
		if nums[i] < nums[left] {
			left--
			break
		}
		left--
	}
	// 找到right
	for i := len(nums) - 1; i > left; i-- {
		if nums[i] > nums[left] {
			right = i
			break
		}
	}
	fmt.Println(left)
	fmt.Println(right)
	// 交换位置
	nums[right], nums[left] = nums[left], nums[right]

	fmt.Println(nums)
	// left下标后面的所有元素倒序
	for i := left + 1; i < (len(nums)+left+1)/2; i++ {
		nums[i], nums[len(nums)-i+left] = nums[len(nums)-i+left], nums[i]
	}
	return
}
