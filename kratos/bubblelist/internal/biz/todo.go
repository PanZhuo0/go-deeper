package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Todo is a todo model.
type Todo struct {
	ID     int64
	Title  string
	Status bool
}

// TodoRepo is a todo repo.

// biz层对数据操作层提出了以下要求,只要结构体实现这些方法即可
type TodoRepo interface {
	Save(context.Context, *Todo) (*Todo, error)
	Update(context.Context, *Todo) error
	Delete(context.Context, int64) error
	FindByID(context.Context, int64) (*Todo, error)
	ListAll(context.Context) ([]*Todo, error)
}

// TodoUsecase is a Todo usecase.
type TodoUsecase struct {
	repo TodoRepo
	log  *log.Helper
}

// NewTodoUsecase new a Todo usecase.
func NewTodoUsecase(repo TodoRepo, logger log.Logger) *TodoUsecase {
	return &TodoUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateTodo creates a Tode, and returns the new Todo.
func (uc *TodoUsecase) Create(ctx context.Context, t *Todo) (*Todo, error) {
	uc.log.WithContext(ctx).Infof("CreateTodo: %v", t)
	return uc.repo.Save(ctx, t) //调用底层的save函数
}

func (uc *TodoUsecase) Get(ctx context.Context, id int64) (*Todo, error) {
	uc.log.WithContext(ctx).Infof("GetTodo:%v", id)
	return uc.repo.FindByID(ctx, id) //调用下一层的server方法
}

func (uc *TodoUsecase) Update(ctx context.Context, t *Todo) error {
	uc.log.WithContext(ctx).Infof("GetTodo:%v", t)
	return uc.repo.Update(ctx, t)
}

func (uc *TodoUsecase) Delete(ctx context.Context, id int64) error {
	uc.log.WithContext(ctx).Infof("GetTodo:%v", id)
	return uc.repo.Delete(ctx, id)
}
