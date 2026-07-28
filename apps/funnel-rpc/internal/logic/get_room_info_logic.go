package logic

import (
	"context"

	"funnel/apps/funnel-rpc/internal/svc"
	"funnel/apps/funnel-rpc/pb"

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

func (l *GetRoomInfoLogic) GetRoomInfo(in *pb.GetRoomInfoRequest) (*pb.GetRoomInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetRoomInfoResponse{}, nil
}
