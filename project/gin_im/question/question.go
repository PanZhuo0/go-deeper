package main

import (
	"context"
	"fmt"

	"log"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// ini_1()
	// ini_2()
	// gen_1()
}

func genID_1() {
	// genID_1 学习怎么产生消息的ID
}

func gen_1() {
	// gen_1 学习如何使用gorm Gen
	// 1.generate.go 用于生成对应的query 文件,可以供我们直接使用
	// 2.使用GORM-Gen生成的CRUD-API
	db, err := gorm.Open(mysql.Open(`root:123123@tcp(localhost:3306)/book?charset=utf8mb4&parseTime=True&loc=Local`))
	if err != nil {
		log.Fatalln(err)
	}
	query.SetDefault(db) //这一步是必须的,需要绑定对应的数据库
	query := query.Q
	query.Book.WithContext(context.Background()).Create(&model.Book{
		Title: "Go",
		Price: 99,
	})
}

type DataBase struct {
	DB         string `ini:"DB"`
	DBHost     string `ini:"DBHost"`
	DBUser     string `ini:"DBUser"`
	DBPort     string `ini:"DBPort"`
	DBPassword string `ini:"DBPassword"`
}

type Redis struct {
	RedisDB   string `ini:"RedisDB"`
	RedisAddr string `ini:"RedisAddr"`
}

type MongoDB struct {
	MongDBName      string `ini:"MongDBName"`
	MongoDBAddr     string `ini:"MongoDBAddr"`
	MongoDBPort     string `ini:"MongoDBPort"`
	MongoDBPassword string `ini:"MongoDBPassword"`
}

type Config struct {
	DataBase `ini:"DataBase"`
	Redis    `ini:"Redis"`
	MongoDB  `ini:"MongoDB"`
}

func ini_2() {
	// 2.结构体与分区(ini中的section)双向映射、也就是可以从配置文件中的分区转移到结构体、也可以从结构体转移到分区
	cfg := new(Config)
	err := ini.MapTo(cfg, "./config/config.ini")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)
}

func ini_1() {
	// 1.不知道怎么解析ini 文件1
	/* ini具有分区(Section)、键(Key)、键值(Value) */
	cfg, err := ini.Load("./config/config.ini")
	if err != nil {
		log.Fatal("解析配置文件失败,err:", err)
	}
	fmt.Println(cfg)
	sec, err := cfg.GetSection("DataBase")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sec.GetKey("DBHost"))
}
