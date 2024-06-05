package error

import "fmt"

type MyError struct {
	Msg  string
	File string
	Line int
}

// 实现error接口的Error方法
func (e *MyError) Error() string {
	return fmt.Sprintf("%s:%d:%s", e.File, e.Line, e.Msg)
}

// func test() error {
// 	return &MyError{"Something happened", "sever.go", 10}
// }

// func main() {
// 	err := test()
// 	switch err := err.(type) {
// 	case nil:
// 		// call succeeded,nothing to do
// 	case *MyError:
// 		fmt.Println("error occured on line:", err.Line)
// 	default:
// 		// unknown error
// 	}
// }
