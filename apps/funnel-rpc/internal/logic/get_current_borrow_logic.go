package logic

import (
	"context"

	"funnel/apps/funnel-rpc/internal/svc"
	"funnel/apps/funnel-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCurrentBorrowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCurrentBorrowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCurrentBorrowLogic {
	return &GetCurrentBorrowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCurrentBorrowLogic) GetCurrentBorrow(in *pb.LibrarySessionPageRequest) (*pb.GetCurrentBorrowResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetCurrentBorrowResponse{}, nil
}
