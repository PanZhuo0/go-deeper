package server

import (
	v1 "bubblelist/api/bubblelist/v1"
	"bubblelist/internal/conf"
	"bubblelist/internal/service"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, todo *service.TodoService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			CustomMiddleWare(),
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
	// replace default ResponseEncoder
	opts = append(opts, http.ResponseEncoder(ResponseEncoder))
	// replace default ErrorEncoder
	opts = append(opts, http.ErrorEncoder(ErrorEncoder))
	srv := http.NewServer(opts...)
	v1.RegisterTodoHTTPServer(srv, todo)
	return srv
}

// 自定义的MiddleWare
func CustomMiddleWare() middleware.Middleware {
	return func(h middleware.Handler) middleware.Handler {
		// 内嵌使用闭包,实现before、after生命周期
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			// 执行handler之前做点事，before hook
			fmt.Println("Middleware:执行Handler之前")
			// token认证
			if tr, ok := transport.FromClientContext(ctx); ok {
				token := tr.RequestHeader().Get("token")
				fmt.Println("token:", token)
			}
			defer func() {
				// 执行之后做点事,after hook
				fmt.Println("Middleware:执行Handler之后")
			}()
			return h(ctx, req)
		}
	}
}
