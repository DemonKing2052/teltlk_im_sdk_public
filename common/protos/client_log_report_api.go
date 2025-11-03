package protos

type (
	ClientLogReportingReq struct {
		ClientLog string `json:"client_log"`
	}
	ClientLogReportingResp struct {
		Data    any    `json:"data"`
		ErrCode int64  `json:"errCode"`
		ErrMsg  string `json:"errMsg"`
	}
)

type (
	GetClientLogReq struct {
		Pagination
	}
	GetClientLogResp struct {
		Data    any    `json:"data"`
		ErrCode int64  `json:"errCode"`
		ErrMsg  string `json:"errMsg"`
	}
)
