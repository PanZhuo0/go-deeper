package svc

import (
	"user/internal/config"
	"user/internal/middleware"
	"user/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	Cost      rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	u := model.NewUserModel(conn, c.CacheRedis)
	return &ServiceContext{
		Config:    c,
		UserModel: u,
		Cost:      middleware.NewCostMiddleware().Handle,
	}
}
