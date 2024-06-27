// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	bookFieldNames          = builder.RawFieldNames(&Book{})
	bookRows                = strings.Join(bookFieldNames, ",")
	bookRowsExpectAutoSet   = strings.Join(stringx.Remove(bookFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	bookRowsWithPlaceHolder = strings.Join(stringx.Remove(bookFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheBookBookIdPrefix = "cache:book:book:id:"
)

type (
	bookModel interface {
		Insert(ctx context.Context, data *Book) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Book, error)
		Update(ctx context.Context, data *Book) error
		Delete(ctx context.Context, id int64) error
	}

	defaultBookModel struct {
		sqlc.CachedConn
		table string
	}

	Book struct {
		Id    int64          `db:"id"`
		Title sql.NullString `db:"title"`
		Price sql.NullInt64  `db:"price"`
	}
)

func newBookModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultBookModel {
	return &defaultBookModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`book`",
	}
}

func (m *defaultBookModel) Delete(ctx context.Context, id int64) error {
	bookBookIdKey := fmt.Sprintf("%s%v", cacheBookBookIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, bookBookIdKey)
	return err
}

func (m *defaultBookModel) FindOne(ctx context.Context, id int64) (*Book, error) {
	bookBookIdKey := fmt.Sprintf("%s%v", cacheBookBookIdPrefix, id)
	var resp Book
	err := m.QueryRowCtx(ctx, &resp, bookBookIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", bookRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBookModel) Insert(ctx context.Context, data *Book) (sql.Result, error) {
	bookBookIdKey := fmt.Sprintf("%s%v", cacheBookBookIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, bookRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Title, data.Price)
	}, bookBookIdKey)
	return ret, err
}

func (m *defaultBookModel) Update(ctx context.Context, data *Book) error {
	bookBookIdKey := fmt.Sprintf("%s%v", cacheBookBookIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, bookRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Title, data.Price, data.Id)
	}, bookBookIdKey)
	return err
}

func (m *defaultBookModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheBookBookIdPrefix, primary)
}

func (m *defaultBookModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", bookRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultBookModel) tableName() string {
	return m.table
}
