package login

import (
	"activity/answer/api/internal/messages"
	"activity/answer/api/internal/svc"
	"activity/answer/api/internal/types"
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	Code      = "cache:code:"
	CodeLimit = "cache:code:limit:"
)

type SendCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *SendCodeLogic) SendCode(req *types.SendCodeRequest) (resp *types.SendCodeResponse, err error) {
	codeCacheKey := fmt.Sprintf("%s%v", Code, req.GameUid)
	codeLimitCacheKey := fmt.Sprintf("%s%v", CodeLimit, req.GameUid)
	codeValue, err := l.svcCtx.KvStore.GetCtx(l.ctx, codeLimitCacheKey)
	var result types.SendCodeResponse
	if err != nil {
		result.Code = http.StatusInternalServerError
		result.Message = messages.DataErr
		return &result, nil
	}
	if codeValue != "" {
		result.Code = http.StatusInternalServerError
		result.Message = messages.AuthCodeExist
		return &result, nil
	}
	/*	uid, err := strconv.ParseInt(req.GameUid, 10, 64)
		if err != nil {
			result.Code = http.StatusInternalServerError
			result.Message = messages.DataErr
			return &result, nil
		}*/
	randCode := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	//randCode = "123456"
	logx.Info("rand code ", randCode)
	err = l.svcCtx.KvStore.SetexCtx(l.ctx, codeLimitCacheKey, "1", 50)
	/*	params := types.SendMailParams{
		AppId:     l.svcCtx.Config.GameApiConfig.AppId,
		SvrRegion: l.svcCtx.Config.GameApiConfig.SvrRegion,
		//AppId:     0,
		//SvrRegion: 0,
		Type:       "send_mail",
		SvrId:      0,
		SenderName: req.SenderName,
		RecverUids: []int64{uid},
		Title:      req.Title,
		Context:    fmt.Sprintf(req.Context, randCode),
		Coins:      nil,
		Items:      nil,
		AffectTime: time.Now().Unix(),
		ExpireTime: time.Now().Unix() + 300,
		Assets:     nil,
	}*/
	params := types.SendCodeParams{
		Type: "mail",
		Uid:  req.GameUid,
		Code: randCode,
	}
	sendMailRes, err := svc.SendCode(l.svcCtx.Config.SendCodeConfig, params)
	logx.Info("send code mail data ", sendMailRes, err)
	if err != nil {
		result.Code = http.StatusInternalServerError
		result.Message = messages.SendMailErr
		return &result, nil
	}
	err = l.svcCtx.KvStore.SetexCtx(l.ctx, codeLimitCacheKey, "1", 50)
	if err != nil {
		result.Code = http.StatusInternalServerError
		result.Message = messages.SaveFail
		return &result, nil
	}
	err = l.svcCtx.KvStore.SetexCtx(l.ctx, codeCacheKey, randCode, 900)
	if err != nil {
		result.Code = http.StatusInternalServerError
		result.Message = messages.SaveFail
		return &result, nil
	}
	result.Code = http.StatusOK
	result.Message = messages.Success
	return &result, nil
}
