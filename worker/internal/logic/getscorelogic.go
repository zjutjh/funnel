package logic

import (
	"context"

	"funnel/api/rpc/funnel"
	"funnel/worker/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScoreLogic {
	return &GetScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetScoreLogic) GetScore(in *funnel.ZfSessionTermRequest) (*funnel.GetScoreResponse, error) {
	// todo: add your logic here and delete this line

	return &funnel.GetScoreResponse{}, nil
}
