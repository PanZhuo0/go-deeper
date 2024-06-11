package main

import "fmt"

type IceCreamMaker interface {
	Hello()
}

type Ben struct {
	id   int
	name string
}

func (b *Ben) Hello() {
	fmt.Println("Ben says: ", b.name)
}

type Jerry struct {
	name string
}

func (j *Jerry) Hello() {
	fmt.Println("Jerry says: ", j.name)
}

func main() {
	ben := &Ben{name: "Ben"}
	jerry := &Jerry{name: "jerry"}
	var maker IceCreamMaker = ben
	var loop0, loop1 func()
	// loop call
	loop0 = func() {
		maker = ben
		go loop1()
	}

	loop1 = func() {
		maker = jerry
		go loop0()
	}

	go loop0()

	for {
		maker.Hello()
	}
}
