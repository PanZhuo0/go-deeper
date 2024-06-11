package main

import "fmt"

func main() {
	// OddTimesNum2([]int{})
	test()
}

// 找出数组中仅有的两个出现奇数次的数
func OddTimesNum2(a []int) {
	a = []int{1, 1, 2, 2, 3, 3, 3, 4, 4, 5, 6, 6} //only 3 & 5 ocurr odd times
	eor := 0
	for i := 0; i < len(a); i++ {
		eor ^= a[i]
	}
	fmt.Println(eor)
	// find the special 1 bit of eor(rightest)
	rightOne := eor & (^eor + 1)
	fmt.Println(rightOne)
	// find one of oddTimesNumbers
	for i := 0; i < len(a); i++ {
	}
}

func test() {
	a := []int{1, 2, 3, 3, 4, 5}
	fmt.Println(len(a))
	fmt.Println(a[:len(a)-1])
}
