package logic

import (
	"context"

	"funnel/api/rpc/funnel"
	"funnel/worker/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExamInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetExamInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExamInfoLogic {
	return &GetExamInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetExamInfoLogic) GetExamInfo(in *funnel.ZfSessionTermRequest) (*funnel.GetExamInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &funnel.GetExamInfoResponse{}, nil
}
