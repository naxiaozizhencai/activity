package login

import (
	"net/http"

	"activity/answer/api/internal/logic/login"
	"activity/answer/api/internal/svc"
	"activity/answer/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryTargetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryTargetRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := login.NewQueryTargetLogic(r.Context(), svcCtx)
		resp, err := l.QueryTarget(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
