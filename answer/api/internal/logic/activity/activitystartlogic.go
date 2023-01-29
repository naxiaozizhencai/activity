package activity

import (
	"activity/answer/api/internal/messages"
	"activity/answer/api/internal/svc"
	"activity/answer/api/internal/types"
	"activity/answer/model"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityStartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityStartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityStartLogic {
	return &ActivityStartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityStartLogic) ActivityStart(req *types.StartAnswerRequest) (resp *types.StartAnswerResponse, err error) {
	userId := l.ctx.Value("userId").(string)
	answerModel := model.NewAnwserModel(l.svcCtx.DbConn, l.svcCtx.Config.CacheRedis)
	answerInfo, err := answerModel.FindOne(l.ctx, req.AnswerId)
	resp = &types.StartAnswerResponse{
		Response: types.Response{
			Code:    http.StatusOK,
			Message: messages.Success,
		},
		Data: types.StartAnswerData{},
	}
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = messages.DataErr
		return resp, nil
	}

	answerLogic := NewAnswerListLogic(l.ctx, l.svcCtx)
	answerStatus := answerLogic.checkAnswerStatus(userId, req.AnswerId, *answerInfo)
	if answerStatus == 0 {
		resp.Code = http.StatusInternalServerError
		resp.Message = messages.ActiveOpenErr
		return resp, nil
	}
	if answerStatus == 2 {
		resp.Code = http.StatusInternalServerError
		resp.Message = messages.AnswerHasSuccess
		return resp, nil
	}
	timesKey := fmt.Sprintf("%s%v%d", model.CacheAnwserIdUserTimesPrefix, userId, req.AnswerId)
	timesStr, err := l.svcCtx.KvStore.Get(timesKey)
	if err != nil {
		return nil, err
	}

	times, _ := strconv.ParseInt(timesStr, 10, 64)
	if times >= int64(model.AnswerTimes) {
		resp.Code = http.StatusInternalServerError
		resp.Message = messages.AnswerTimesOut
		return resp, nil
	}
	expireKey := fmt.Sprintf("%s%v", model.CacheAnwserIdUserExpireTimePrefix, userId)
	isExpire, err := l.svcCtx.KvStore.SetnxEx(expireKey, "1", 180)
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = messages.DataErr
		return resp, nil
	}
	if !isExpire {
		return resp, nil
	}
	_, err = l.svcCtx.KvStore.IncrCtx(l.ctx, timesKey)
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = messages.DataErr
		return resp, nil
	}
	nowDate := time.Now().Format("2006-01-02")
	loc, _ := time.LoadLocation("Local")
	endDateTime, err := time.ParseInLocation("2006-01-02 15:04:05", nowDate+" 23:59:59", loc)
	err = l.svcCtx.KvStore.ExpireatCtx(l.ctx, timesKey, endDateTime.Unix())
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = messages.SaveFail
		return resp, nil
	}
	var surplusTimes int
	if times == 0 {
		surplusTimes = 1
	} else {
		surplusTimes = model.AnswerTimes - int(times)
	}
	resp.Data.SurplusTimes = surplusTimes
	return resp, nil
}
