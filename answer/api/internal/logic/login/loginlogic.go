package login

import (
	"activity/answer/api/internal/messages"
	"activity/answer/model"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"

	"activity/answer/api/internal/svc"
	"activity/answer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	codeCacheKey := fmt.Sprintf("%s%v", Code, req.GameUid)
	codeValue, err := l.svcCtx.KvStore.Get(codeCacheKey)
	result := types.LoginResponse{
		Response: types.Response{
			Code:    http.StatusOK,
			Message: messages.Success,
		},
		Data: types.LoginData{
			GameUid:         "",
			Token:           "",
			AccessExpire:    0,
			RefreshAfter:    0,
			FragmentRewards: make([]string, 0),
		},
	}
	if err != nil {
		result.Code = http.StatusInternalServerError
		result.Message = messages.DataErr
		return &result, nil
	}
	if codeValue != req.AuthCode {
		result.Code = http.StatusInternalServerError
		result.Message = messages.AuthCodeErr
		return &result, nil
	}
	token, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.Auth.AccessExpire, req.GameUid, req.Language)
	result.Data.GameUid = req.GameUid
	result.Data.Language = req.Language
	result.Data.Token = token
	result.Data.AccessExpire = time.Now().Unix()
	result.Data.RefreshAfter = time.Now().Unix() + l.svcCtx.Config.Auth.AccessExpire
	rewardModel := model.NewRewardModel(l.svcCtx.DbConn, l.svcCtx.Config.CacheRedis)
	fragmentRewards, err := rewardModel.FindUserFragmentRewards(l.ctx, req.GameUid, []string{"1", "2", "3", "4", "5"})
	if err != nil {
		result.Code = http.StatusInternalServerError
		result.Message = messages.DataErr
		return &result, nil
	}
	for _, v := range fragmentRewards {
		result.Data.FragmentRewards = append(result.Data.FragmentRewards, v.ItemId)
	}
	loginModel := model.NewLoginLogModel(l.svcCtx.DbConn)
	loginModel.Insert(l.ctx, &model.LoginLog{
		Language:  req.Language,
		Uid:       req.GameUid,
		LoginTime: time.Now(),
		LogTime:   time.Now(),
	})
	//删除验证码
	l.svcCtx.KvStore.Del(codeCacheKey)
	return &result, nil
}
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds int64, userId string, language string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	claims["language"] = language
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
