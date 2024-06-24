package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"user/internal/svc"
	"user/internal/types"
	"user/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type SignupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignupLogic {
	return &SignupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

/*
	参数校验
	查询username是否已经被注册
	生成用户ID(雪花算法生成)
	加密密码(加盐MD5)


	下面的代码中只是完成了基本功能，在错误处理等方面存在这一些问题
*/
// 将用户信息插入到数据库中
var secret = []byte("张三")

func (l *SignupLogic) Signup(req *types.SignupRequest) (resp *types.SignupResponse, err error) {

	// todo: add your logic here and delete this line
	fmt.Println("req:", req)

	// 1.参数校验
	if req.Password != req.RePassword {
		return nil, errors.New("两次输入的密码不一致")
	}
	// 2.查询username是否已经被注册
	u, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.UserName)
	if err != nil && err != sqlx.ErrNotFound {
		// 出错了
		fmt.Println("查询数据库失败,", err)
		return nil, errors.New("内部错误,服务异常")
	}

	if u != nil {
		// 该username已被注册
		fmt.Println("用户已存在")
		return nil, errors.New("用户已存在")
	}
	// 3.生成ID(这里以时间戳作为ID,分布式系统可以尝试使用雪花算法)
	// 4.加密密码
	h := md5.New()
	h.Write([]byte(req.Password)) //密码计算MD5
	h.Write(secret)               //加盐
	passwordStr := hex.EncodeToString(h.Sum(nil))
	// 5.存入数据库
	user := &model.User{
		UserId:   time.Now().Unix(),
		Username: req.UserName,
		Password: passwordStr,
		Gender:   int64(req.Gender),
	}
	if _, err := l.svcCtx.UserModel.Insert(context.Background(), user); err != nil {
		return nil, err
	}
	return &types.SignupResponse{Message: "success"}, nil
}
