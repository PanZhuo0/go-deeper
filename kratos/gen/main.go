package main

import (
	"context"
	"fmt"
	"gen/dal/query"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const MySQLDSN = "root:123123@tcp(127.0.0.1:3306)/gen?charset=utf8mb4&parseTime=True"

func connectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	return db
}

func main() {
	fmt.Println("gen demo start...")
	db := connectDB(MySQLDSN)

	// 重要！
	query.SetDefault(db)

	// CRUD
	// Create
	/* 	b1 := model.Book{
	   		Title:       "Book",
	   		Author:      "zhangsan",
	   		Price:       100,
	   		PublishDate: time.Now(),
	   	}
	   	err := query.Book.WithContext(context.Background()).Create(&b1)
	   	if err != nil {
	   		fmt.Printf("create book failed,err:%v\n", err)
	   		return
	   	}
	   	fmt.Printf("b1:%#v\n", b1) */

	// Retrieval
	/* 	ret, err := query.Book.
	   		WithContext(context.Background()).
	   		Where(query.Book.ID.Eq(1)).
	   		First()
	   	fmt.Printf("%#v\n", ret)
	   	fmt.Println(err)

	   	// Retrieval 2
	   	book := query.Book
	   	b, err := book.
	   		WithContext(context.Background()).
	   		Where(book.ID.Eq(1)).
	   		First()
	   	fmt.Printf("%#v\n", b)
	   	fmt.Println(err) */

	// Update
	/* 	ret, err := query.Book.
	   		WithContext(context.Background()).
	   		Where(query.Book.ID.Eq(1)).
	   		Update(query.Book.Price, 200)
	   	if err != nil {
	   		fmt.Printf("update book failed,err:%v\n", err)
	   		return
	   	}
	   	fmt.Println(ret.RowsAffected) */

	// Delete
	/* 	b3 := model.Book{
	   		ID: 1,
	   	}
	   	ret, err := query.Book.
	   		WithContext(context.Background()).
	   		Delete(&b3)
	   	if err != nil {
	   		fmt.Printf("delete failed,err:%#v", err)
	   		return
	   	}
	   	fmt.Println(ret.RowsAffected) */

	// Delete 2
	// ret, err := query.Book.
	// 	WithContext(context.Background()).
	// 	Where(query.Book.ID.Eq(2)).
	// 	Delete()
	// if err != nil {
	// 	fmt.Printf("delete failed,err:%#v", err)
	// 	return
	// }
	// fmt.Println(ret.RowsAffected)

	// Use custom GORM func
	books, err := query.Book.WithContext(context.Background()).GetBooksByAuthor("zhangsan")
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}
	for k, v := range books {
		fmt.Printf("%d: book:%#v\n", k, v)
	}

	// query.Book.WithContext(context.Background()).GetByID()
	// query.Book.WithContext(context.Background()).GetByIDReturnMap()
}
