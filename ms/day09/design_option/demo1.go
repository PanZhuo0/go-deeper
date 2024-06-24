package main

import "fmt"

type DoSomethingOption struct {
	a string
	b int
	c bool
	y Info
	//....
}
type Info struct {
	Addr string
	City string
}

var Default = &DoSomethingOption{
	a: `zhangsan`,
	b: 20,
	c: false,
}

/*
// Traditional
func NewDoSomethingOption(a string, b int, c bool) *DoSomethingOption {
	return &DoSomethingOption{
		a: a,
		b: b,
		c: c,
	}
}
*/

/*
	存在的问题:

1.如果DoSomethingOption有很多字段，那构造函数就需要十几个参数,如何为这些配置项指定默认值呢?
2.DoSomethingOption随着业务不断增加字段后，我们又需要同步变更它的NewDoSomethingOption方法
(变更了构造函数又会影响既有代码)
*/

/* Option Pattern */
type OptionFunc func(*DoSomethingOption)

// WithXxx 是Option Pattern 中约定成俗的函数命名格式
func WithB(b int) OptionFunc {
	return func(o *DoSomethingOption) {
		o.b = b
	}
}

// WithY 实现将Info类型的数据赋值给Option中的y字段
func WithY(info Info) OptionFunc {
	return func(o *DoSomethingOption) {
		o.y = info
	}
}

func WithC(c bool) OptionFunc {
	// 闭包函数、函数使用到了函数外的c这个值
	return func(o *DoSomethingOption) {
		o.c = c
	}
}

func NewDoSomethingOption(a string, opts ...OptionFunc) *DoSomethingOption {
	o := Default
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func main() {
	// option pattern
	opt := NewDoSomethingOption(`zhangsan`, WithB(12), WithC(true), WithB(20), WithB(30), WithY(Info{Addr: `天安门`, City: `北京`}))
	opt.y.Addr = `故宫` //存在的问题,可以直接修改值,kit包使用者可以不必使用With
	fmt.Println(opt)
}
