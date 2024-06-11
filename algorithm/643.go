package main

import "fmt"

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}

func findMaxAverage(nums []int, k int) float64 {
	if len(nums) < k {
		return 0
	}
	var max_avg float64
	var sum float64
	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}
	max_avg = sum / float64(k)
	for i := 1; i <= len(nums)-k; i++ {
		// sliding windows
		sum = sum - float64(nums[i-1])
		sum = sum + float64(nums[i+k-1])
		if max_avg < (sum / float64(k)) {
			max_avg = sum / float64(k)
		}
		fmt.Println(sum)
	}
	return max_avg
}
