// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"net/http"

	"funnel/apps/funnel-api/internal/logic"
	"funnel/apps/funnel-api/internal/svc"
	"funnel/apps/funnel-api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetClassTableHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SessionTermReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetClassTableLogic(r.Context(), svcCtx)
		resp, err := l.GetClassTable(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
