package logic

import (
	"context"

	"funnel/apps/funnel-rpc/internal/svc"
	"funnel/apps/funnel-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBorrowHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBorrowHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBorrowHistoryLogic {
	return &GetBorrowHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBorrowHistoryLogic) GetBorrowHistory(in *pb.LibrarySessionPageRequest) (*pb.GetBorrowHistoryResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetBorrowHistoryResponse{}, nil
}
