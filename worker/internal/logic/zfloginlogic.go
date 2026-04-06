package logic

import (
	"context"

	"funnel/api/rpc/funnel"
	"funnel/worker/internal/svc"

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

func (l *ZfLoginLogic) ZfLogin(in *funnel.ZfLoginRequest) (*funnel.ZfLoginResponse, error) {
	// todo: add your logic here and delete this line

	return &funnel.ZfLoginResponse{}, nil
}
