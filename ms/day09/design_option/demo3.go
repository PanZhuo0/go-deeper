package main

import "fmt"

type config struct {
	name string
	age  int
}

type ConfigOption interface {
	apply(*config)
}

/* start*/
// 这个结构体用来实现ConfigOption这个接口,本质上这是一个辅助结构体
type funcObj struct {
	f func(*config)
}

func (f funcObj) apply(cfg *config) {
	f.f(cfg)
}

// 这个结构体对应的构造方法
func NewFuncObj(f func(*config)) funcObj {
	return funcObj{f: f}
}

/* end*/

func WithConfigName(name string) ConfigOption {
	f := func(cfg *config) {
		cfg.name = name
	}
	return NewFuncObj(f)
}

func WithConfigAge(age int) ConfigOption {
	f := func(cfg *config) {
		cfg.age = age
	}
	return NewFuncObj(f)
}

// 增加Config
func NewConfig(opts ...ConfigOption) *config {
	o := &config{}
	for _, opt := range opts {
		opt.apply(o)
	}
	return o
}

func main() {
	c := NewConfig(WithConfigName(`zhangsan`), WithConfigAge(100))
	fmt.Println(c)
}
