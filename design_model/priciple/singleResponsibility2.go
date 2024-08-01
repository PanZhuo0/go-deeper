package main

import "fmt"

// 1.遵守单一职责原则
// 2.改动很大
// 3.如何改进:直接改动Vehicle类,而不是新建类
func main() {
	r := RoadVehicle{}
	w := WaterVehicle{}
	a := AirVehicle{}
	a.run("飞机")
	w.run("摩托车")
	r.run("汽车")
}

type RoadVehicle struct{}

func (r *RoadVehicle) run(vehicle string) {
	fmt.Println(vehicle + "在公路上运行")
}

type AirVehicle struct{}

func (a *AirVehicle) run(vehicle string) {
	fmt.Println(vehicle + "在空中运行")
}

type WaterVehicle struct{}

func (a *WaterVehicle) run(vehicle string) {
	fmt.Println(vehicle + "在水中运行")
}
