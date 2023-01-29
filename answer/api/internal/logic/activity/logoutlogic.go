package activity

import (
	"activity/answer/api/internal/logic/login"
	"activity/answer/api/internal/messages"
	"context"
	"fmt"
	"net/http"

	"activity/answer/api/internal/svc"
	"activity/answer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout() (resp *types.Response, err error) {
	gameUid := l.ctx.Value("userId").(string)
	codeCacheKey := fmt.Sprintf("%s%v", login.Code, gameUid)
	l.svcCtx.KvStore.Del(codeCacheKey)
	resp = &types.Response{
		Code:    http.StatusOK,
		Message: messages.Success,
	}
	return resp, nil
}
