package main

import "fmt"

// 午餐接口,只需要工厂实现Cook函数，则属于能制造午餐的工厂
type Lunch interface {
	Cook()
}

// 米饭类
type Rice struct{}

// 米饭实现cook
func (r *Rice) Cook() {
	fmt.Println("it's rice!")
}

type Tomato struct{}

func (t *Tomato) Cook() {
	fmt.Println("it 's tomato!")
}

type LunchFactory interface {
	// 这就是抽象工厂？，就是返回值是抽象类?
	CreateFood() Lunch
	CreateVegetable() Lunch
}

type SimpleFactory struct{}

func NewSimpleFactory() LunchFactory {
	return &SimpleFactory{}
}

func (s *SimpleFactory) CreateFood() Lunch {
	return &Rice{}
}

func (s *SimpleFactory) CreateVegetable() Lunch {
	return &Tomato{}
}

/*
这个抽象工厂确实抽象啊，根本看不懂啊,先跳过，看看后面的模式，之后再回来学习
*/
func main() {
	factory := NewSimpleFactory()
	food := factory.CreateFood()
	food.Cook()
	vegetable := factory.CreateVegetable()
	vegetable.Cook()
}
