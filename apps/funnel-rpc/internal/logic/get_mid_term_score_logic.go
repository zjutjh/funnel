package logic

import (
	"context"

	"funnel/apps/funnel-rpc/internal/svc"
	"funnel/apps/funnel-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMidTermScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMidTermScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMidTermScoreLogic {
	return &GetMidTermScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMidTermScoreLogic) GetMidTermScore(in *pb.ZfSessionTermRequest) (*pb.GetMidTermScoreResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetMidTermScoreResponse{}, nil
}
