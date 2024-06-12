package main

import (
	"d3/d3"
	"fmt"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	// client
	b := d3.Book{
		Title:  "book_title",
		Price:  &wrapperspb.Int64Value{Value: 90},
		Price2: &wrapperspb.DoubleValue{Value: 90.99},
		Foo:    &wrapperspb.StringValue{Value: "Foo"},
	}
	// server
	if b.GetPrice() == nil {
		//if nil means default value
		fmt.Println("Need assign value:Price")
	} else {
		// get value
		fmt.Println(b.GetPrice().GetValue())
	}

	if b.GetFoo() == nil {
		fmt.Println("Need assign value:Foo")
	} else {
		fmt.Println(b.GetFoo().GetValue())
	}
}
