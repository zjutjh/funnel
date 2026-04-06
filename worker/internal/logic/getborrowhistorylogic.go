package logic

import (
	"context"

	"funnel/api/rpc/funnel"
	"funnel/worker/internal/svc"

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

func (l *GetBorrowHistoryLogic) GetBorrowHistory(in *funnel.LibrarySessionPageRequest) (*funnel.GetBorrowHistoryResponse, error) {
	// todo: add your logic here and delete this line

	return &funnel.GetBorrowHistoryResponse{}, nil
}
