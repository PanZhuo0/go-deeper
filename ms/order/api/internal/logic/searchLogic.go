package logic

import (
	"context"
	"user/rpc/userclient"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchRequset) (resp *types.SearchResponse, err error) {
	// todo: add your logic here and delete this line
	// 1.根据请求参数中的订单号查询数据库找到订单服务（课后作业:用go-zero框架生成对应的代码）
	// 2.更具订单记录中的user_id 去查询用户数据（通过RPC调用user服务）
	// 假设user_id=xxxxxx
	userResp, err := l.svcCtx.UserRPC.GetUser(l.ctx, &userclient.GetUserReq{UserID: xxxxx})
	if err != nil {
		logx.Errorw("UserRPC.GetUser failed,", logx.Field("err", err))
		return nil, err
	}
	// 3.拼接返回结果（因为我们这个接口的数据不是由当前这一个服务组成的）
	return &types.SearchResponse{
		OrderID:  "xxxx", //根据实际结果来赋值，这里是假数据
		Status:   100,    //根据实际结果来赋值，这里是假数据
		Username: userResp.GetUsername(),
	}, nil
}
