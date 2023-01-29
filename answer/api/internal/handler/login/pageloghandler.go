package login

import (
	"net/http"

	"activity/answer/api/internal/logic/login"
	"activity/answer/api/internal/svc"
	"activity/answer/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PageLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := login.NewPageLogLogic(r.Context(), svcCtx)
		resp, err := l.PageLog(&req, r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
