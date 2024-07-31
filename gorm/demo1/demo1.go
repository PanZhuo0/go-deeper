package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(mysql.Open(`root:123123@tcp(localhost:3006)/book`), &gorm.Config{
		// SkipDefaultTransaction: true, //是否跳过默认事务?
		// NamingStrategy: schema.NamingStrategy{
		// 	TablePrefix:   "tab_", //表名前缀
		// 	SingularTable: false,  //是否单数表名?
		// 	NoLowerCase:   true,   //不要小写转换?
		// }, //表名命名策略
		// Logger: ,设置日志

	})
	if err != nil {
		panic("连接数据库出错," + err.Error())
	}
}

func main() {
	// db.AutoMigrate()
}
