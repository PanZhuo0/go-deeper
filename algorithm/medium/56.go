package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = [][]int{}
	a = append(a, []int{1, 3})
	a = append(a, []int{2, 6})
	a = append(a, []int{8, 10})
	a = append(a, []int{15, 18})
	fmt.Println(merge(a))
}

func merge(intervals [][]int) [][]int {
	// 先排序,按照升序排序
	f := func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	}
	sort.Slice(intervals, f)
	//这个时候我们就要考虑一下区间重叠的问题了
	//对于两个区间[x1, y1],[x2, y2]
	//总共可以分为两种情况，一种是重叠的，另一种是不重叠的
	//发生重叠时，x2 <= y1,这里因为是排序后的，所以x2 >= x1,这个时候只需要比较y2和y1的大小。还有一种情况是[x2,y2]是[x1,y1]的真子集，这个时候不需要更新
	//不发生重叠，那么 x2 > y1， 不需要更新

	//更新两个辅助start,end
	start, end := intervals[0][0], intervals[0][1]
	ans := [][]int{}
	for i := 1; i < len(intervals); i++ {
		// 如果第二个区间的起始位置大于第一个区间的终止位置，不需要合并区间，只需要添加到结果里
		if intervals[i][0] > end {
			ans = append(ans, []int{start, end})
			// 然后向后移动strat和end
			start = intervals[i][0]
			end = intervals[i][1]
		} else {
			end = max(intervals[i][1], end)
		}
	}
	//然后根据我们搜集到的end和start进行更新
	ans = append(ans, []int{start, end})
	return ans
}
