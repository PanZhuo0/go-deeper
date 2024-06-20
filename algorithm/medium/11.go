package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	/*
		left:		水箱的左侧
		right:		水箱的右侧
		base:		水箱的底部长度
		area:		形成的水箱的容积
	*/
	left := 0
	right := len(height) - 1
	base := len(height) - 1
	area := 0
	for left != right {
		sum := base * min(height[left], height[right]) //计算当前容积
		// 更新水箱底部长度、更新水箱的左侧/右侧
		base--
		if sum > area {
			area = sum
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return area
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
