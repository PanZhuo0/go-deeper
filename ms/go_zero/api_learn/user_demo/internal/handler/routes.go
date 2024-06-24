// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"user_demo/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/signup",
				Handler: SignupHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api"),
	)
}