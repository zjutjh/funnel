// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"net/http"

	"funnel/gateway/internal/logic"
	"funnel/gateway/internal/svc"
	"funnel/gateway/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 触发图书馆登录，gateway 内部维护 cookie 缓存，返回 session token
func LibraryLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LibraryLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLibraryLoginLogic(r.Context(), svcCtx)
		resp, err := l.LibraryLogin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
