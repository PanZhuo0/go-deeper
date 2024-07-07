package data

import (
	"bubblelist/internal/biz"
	"bubblelist/internal/conf"
	"errors"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDB, NewData, NewTodoRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
	// db sqlx.DB
}

// NewData .
func NewData(db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	/*
		如果在这里直接使用gorm连接DB，就不符合IOC要求
		db, err := gorm.Open(mysql.Open(c.Database.Dsn), &gorm.Config{})
	*/
	// 应该使用依赖注入,将依赖通过参数传进来
	return &Data{db: db}, cleanup, nil
}

func NewDB(c *conf.Data) (*gorm.DB, error) {
	// 根据配置文件中指定的driver
	var db *gorm.DB
	var err error
	switch strings.ToLower(c.Database.Driver) {
	case "mysql":
		db, err = gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(c.Database.Source), &gorm.Config{})
	default:
		return nil, errors.New("invalid driver")
	}
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&biz.Todo{}) //自动建表
	return db, err
}
