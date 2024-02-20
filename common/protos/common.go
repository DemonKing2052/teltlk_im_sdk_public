package protos

type ImCommonResp struct {
	ErrCode int64  `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
	Data    any    `json:"data"`
}
