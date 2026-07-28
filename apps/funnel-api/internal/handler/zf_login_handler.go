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

// 触发正方登录，gateway 内部维护 cookie 缓存，返回 session token
func ZfLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ZfLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewZfLoginLogic(r.Context(), svcCtx)
		resp, err := l.ZfLogin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
