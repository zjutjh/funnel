package logic

import (
	"context"

	"funnel/apps/funnel-rpc/internal/svc"
	"funnel/apps/funnel-rpc/pb"

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

func (l *LibraryLoginLogic) LibraryLogin(in *pb.LibraryLoginRequest) (*pb.LibraryLoginResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.LibraryLoginResponse{}, nil
}
