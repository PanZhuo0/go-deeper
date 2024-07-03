package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
}

func combine(n int, k int) [][]int {
	ans := [][]int{}
	tmp := []int{}
	var dfs func(x, y int)
	dfs = func(layer, start int) {
		if layer == k {
			// fmt.Println(tmp) need append(tmp ...)
			ans = append(ans, append([]int(nil), tmp...))
			return
		}
		// through from start
		for i := start; i < n; i++ {
			tmp = append(tmp, i+1) //push
			dfs(layer+1, i+1)
			tmp = tmp[:len(tmp)-1] //pop
		}
	}
	dfs(0, 0)
	return ans
}

func combine1(n int, k int) [][]int {
	path := []int{}
	var dfs func(int)
	ans := make([][]int, 0)
	dfs = func(cur int) {
		// 当前已用数据+可用的其他数字<总需要的位数,已经没有可用的结果了
		if len(path)+(n-cur+1) < k {
			return
		}
		// 长度到达k
		if len(path) == k {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		// push
		path = append(path, cur)
		dfs(cur + 1) //考虑当前位置
		// 回溯完后 pop
		path = path[:len(path)-1]
		dfs(cur + 1) //不考虑当前位置
	}
	dfs(1)
	return ans
}
