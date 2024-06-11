package error

import (
	"errors"
)

// 模拟源码(源码中是一个结构体对象、这里是string类型别名为errorString)
type errorString string

// 模拟源码
func (e errorString) Error() string {
	return string(e)
}

func New(text string) error {
	// 源码中这里返回的是&errorString
	return errorString(text)
}

// 模拟源码的哨兵级别的错误
var ErrNamedType = New("EOF")
var ErrStructType = errors.New("EOF")

// func main() {
// 	// 模拟源码的返回的是string的值
// 	if ErrNamedType == New("EOF") {
// 		fmt.Println("Name Type Error") //这里会输出
// 	}
// 	// 源码中的errors.New()返回的是errorString对象的指针
// 	if ErrStructType == errors.New("EOF") {
// 		fmt.Println("Struct Type Error") //这里不会输出,因为指向的不是同一个对象
// 	}
// }

/*为什么源码要这样定义?
1.使用对象地址而不是值来定义,可以避免工程中出现用户不经意定义了相同的errorString(值相同),导致错误的认为是同一个error
2.即便定义相同的errorString也不会是同一个Error
*/
