package data

import (
	"context"
	"fmt"

	"bubblelist/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// todoRepo 实现了biz层定义的repo接口
type todoRepo struct {
	data *Data
	log  *log.Helper //透传下来方便做日志的
}

func NewTodoRepo(data *Data, logger log.Logger) biz.TodoRepo {
	return &todoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

/*
Save(context.Context, *Todo) (*Todo, error)
Update(context.Context, *Todo) error
Delete(context.Context, int64) error
FindByID(context.Context, int64) (*Todo, error)
ListByHello(context.Context, string) ([]*Todo, error)
ListAll(context.Context) ([]*Todo, error)
*/

func (r *todoRepo) Save(ctx context.Context, t *biz.Todo) (*biz.Todo, error) {
	// Save 这一步真正实现存储的操作
	fmt.Println("saving ... t:", t)
	// GORM-mysql Create  不建议AutoMigrate,这样对SQL表的可视性太差了,推荐自己建表格
	err := r.data.db.Create(t).Error //链式调用
	r.data.db.Save(&biz.Todo{})      //把biz.Todo这张表对应的数据进行保存
	return t, err
}

func (r *todoRepo) Update(ctx context.Context, t *biz.Todo) error {
	// 更新
	// err := r.data.db.WithContext(ctx).Model(t).Where("id = ?", t.ID).Update("title", t.Title).Update("status", t.Status).Error
	err := r.data.db.WithContext(ctx).Model(t).UpdateColumns(t).Error //uodateColumu 更新行
	// Model()指定那个表格,需要表有主键，支持直接传入结构体实现更新
	return err
}

func (r *todoRepo) FindByID(ctx context.Context, id int64) (*biz.Todo, error) {
	// First 查单条,条件ID
	t := &biz.Todo{ID: id}
	err := r.data.db.First(t).Error
	if err != nil {
		return nil, err
	}
	return t, err
}

func (r *todoRepo) Delete(ctx context.Context, id int64) error {
	// 你真别说这个GORM用起来是真轻松啊
	return r.data.db.WithContext(ctx).
		Where("id=?", id).Delete(&biz.Todo{}).Error
}

func (r *todoRepo) ListAll(context.Context) ([]*biz.Todo, error) {
	return nil, nil
}
