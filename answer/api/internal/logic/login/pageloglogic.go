package login

import (
	"activity/answer/api/internal/messages"
	"activity/answer/model"
	"context"
	"net/http"
	"strings"
	"time"

	"activity/answer/api/internal/svc"
	"activity/answer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPageLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageLogLogic {
	return &PageLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PageLogLogic) PageLog(req *types.PageRequest, r *http.Request) (resp *types.Response, err error) {
	resp = &types.Response{
		Code:    http.StatusOK,
		Message: messages.Success,
	}
	ip := strings.Split(r.RemoteAddr, ":")
	pageLogModel := model.NewPageLogModel(l.svcCtx.DbConn)
	pageLogModel.Insert(l.ctx, &model.PageLog{
		Uid:     req.GameUid,
		PageUrl: req.PageUrl,
		Ip:      ip[0],
		LogTime: time.Now(),
	})
	return resp, nil
}
