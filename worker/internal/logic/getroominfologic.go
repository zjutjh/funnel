package logic

import (
	"context"

	"funnel/api/rpc/funnel"
	"funnel/worker/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoomInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoomInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoomInfoLogic {
	return &GetRoomInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoomInfoLogic) GetRoomInfo(in *funnel.GetRoomInfoRequest) (*funnel.GetRoomInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &funnel.GetRoomInfoResponse{}, nil
}
