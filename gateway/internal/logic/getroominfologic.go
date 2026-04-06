// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"funnel/gateway/internal/svc"
	"funnel/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoomInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoomInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoomInfoLogic {
	return &GetRoomInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoomInfoLogic) GetRoomInfo(req *types.GetRoomInfoReq) (resp *types.GetRoomInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
