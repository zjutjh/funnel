package logic

import (
	"context"

	"funnel/api/rpc/funnel"
	"funnel/worker/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LibraryLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLibraryLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LibraryLoginLogic {
	return &LibraryLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LibraryLoginLogic) LibraryLogin(in *funnel.LibraryLoginRequest) (*funnel.LibraryLoginResponse, error) {
	// todo: add your logic here and delete this line

	return &funnel.LibraryLoginResponse{}, nil
}
