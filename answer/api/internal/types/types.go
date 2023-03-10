// Code generated by goctl. DO NOT EDIT.
package types

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SendCodeRequest struct {
	GameUid string `json:"gameUid"`
}

type SendCodeResponse struct {
	Response
}

type LoginRequest struct {
	GameUid  string `json:"gameUid"`
	AuthCode string `json:"authCode"`
	Language string `json:"language"`
}

type PageRequest struct {
	GameUid string `json:"gameUid"`
	PageUrl string `json:"pageUrl"`
}

type QueryTargetRequest struct {
	Sql string `json:"sql"`
}

type LoginData struct {
	GameUid         string   `json:"gameUid"`
	Token           string   `json:"token"`
	AccessExpire    int64    `json:"accessExpire"`
	RefreshAfter    int64    `json:"refreshAfter"`
	FragmentRewards []string `json:"fragmentRewards"`
	Language        string   `json:"language"`
}

type LoginResponse struct {
	Response
	Data LoginData `json:"data"`
}

type AnswerListData struct {
	AnswerId     int64  `json:"answerId"`
	AnwerName    string `json:"anwerName"`
	AnserStatus  int    `json:"anserStatus"`
	RewardStatus int    `json:"rewardStatus"`
	OpenTime     string `json:"openTime"`
	SurplusTimes int    `json:"surplusTimes"`
}

type AnswerListResponse struct {
	Response
	Data []AnswerListData `json:"data"`
}

type StartAnswerRequest struct {
	AnswerId int64 `json:"answerId"`
}

type StartAnswerData struct {
	SurplusTimes int `json:"surplusTimes"`
}

type StartAnswerResponse struct {
	Response
	Data StartAnswerData `json:"data"`
}

type AnswerRequest struct {
	AnswerId int64  `json:"answerId"`
	Result   string `json:"result"`
}

type AnswerData struct {
	ItemId       string `json:"itemId"`
	SurplusTimes int    `json:"surplusTimes"`
}

type AnswerReponse struct {
	Response
	Data AnswerData `json:"data"`
}

type RewardRequest struct {
	AnswerId   int64  `json:"answerId"`
	SenderName string `json:"senderName"`
	Title      string `json:"title"`
	Context    string `json:"context"`
}
