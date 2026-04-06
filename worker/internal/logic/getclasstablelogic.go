package logic

import (
	"context"

	"funnel/api/rpc/funnel"
	"funnel/worker/internal/svc"

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

func (l *GetClassTableLogic) GetClassTable(in *funnel.ZfSessionTermRequest) (*funnel.GetClassTableResponse, error) {
	// todo: add your logic here and delete this line

	return &funnel.GetClassTableResponse{}, nil
}
