// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"funnel/apps/funnel-api/internal/svc"
	"funnel/apps/funnel-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClassTableLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetClassTableLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClassTableLogic {
	return &GetClassTableLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetClassTableLogic) GetClassTable(req *types.SessionTermReq) (resp *types.GetClassTableResp, err error) {
	// todo: add your logic here and delete this line

	return
}
