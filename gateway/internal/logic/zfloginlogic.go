// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"funnel/gateway/internal/svc"
	"funnel/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ZfLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 触发正方登录，gateway 内部维护 cookie 缓存，返回 session token
func NewZfLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ZfLoginLogic {
	return &ZfLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ZfLoginLogic) ZfLogin(req *types.ZfLoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line

	return
}
