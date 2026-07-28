package logic

import (
	"context"

	"funnel/apps/funnel-rpc/internal/svc"
	"funnel/apps/funnel-rpc/pb"

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

func (l *GetExamInfoLogic) GetExamInfo(in *pb.ZfSessionTermRequest) (*pb.GetExamInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetExamInfoResponse{}, nil
}
