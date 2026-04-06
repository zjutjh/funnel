package logic

import (
	"context"

	"funnel/api/rpc/funnel"
	"funnel/worker/internal/svc"

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

func (l *GetMidTermScoreLogic) GetMidTermScore(in *funnel.ZfSessionTermRequest) (*funnel.GetMidTermScoreResponse, error) {
	// todo: add your logic here and delete this line

	return &funnel.GetMidTermScoreResponse{}, nil
}
