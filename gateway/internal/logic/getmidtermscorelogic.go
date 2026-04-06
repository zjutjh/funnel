// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"funnel/gateway/internal/svc"
	"funnel/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMidTermScoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMidTermScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMidTermScoreLogic {
	return &GetMidTermScoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMidTermScoreLogic) GetMidTermScore(req *types.SessionTermReq) (resp *types.GetMidTermScoreResp, err error) {
	// todo: add your logic here and delete this line

	return
}
