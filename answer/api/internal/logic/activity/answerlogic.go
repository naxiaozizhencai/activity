package activity

import (
	"activity/answer/api/internal/messages"
	"activity/answer/api/internal/svc"
	"activity/answer/api/internal/types"
	"activity/answer/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnswerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var cacheUserFragmentRewardsPrefix = "cache:user:fragment_reward:items:"
var fragmentRewardsItemIds = []string{"1", "2", "3", "4", "5"}

func NewAnswerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnswerLogic {
	return &AnswerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AnswerLogic) Answer(req *types.AnswerRequest) (resp *types.AnswerReponse, err error) {
	gameUid := l.ctx.Value("userId").(string)
	answerModel := model.NewAnwserModel(l.svcCtx.DbConn, l.svcCtx.Config.CacheRedis)
	answerInfo, err := answerModel.FindOne(l.ctx, req.AnswerId)
	if err != nil {
		return nil, err
	}
	resp = &types.AnswerReponse{
		Response: types.Response{
			Code:    http.StatusOK,
			Message: messages.Success,
		},
		Data: types.AnswerData{
			ItemId:       "",
			SurplusTimes: 0,
		},
	}
	answerLogic := NewAnswerListLogic(l.ctx, l.svcCtx)
	answerStatus := answerLogic.checkAnswerStatus(gameUid, req.AnswerId, *answerInfo)
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

	timesKey := fmt.Sprintf(model.CacheAnwserIdUserTimesPrefix, gameUid, req.AnswerId)
	timesStr, err := l.svcCtx.KvStore.Get(timesKey)
	if err != nil {
		return nil, err
	}
	times, _ := strconv.ParseInt(timesStr, 10, 64)
	if times >= int64(model.AnswerTimes) {
		resp.Code = http.StatusInternalServerError
		resp.Message = messages.AnswerTimesOut
		resp.Data.SurplusTimes = 0
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
		resp.Message = messages.DataErr
		return resp, nil
	}
	/*	expireKey := fmt.Sprintf("%s%v", cacheAnwserIdUserExpireTimePrefix, gameUid)
		answerExpireData, err := l.svcCtx.KvStore.Get(expireKey)
		if err != nil {
			resp.Code = http.StatusInternalServerError
			resp.Message = messages.DataErr
			return resp, nil
		}
		if answerExpireData == "" {
			resp.Code = http.StatusInternalServerError
			resp.Message = messages.AnswerTimeOut
			return resp, nil
		}*/
	var surplusTimes int
	if times == 0 {
		surplusTimes = 1
	} else {
		surplusTimes = model.AnswerTimes - int(times+1)
	}
	var languageStr string
	if language, ok := l.ctx.Value("language").(string); ok {
		languageStr = language
	}

	answerLogModel := model.NewAnswerLogModel(l.svcCtx.DbConn)
	answerLogModel.Insert(l.ctx, &model.AnswerLog{
		Uid:          gameUid,
		Languge:      languageStr,
		AnswerId:     req.AnswerId,
		UserResult:   req.Result,
		AnswerResult: answerInfo.Result,
		LogTime:      time.Now(),
	})
	if req.AnswerId != 5 {
		if answerInfo.Result != req.Result {
			//删除过期事件
			//l.svcCtx.KvStore.Del(expireKey)
			resp.Code = http.StatusInternalServerError
			resp.Message = messages.AnswerErr
			resp.Data.SurplusTimes = surplusTimes
			return resp, nil
		}
	}

	fragmentRewardsPrefix := fmt.Sprintf("%s%v", cacheUserFragmentRewardsPrefix, gameUid)
	itemId, err := l.getFragmentRewardsItemId(gameUid)
	if err != nil {
		//result.Message = "碎片已获取完毕"
		resp.Code = http.StatusInternalServerError
		resp.Message = messages.AnswerErr
		return resp, nil
	}
	_, err = l.svcCtx.KvStore.SaddCtx(l.ctx, fragmentRewardsPrefix, itemId)
	if err != nil {
		return nil, err
	}
	//记录答题结果
	answerResultModel := model.NewAnwserResultModel(l.svcCtx.DbConn, l.svcCtx.Config.CacheRedis)
	answerResult, err := answerResultModel.FindOneByGameUidAnswerId(l.ctx, gameUid, req.AnswerId)
	if err != nil && err != sqlx.ErrNotFound {
		resp.Code = http.StatusInternalServerError
		resp.Message = messages.DataErr
		return resp, nil
	}
	if answerResult != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = messages.AnswerHasSuccess
		return resp, nil
	}
	_, err = answerResultModel.Insert(l.ctx, &model.AnwserResult{
		GameUid:  gameUid,
		AnswerId: req.AnswerId,
		Status:   1,
		AddTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = messages.SaveFail
		return resp, nil
	}

	//发送碎片
	rewardModel := model.NewRewardModel(l.svcCtx.DbConn, l.svcCtx.Config.CacheRedis)
	_, err = rewardModel.Insert(l.ctx, &model.Reward{
		GameUid:  gameUid,
		AnswerId: req.AnswerId,
		ItemId:   itemId,
		Nums:     1,
		Status:   1,
		AddTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = messages.SaveFail
		return resp, nil
	}

	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = messages.SaveFail
		return resp, nil
	}

	resp.Data.ItemId = itemId
	resp.Data.SurplusTimes = surplusTimes
	return resp, nil
}

func (l *AnswerLogic) getFragmentRewardsItemId(gameUid string) (string, error) {
	fragmentRewardsPrefix := fmt.Sprintf("%s%v", cacheUserFragmentRewardsPrefix, gameUid)
	items, err := l.svcCtx.KvStore.Smembers(fragmentRewardsPrefix)
	if err != nil {
		return "", err
	}
	existItems := make(map[string]int, 0)
	for _, v := range items {
		existItems[v] = 1
	}
	randItems := make([]string, 0)
	for _, v := range fragmentRewardsItemIds {
		if _, ok := existItems[v]; ok {
			continue
		}
		randItems = append(randItems, v)
	}
	if len(randItems) == 0 {
		return "", errors.New("碎片已获取完毕")
	}
	rand.Seed(time.Now().UnixNano())
	randKey := rand.Intn(len(randItems))
	return randItems[randKey], nil
}
