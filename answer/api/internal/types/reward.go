package types

type SendMailParams struct {
	AppId      int     `json:"app_id"`
	SvrRegion  string  `json:"svr_region"`
	Type       string  `json:"type,default=send_mail"`
	SvrId      int     `json:"svr_id"`
	SenderName string  `json:"sender_name"`
	RecverUids []int64 `json:"recver_uids"`
	Title      string  `json:"title"`
	Context    string  `json:"context"`
	Currency   int     `json:"currency,omitempty"`
	Coins      []Coin  `json:"coins"`
	Items      []Item  `json:"items,omitempty"`
	AffectTime int64   `json:"affectTime"`
	ExpireTime int64   `json:"expireTime"`
	Assets     []Asset `json:"assets,omitempty"`
}
type Coin struct {
	CoinName  string `json:"coin_name"`
	CoinValue int    `json:"coin_value"`
}
type Item struct {
	ItemId    interface{} `json:"item_id"`
	ItemCount int64       `json:"item_count"`
}
type Asset struct {
	Type   string `json:"type"`
	Id     string `json:"id"`
	Amount int    `json:"amount"`
}

type SendMailResp struct {
	Result int        `json:"result"`
	Msg    string     `json:"msg"`
	TaskId string     `json:"taskId"`
	Data   []GameResp `json:"data"`
}

type GameResp struct {
	Ret int `json:"ret"`
}
