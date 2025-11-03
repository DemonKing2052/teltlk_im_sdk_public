package protos

type ImCommonResp struct {
	ErrCode int64  `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
	Data    any    `json:"data"`
}

// 通用翻页
type Pagination struct {
	Page     int64 `json:"page,optional,default=1" gorm:"-"`
	PageSize int64 `json:"page_size,optional,default=20" gorm:"-"`
}

// CommonListResp 通用列表模块
type CommonListResp struct {
	List     interface{} `json:"list"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	Total    int64       `json:"total"`
}
