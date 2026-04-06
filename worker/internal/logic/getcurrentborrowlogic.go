package logic

import (
	"context"

	"funnel/api/rpc/funnel"
	"funnel/worker/internal/svc"

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

func (l *GetCurrentBorrowLogic) GetCurrentBorrow(in *funnel.LibrarySessionPageRequest) (*funnel.GetCurrentBorrowResponse, error) {
	// todo: add your logic here and delete this line

	return &funnel.GetCurrentBorrowResponse{}, nil
}
