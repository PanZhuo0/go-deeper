package middleware

import (
	"bytes"
	"fmt"
	"net/http"
)

/*
核心逻辑

将ResponseWrtier改为自定义的bodyCopy(该结构体实现了RepsonseWriter,并自带一个bytes.Buffer结构体)
*/
type bodyCopy struct {
	http.ResponseWriter               //结构体嵌入接口类型
	body                *bytes.Buffer //用来记录响应体内容
}

// 实现对应的write方法
func (bc bodyCopy) Write(b []byte) (int, error) {
	// 将要响应的数据写入到body
	bc.body.Write(b)
	// 将要响应的数据写入HTTP 响应
	return bc.ResponseWriter.Write(b)
}

// 复制请求的响应体
func CopyResp(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 处理请求前
		body := bytes.NewBuffer(make([]byte, 0))
		var bc = bodyCopy{ResponseWriter: w, body: body}
		next(bc, r)
		// 处理请求后:把响应结果打印出来
		// 怎么把这个bc.body 读完？
		fmt.Printf("----> req:%v  resp:%v\n", r.URL, bc.body.String())
	}
}

/*
如何在中间件调用其他的服务---->闭包
func WithAnotherService(s *AnotherService) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("X-middleware", s.GetToken())
			next(w, r)
		}
	}
}
*/

