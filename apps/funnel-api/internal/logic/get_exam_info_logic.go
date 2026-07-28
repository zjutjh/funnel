// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"funnel/apps/funnel-api/internal/svc"
	"funnel/apps/funnel-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExamInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetExamInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExamInfoLogic {
	return &GetExamInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetExamInfoLogic) GetExamInfo(req *types.SessionTermReq) (resp *types.GetExamInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
