// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"funnel/apps/funnel-api/internal/svc"
	"funnel/apps/funnel-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCurrentBorrowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCurrentBorrowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCurrentBorrowLogic {
	return &GetCurrentBorrowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCurrentBorrowLogic) GetCurrentBorrow(req *types.LibraryBorrowReq) (resp *types.GetCurrentBorrowResp, err error) {
	// todo: add your logic here and delete this line

	return
}
