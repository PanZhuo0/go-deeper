package main

import "fmt"

type Resturant interface {
	GetFood()
}

type ShaXian struct{}
type Hualaishi struct{}

func (s ShaXian) GetFood() {
	fmt.Println("沙县大酒店!欢迎来吃")
}

func (h Hualaishi) GetFood() {
	fmt.Println("华莱士，来的快，去的也快!")
}

func NewRestaurant(name string) Resturant {
	switch name {
	case "s":
		return &ShaXian{}
	case "h":
		return &Hualaishi{}
	}
	return nil
}

/* 简单工厂模式:
好处:
	调用者只需要指明厂商/工厂名字就可以获取对应的工厂对象，可执行GETxxx方法

Go中不适合用这个模式
	1.当我们NewXXX的时候Go不会对你输入的参数(工厂名字)进行判断，可能导致NewXXX失败,Go不会进行编译时检查
	2.需要传入实例key string。 这个key需要多余的维护，并且容易写错。

// B站评论
工厂模式不适合在go里使用，需要哪个实例，建议直接从实例去New出来。基于以下原因，不适合:

- 工厂模式，需要传入实例key string。 这个key需要多余的维护，并且容易写错。这个key一旦写错，无法享受go的编译前错误，将他报出来，而只有运行期才可以报错。
- 和通过pkg.Instance的调用方式相比，没有管理上，性能上的优势，仅仅是为了模仿java而衍生出来的东西。
- 继承了一些接口的缺点。如果一个功能，可以抛开接口来实现，那么就不要使用接口。接口在编程逻辑学里很美，但是在实际开发中，有以下问题:
- 无法通过调用的位置，点出实例的实现位置。而会跳到接口声明，进而凭感觉去挑选哪个实例是你要看的。
- 实例在编写时，容易因为返回值漏写一个error，或者参数写偏，而造成实例未实现该接口。这个未实现在实例页不会即时抛错，等到你想要寻早这个实例为什么没有实现该接口时，能找到你暴跳如雷。
- 业务侧容易做需求迭代。一个接口很难在声明初期就预设出他的所有限制要求，而任何的变动，都依赖所有实例做变动，它跨越了作者，时间，版本。
*/
func main() {
	NewRestaurant("s").GetFood()
	NewRestaurant("h").GetFood()
}
