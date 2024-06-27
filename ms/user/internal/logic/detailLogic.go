package logic

import (
	"context"
	"errors"

	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	if err == sqlx.ErrNotFound {
		logx.Error(err)
		return nil, err
	}
	if err != nil {
		logx.Error(err)
		return nil, errors.New("内部错误")
	}
	return &types.UserInfoResponse{
		Username: user.Username,
		Gender:   user.Gender,
		Age:      18,
	}, nil
}
