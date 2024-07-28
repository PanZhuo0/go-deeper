package main

import "fmt"

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	pace uint8
}

func (c cat) move() {
	fmt.Printf(c.name + " moving...\n")
}

func (c cat) eat(f string) {
	fmt.Printf(c.name+" eating %v\n", f)
}
func main() {
	c1 := new(cat)
	c1.name = "tom"
	c1.pace = 10
	fmt.Printf("%T\n", c1)
	fmt.Printf("%#v", c1)
	c1.move()
	c1.eat("food")
}
