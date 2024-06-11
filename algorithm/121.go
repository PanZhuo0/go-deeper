package main

import "fmt"

func main() {
	r := maxProfit([]int{2, 1, 2, 1, 0, 1, 2})
	fmt.Println(r)
}

// func maxProfit(prices []int) int {
// 	// 约束:只有先买入才能卖出
// 	// 动态规划
// 	// dp[i][0/1]: 从第0天到第i天,第i天持有/不持有用户能拥有最好的收益
// 	// 取值:dp[i][0]= max(dp[i-1][0],dp[i-1][1]+prices[i]) :昨天0 今天0 昨天 1 今天0
// 	// dp[i][1]= max(dp[i][1],dp[i][0]-prices[i]):昨天0 今天1 昨天1 今天1
// 	if len(prices) < 2 {
// 		return 0
// 	}
// 	dp := make([][]int, len(prices))
// 	for i := 0; i < len(dp); i++ {
// 		dp[i] = make([]int, 2)
// 	}
// 	dp[0][0] = 0
// 	dp[0][1] = -prices[0]
// 	for i := 1; i < len(prices); i++ {
// 		// 今天不持有
// 		if dp[i-1][0] > dp[i-1][1]+prices[i] {
// 			//昨天没有，今天不买
// 			dp[i][0] = dp[i-1][0]
// 		} else {
// 			//昨天有今天卖出
// 			dp[i][0] = dp[i-1][1] + prices[i]
// 		}
// 		// 今天持有
// 		if dp[i-1][1] > -prices[i] {
// 			//昨天持有，今天持有
// 			dp[i][1] = dp[i-1][1]
// 		} else {
// 			// 昨天没有，今天买入
// 			dp[i][1] = -prices[i]
// 		}
// 	}
// 	// 要获得收益、今天一定要抛出才有收益，因此返回dp[today][0]
// 	return dp[len(dp)-1][0]
// }

func maxProfit(prices []int) int {
	/*
		1.找出列表中最小的元素
		2.从最小的元素开始往后找到最大的元素,计算两者的差值
	*/
	if len(prices) < 2 {
		return 0
	}

	min := prices[0] //最小元素
	ret := 0         //最大差距
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i] //修改最小元素
		} else if prices[i]-min > ret {
			ret = prices[i] - min //修改最大差距
		}
	}
	return ret
}
