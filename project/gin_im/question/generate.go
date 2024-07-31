package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type Book struct {
	ID    int `db:"id"`
	Title int `db:"title"`
	Price int `db:"price"`
}

func main() {
	// 1.配置Genearator
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface, //生成默认的查询以及查询接口代码
	})
	// 2.新建一个GORM DB连接
	db, err := gorm.Open(mysql.Open(`root:123123@tcp(localhost:3306)/book?charset=utf8mb4&parseTime=True&loc=Local`)) //DB:book
	if err != nil {
		log.Fatalln(db)
	}
	// 3.让这个Genearator使用这个GORM DB 连接
	g.UseDB(db) //generator需要用上这个gormDB连接才行
	// 4.生成Model层
	g.GenerateModel("book")
	allTableModels := g.GenerateAllTable()
	// 5.生成query-CRUD API代码
	g.ApplyBasic(allTableModels...) //对该数据库中所有表格产生对应的CRUD API,这样不需要定义model
	// g.ApplyBasic(Book{}) //Table:book
	g.Execute()
}
