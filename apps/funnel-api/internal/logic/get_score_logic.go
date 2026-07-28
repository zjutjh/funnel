// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"funnel/apps/funnel-api/internal/svc"
	"funnel/apps/funnel-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetScoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScoreLogic {
	return &GetScoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetScoreLogic) GetScore(req *types.SessionTermReq) (resp *types.GetScoreResp, err error) {
	// todo: add your logic here and delete this line

	return
}
