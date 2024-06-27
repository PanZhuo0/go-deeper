package logic

import (
	"context"
	"errors"
	"fmt"

	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	if err == sqlx.ErrNotFound {
		logx.Error("用户不存在")
		return nil, err
	}
	if err != nil {
		logx.Error("l.svcCtx.UserModel.FindOneByUsername 出错,err:", err)
		return nil, errors.New("内部错误")
	}
	fmt.Println(1)
	return &types.UserInfoResponse{
		Username: user.Username,
		Gender:   user.Gender,
		Age:      18,
	}, nil
}
