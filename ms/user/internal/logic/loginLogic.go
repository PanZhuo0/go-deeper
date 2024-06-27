package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"

	"user/internal/svc"
	"user/internal/types"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

/*
1.参数检验

	用户名、密码是否都有

2.业务逻辑

	判断用户、密码是否正确？
	结果一致登录成功，否则失败
	两种方式:
		1.输入用户名和密码(加密后) 去查数据库 是否有这条记录
		2.根据用户名获取密码，比较用户密码加密后是否与数据库中密码相同(这个例子用这个)

3.响应结果

	成功/失败
*/
func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOneByUsername(context.Background(), req.Username)
	if err == sqlx.ErrNotFound {
		return &types.LoginResponse{
			Message: "用户名不存在",
		}, nil
	}
	if err != nil {
		logx.Errorw("UserModel.FindOneByUsername failed", logx.Field("err", err))
		return &types.LoginResponse{
			Message: "用户不存在",
		}, errors.New("内部错误")
	}
	if user.Password != passwordMD5([]byte(req.Password)) {
		return &types.LoginResponse{
			Message: "用户名或密码错误",
		}, nil
	}
	return &types.LoginResponse{Message: "登录成功"}, nil
}

func passwordMD5(password []byte) string {
	h := md5.New()
	h.Write([]byte(password)) //密码计算MD5
	h.Write(secret)           //加盐
	passwordStr := hex.EncodeToString(h.Sum(nil))
	return passwordStr
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userID int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userid"] = userID
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey)) //签名，防止篡改
}
