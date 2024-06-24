package svc

import (
	"user/internal/config"
	"user/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	u := model.NewUserModel(conn)
	return &ServiceContext{
		Config:    c,
		UserModel: u,
	}
}
