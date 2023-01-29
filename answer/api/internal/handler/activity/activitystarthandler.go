package activity

import (
	"net/http"

	"activity/answer/api/internal/logic/activity"
	"activity/answer/api/internal/svc"
	"activity/answer/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ActivityStartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StartAnswerRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := activity.NewActivityStartLogic(r.Context(), svcCtx)
		resp, err := l.ActivityStart(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
