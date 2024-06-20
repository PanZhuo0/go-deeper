package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 使用GORM
func NewDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	// 迁移Schema
	db.AutoMigrate(&Shelf{}, &Book{})
	return db, nil
}

// 定义Model
// Shelf书架
type Shelf struct {
	ID       int64 `gorm:"primaryKey"`
	Theme    string
	Size     int64
	CreateAt time.Time
	UpdateAt time.Time
}

// Book 图书
type Book struct {
	ID       int64 `gorm:"primaryKey"`
	Author   string
	Title    string
	ShelfID  int64
	CreateAt time.Time
	UpdateAt time.Time
}

// 数据库操作
type bookstore struct {
	db *gorm.DB
}

const (
	defaultShelfSize = 100
)

// CreateShelf 创建书架
func (b *bookstore) CreateShelf(ctx context.Context, data Shelf) (*Shelf, error) {
	if len(data.Theme) <= 0 {
		return nil, errors.New("invalid theme")
	}
	if data.Size <= 0 {
		data.Size = defaultShelfSize
	}
	v := Shelf{Theme: data.Theme, Size: data.Size, CreateAt: time.Now(), UpdateAt: time.Now()}
	err := b.db.WithContext(ctx).Create(&v).Error
	return &v, err
}

// GetShelf 获取书架
func (b *bookstore) GetShelf(ctx context.Context, id int64) (*Shelf, error) {
	v := Shelf{}
	err := b.db.WithContext(ctx).First(&v, id).Error
	return &v, err
}

// ListShelves 书架列表
func (b *bookstore) ListShelves(ctx context.Context) ([]*Shelf, error) {
	var vl []*Shelf
	err := b.db.WithContext(ctx).Find(&vl).Error
	return vl, err
}

// DeleteShelv 删除书架
func (b *bookstore) DeleteShelv(ctx context.Context, id int64) error {
	n := b.db.WithContext(ctx).Delete(&Shelf{}, id).RowsAffected
	if n == 1 {
		return nil
	}
	return errors.New("#RowAffected!=1")
}

// 根据shelfID 查询结果
func (b *bookstore) GetBookListByShelfID(ctx context.Context, shelfID int64, pageSize int) ([]*Book, error) {
	var vl []*Book
	err := b.db.WithContext(ctx).Where("shelf_id=? and id > ?").Order("id asc").Limit(pageSize).Find(&vl).Error
	return vl, err
}
