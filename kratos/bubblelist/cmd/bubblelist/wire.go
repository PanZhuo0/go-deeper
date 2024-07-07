//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"bubblelist/internal/biz"
	"bubblelist/internal/conf"
	"bubblelist/internal/data"
	"bubblelist/internal/server"
	"bubblelist/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	/*
		提供了
			server层的ProviderSet
			service层的ProviderSet
			biz层的ProviderSet
			data层的ProviderSet
	*/
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
