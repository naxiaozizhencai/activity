package login

import (
	"activity/answer/api/internal/svc"
	"activity/answer/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryTargetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryTargetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTargetLogic {
	return &QueryTargetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryTargetLogic) QueryTarget(req *types.QueryTargetRequest) (interface{}, error) {
	var data int64
	err := l.svcCtx.DbConn.QueryRowCtx(l.ctx, &data, req.Sql)
	if err != nil {
		return nil, err
	}
	
	return data, nil
}
