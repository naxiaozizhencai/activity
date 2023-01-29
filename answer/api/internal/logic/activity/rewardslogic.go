package activity

import (
	"activity/answer/api/internal/messages"
	"activity/answer/api/internal/svc"
	"activity/answer/api/internal/types"
	"activity/answer/model"
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"net/http"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RewardsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRewardsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RewardsLogic {
	return &RewardsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RewardsLogic) Rewards(req *types.RewardRequest) (resp *types.Response, err error) {
	gameUid := l.ctx.Value("userId").(string)
	answerModel := model.NewAnwserModel(l.svcCtx.DbConn, l.svcCtx.Config.CacheRedis)
	answerData, err := answerModel.FindOne(l.ctx, req.AnswerId)
	resp = &types.Response{
		Code:    http.StatusOK,
		Message: messages.Success,
	}
	if err != nil {
		return
	}
	answerLogic := NewAnswerListLogic(l.ctx, l.svcCtx)
	answerStatus := answerLogic.checkAnswerStatus(gameUid, req.AnswerId, *answerData)
	if answerStatus != 2 {
		resp.Code = http.StatusInternalServerError
		resp.Message = messages.Reward_Err1
		return resp, nil
	}
	rewardStatus := answerLogic.checkRewardStatus(gameUid, req.AnswerId, *answerData, answerStatus)
	if rewardStatus != 1 {
		resp.Code = http.StatusInternalServerError
		resp.Message = messages.Reward_Err1
		return resp, nil
	}

	rewardModel := model.NewRewardModel(l.svcCtx.DbConn, l.svcCtx.Config.CacheRedis)
	reward, err := rewardModel.FindUserAward(l.ctx, gameUid, req.AnswerId, answerData.ItemId.String)
	if err != nil && err != sqlx.ErrNotFound {
		resp.Code = http.StatusInternalServerError
		resp.Message = messages.DataErr
		return resp, nil
	}
	var newRewardId int64
	if reward == nil {
		rewardResult, err := rewardModel.Insert(l.ctx, &model.Reward{
			GameUid:  gameUid,
			AnswerId: req.AnswerId,
			ItemId:   answerData.ItemId.String,
			Nums:     answerData.ItemNum.Int64,
			Status:   0,
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
		newRewardId, _ = rewardResult.LastInsertId()
	} else {
		newRewardId = reward.Id
	}
	uid, err := strconv.ParseInt(gameUid, 10, 64)
	items := make([]types.Item, 0)
	//itemId, _ := strconv.ParseInt(answerData.ItemId.String, 10, 64)
	items = append(items, types.Item{
		ItemId:    answerData.ItemId.String,
		ItemCount: answerData.ItemNum.Int64,
	})
	content := "<m01><arg>" + req.Context + "</arg></m01><m07></m07>"
	params := types.SendMailParams{
		AppId:      l.svcCtx.Config.GameApiConfig.AppId,
		SvrRegion:  l.svcCtx.Config.GameApiConfig.SvrRegion,
		Type:       "send_mail",
		SvrId:      0,
		SenderName: req.SenderName,
		RecverUids: []int64{uid},
		Title:      req.Title,
		Context:    content,
		Coins:      make([]types.Coin, 0),
		Items:      items,
		AffectTime: time.Now().Unix(),
		ExpireTime: time.Now().Unix() + 30*86400,
		Assets:     nil,
	}
	sendmailResult, err := svc.SendGameEmail(l.svcCtx.Config.GameApiConfig, params)
	logx.Info("send mail result ", sendmailResult, err)
	/*sendmailResult = 0
	err = nil*/
	if err != nil || sendmailResult != 0 {
		resp.Code = http.StatusInternalServerError
		resp.Message = messages.SendMailErr
		return resp, nil
	}
	err = rewardModel.UpdateStatus(l.ctx, newRewardId, 1)
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = messages.SendMailErr
		return resp, nil
	}
	rewardLogModel := model.NewRewardLogModel(l.svcCtx.DbConn)
	rewardLogModel.Insert(l.ctx, &model.RewardLog{
		AnwserId: req.AnswerId,
		ItemId:   answerData.ItemId.String,
		ItemNum:  answerData.ItemNum.Int64,
		Uid:      gameUid,
		LogTime:  time.Now(),
	})
	return
}
