package logic

import (
	"context"

	"funnel/apps/funnel-rpc/internal/svc"
	"funnel/apps/funnel-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ZfLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewZfLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ZfLoginLogic {
	return &ZfLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ZfLoginLogic) ZfLogin(in *pb.ZfLoginRequest) (*pb.ZfLoginResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.ZfLoginResponse{}, nil
}
