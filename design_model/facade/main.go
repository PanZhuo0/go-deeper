package main

import "fmt"

// 以车子Car为例,车子的其他部分都不需要对外显示，只需要显示车子本身就行,Facade
type CarModel struct{}
type CarEngine struct{}
type CarBody struct{}

// 构造方法
func NewCarModel() *CarModel {
	return &CarModel{}
}

func NewCarEngine() *CarEngine {
	return &CarEngine{}
}

func NewCarBody() *CarBody {
	return &CarBody{}
}

// set方法
func (c *CarModel) SetCarModel() {
	fmt.Println("CarModel setting car model....")
}

func (c *CarEngine) SetCarEngine() {
	fmt.Println("CarEngine setting car Engine....")
}
func (c *CarBody) SetCarBody() {
	fmt.Println("CarBody setting car body....")
}

type CarFacade struct {
	model  CarModel
	engine CarEngine
	body   CarBody
}

// 车的外观
func NewCarFacade() *CarFacade {
	return &CarFacade{
		model:  CarModel{},
		engine: CarEngine{},
		body:   CarBody{},
	}
}

func (c *CarFacade) CreateCompleteCar() {
	c.model.SetCarModel()
	c.engine.SetCarEngine()
	c.body.SetCarBody()
}

// 外观模式Facade Pattern
func main() {
	c := NewCarFacade()
	c.CreateCompleteCar()
	fmt.Println(c)
}
