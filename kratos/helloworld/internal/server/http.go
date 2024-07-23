package server

import (
	"context"
	"fmt"
	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/conf"
	"helloworld/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// 自定义MiddleWare
/* 扩展知识:
为什么要用这种闭包的方式写Middleware，而不直接定义一个方法?

为什么用这个
	func FuncXxx() middleware.Middleware {
		return func(h middleware.Handler) middleware.Handler {

		}
	}

而不用这个？
	func(h middleware.Handler) FuncXxx middleware.Handler {

	}

答:使用闭包允许调用者拥有更高的扩展性


eg:
	func MiddleWare(opts ...any) middleware.Middleware {
		return func(h middleware.Handler) middleware.Handler {
			// switch depend on opts
			switch opts{
				...
			}
		}
	}

*/

// 自定义的MiddleWare
func CustomMiddleWare() middleware.Middleware {
	return func(h middleware.Handler) middleware.Handler {
		// 内嵌使用闭包,实现before、after生命周期
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			// 执行handler之前做点事，before hook
			fmt.Println("Middleware:执行Handler之前")
			defer func() {
				// 执行之后做点事,after hook
				fmt.Println("Middleware:执行Handler之后")
			}()
			return h(ctx, req)
		}
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			CustomMiddleWare(),
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	return srv
}
