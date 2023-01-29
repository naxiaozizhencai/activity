package activity

import (
	"net/http"

	"activity/answer/api/internal/logic/activity"
	"activity/answer/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FragmentRewardsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := activity.NewFragmentRewardsLogic(r.Context(), svcCtx)
		resp, err := l.FragmentRewards(nil)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
