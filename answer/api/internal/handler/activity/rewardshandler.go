package activity

import (
	"net/http"

	"activity/answer/api/internal/logic/activity"
	"activity/answer/api/internal/svc"
	"activity/answer/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RewardsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RewardRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		var resp *types.Response
		var err error
		if req.AnswerId < 5 {
			l := activity.NewRewardsLogic(r.Context(), svcCtx)
			resp, err = l.Rewards(&req)
		} else {
			ll := activity.NewFragmentRewardsLogic(r.Context(), svcCtx)
			resp, err = ll.FragmentRewards(&req)
		}

		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
