package main

import "context"

// A message processes parameter and returns the result on responseChan
// ctx is places in a  stuct,but this is ok to do
type message struct {
	responseChan chan<- int
	parameter    string
	ctx          context.Context
}

func main() {
	// with Value
	context.WithValue()
	// with Timeout
	context.WithTimeout()
}
