package types

type SendCodeParams struct {
	Type string `json:"type,default=mail"`
	Uid  string `json:"uid"`
	Code string `json:"code"`
}
type SendCodeReq struct {
	Type string `json:"type,default=mail"`
	Uid  int64  `json:"uid"`
	Code int64  `json:"code"`
	Sign string `json:"sign"`
}

type SendCodeRes struct {
	Result      string   `json:"result"`
	CreateTime  int64    `json:"create_time"`
	Name        string   `json:"name"`
	Power       int64    `json:"power"`
	OfficerInfo []string `json:"officer_info"`
}
