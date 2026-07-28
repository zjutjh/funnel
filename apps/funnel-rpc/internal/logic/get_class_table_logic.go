package logic

import (
	"context"

	"funnel/apps/funnel-rpc/internal/svc"
	"funnel/apps/funnel-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClassTableLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetClassTableLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClassTableLogic {
	return &GetClassTableLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetClassTableLogic) GetClassTable(in *pb.ZfSessionTermRequest) (*pb.GetClassTableResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetClassTableResponse{}, nil
}
