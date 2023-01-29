package activity

import (
	"activity/answer/api/internal/messages"
	"activity/answer/api/internal/svc"
	"activity/answer/api/internal/types"
	"activity/answer/model"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"net/http"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnswerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAnswerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnswerListLogic {
	return &AnswerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AnswerListLogic) AnswerList() (resp *types.AnswerListResponse, err error) {
	gameUid := l.ctx.Value("userId").(string)

	answerModel := model.NewAnwserModel(l.svcCtx.DbConn, l.svcCtx.Config.CacheRedis)
	answerData, err := answerModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}
	result := &types.AnswerListResponse{
		Response: types.Response{
			Code:    http.StatusOK,
			Message: messages.Success,
		},
		Data: make([]types.AnswerListData, 0),
	}
	loc, _ := time.LoadLocation("Local")
	endDateTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2023-01-12 08:00:00", loc)
	if time.Now().Unix() > endDateTime.Unix() {
		result.Code = http.StatusInternalServerError
		result.Message = messages.ActivityOver
		return result, nil
	}

	for _, v := range answerData {
		answerStatus := l.checkAnswerStatus(gameUid, v.AnwserId, v)
		timesKey := fmt.Sprintf(model.CacheAnwserIdUserTimesPrefix, gameUid, v.AnwserId)
		timesStr, err := l.svcCtx.KvStore.Get(timesKey)
		if err != nil {
			return nil, err
		}
		var surplusTimes int
		if answerStatus == 1 {
			times, _ := strconv.ParseInt(timesStr, 10, 64)
			if times == 0 {
				surplusTimes = model.AnswerTimes
			} else {
				surplusTimes = model.AnswerTimes - int(times)
			}
		}
		result.Data = append(result.Data, types.AnswerListData{
			AnswerId:     v.AnwserId,
			AnwerName:    v.AnwserName.String,
			AnserStatus:  answerStatus,
			RewardStatus: l.checkRewardStatus(gameUid, v.AnwserId, v, answerStatus),
			OpenTime:     v.StartTime.Format("2006-01-02 15:04:05"),
			SurplusTimes: surplusTimes,
		})
	}

	return result, nil
}

func (l *AnswerListLogic) checkAnswerStatus(gameUid string, anwserId int64, anwserData model.Anwser) int {
	answerResultModel := model.NewAnwserResultModel(l.svcCtx.DbConn, l.svcCtx.Config.CacheRedis)
	answerResultInfo, err := answerResultModel.FindOneByGameUidAnswerId(l.ctx, gameUid, anwserId)
	if err != nil && err != sqlx.ErrNotFound {
		return 0
	}
	if answerResultInfo != nil && answerResultInfo.Status == 1 {
		return 2
	}
	//判断上一个关卡是否开启
	if anwserData.LastAnswerId != 0 {
		_, err := answerResultModel.FindOneByGameUidAnswerId(l.ctx, gameUid, anwserData.LastAnswerId)
		if err != nil {
			return 0
		}
	}
	now := time.Now().Unix()
	if now >= anwserData.StartTime.Unix() {
		return 1
	}
	return 0
}
func (l *AnswerListLogic) checkRewardStatus(gameUid string, anwserId int64, anwserData model.Anwser, answerStatus int) int {
	if answerStatus == 0 || answerStatus == 1 {
		return 0
	}
	rewardModel := model.NewRewardModel(l.svcCtx.DbConn, l.svcCtx.Config.CacheRedis)
	rewardInfo, err := rewardModel.FindUserAward(l.ctx, gameUid, anwserId, anwserData.ItemId.String)
	if err != nil && err != sqlx.ErrNotFound {
		return 0
	}
	if rewardInfo != nil && rewardInfo.Status == 1 {
		return 2
	}
	//碎片奖励回答完五道题才可以领取
	if anwserId == 5 {
		fragmentRewards, err := rewardModel.FindUserFragmentRewards(l.ctx, gameUid, fragmentRewardsItemIds)
		if err != nil {
			return 0
		}
		fragmentRewardsList := make(map[string]int64, 0)
		for _, v := range fragmentRewards {
			fragmentRewardsList[v.ItemId] = v.Nums
		}
		for _, v := range fragmentRewardsItemIds {
			if _, ok := fragmentRewardsList[v]; !ok {
				return 0
			}
		}
	}
	if answerStatus == 2 {
		return 1
	}

	return 0
}
