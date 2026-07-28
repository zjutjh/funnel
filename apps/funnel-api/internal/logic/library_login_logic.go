// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"funnel/apps/funnel-api/internal/svc"
	"funnel/apps/funnel-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LibraryLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 触发图书馆登录，gateway 内部维护 cookie 缓存，返回 session token
func NewLibraryLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LibraryLoginLogic {
	return &LibraryLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LibraryLoginLogic) LibraryLogin(req *types.LibraryLoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line

	return
}
