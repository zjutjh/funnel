// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"funnel/apps/funnel-api/internal/svc"
	"funnel/apps/funnel-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBorrowHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBorrowHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBorrowHistoryLogic {
	return &GetBorrowHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBorrowHistoryLogic) GetBorrowHistory(req *types.LibraryBorrowReq) (resp *types.GetBorrowHistoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
