package main

import "fmt"

type Vehicle struct{}

// RUN方法违反了单一职责原则
func (v *Vehicle) Run(s string) {
	fmt.Println(s + "在公路上运行...")
}

func main() {
	v := Vehicle{}
	v.Run("摩托车")
	v.Run("汽车")
	v.Run("飞机")
}
