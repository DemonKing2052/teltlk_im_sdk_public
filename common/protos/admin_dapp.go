package protos

import "time"

type GetManageBannerListReq struct {
	Pagination
}
type GetManageBannerListResp struct {
	Id             int64      `json:"id" `              // 主键ID
	Position       string     `json:"position" `        // 位置 key
	Title          string     `json:"title" `           // 标题
	Sort           int64      `json:"sort" `            // 顺序
	SkipUrl        string     `json:"skip_url" `        // 跳转URL
	ImgHref        string     `json:"img_href" `        // 图片地址
	Status         uint64     `json:"status" `          // 状态：1-上架，2-下架
	NeedLogin      uint64     `json:"need_login" `      // 是否需要登录：1-不需要，2-需要
	SkipTarget     uint64     `json:"skip_target" `     // 跳转目标：1-H5,2-客服
	StartAt        *time.Time `json:"start_at" `        // 开始时间
	ExpirationTime *time.Time `json:"expiration_time" ` // 过期时间
	CreatedTime    time.Time  `json:"created_time" `    // 创建时间
	UpdatedTime    time.Time  `json:"updated_time" `    // 更新时间

}

type ManageBannerOperationReq struct {
	OperationType  int64  `json:"operation_type" `  // 操作类型：1-新增，2-编辑，3-删除
	Id             int64  `json:"id" `              // 主键ID
	Position       string `json:"position" `        // 位置 key
	Title          string `json:"title" `           // 标题
	Sort           int64  `json:"sort" `            // 顺序
	SkipUrl        string `json:"skip_url" `        // 跳转URL
	ImgHref        string `json:"img_href" `        // 图片地址
	Status         uint64 `json:"status" `          // 状态：1-上架，2-下架
	NeedLogin      uint64 `json:"need_login" `      // 是否需要登录：1-不需要，2-需要
	SkipTarget     uint64 `json:"skip_target" `     // 跳转目标：1-H5,2-客服
	StartAt        string `json:"start_at" `        // 开始时间
	ExpirationTime string `json:"expiration_time" ` // 过期时间

}
type ManageBannerOperationResp struct {
}

type GetManageDiscoverToolCategoriesListReq struct {
	Pagination
}
type GetManageDiscoverToolCategoriesListResp struct {
	Id          int64     `json:"id" `           // 主键ID
	Title       string    `json:"title" `        // 标题
	Sort        int64     `json:"sort" `         // 顺序
	SkipUrl     string    `json:"skip_url" `     // 跳转URL
	ImgHref     string    `json:"img_href" `     // 图片地址
	Description string    `json:"description" `  // 说明
	Status      uint64    `json:"status" `       // 状态：1-上架，2-下架
	IsShow      int64     `json:"is_show" `      // 1-显示，2-隐藏
	CreatedTime time.Time `json:"created_time" ` // 创建时间
	UpdatedTime time.Time `json:"updated_time" ` // 更新时间
}
type ManageDiscoverToolCategoriesOperationReq struct {
	OperationType int64  `json:"operation_type" ` // 操作类型：1-新增，2-编辑，3-删除
	Id            int64  `json:"id" `             // 主键ID
	Title         string `json:"title"`           // 标题
	Sort          int64  `json:"sort" `           // 顺序
	SkipUrl       string `json:"skip_url" `       // 跳转URL
	ImgHref       string `json:"img_href" `       // 图片地址
	Description   string `json:"description" `    // 说明
	Status        uint64 `json:"status" `         // 状态：1-上架，2-下架
	IsShow        int64  `json:"is_show" `        // 1-显示，2-隐藏
}
type ManageDiscoverToolCategoriesOperationResp struct {
}

type GetManageNetworkListReq struct {
	Pagination
}
type GetManageNetworkListResp struct {
	Id          int64     `json:"id" `
	Name        string    `json:"name" `         // 网络名称，如 Ethereum、BSC、Solana
	ChainId     string    `json:"chain_id" `     // 链ID，如 1、56、solana
	Symbol      string    `json:"symbol" `       // 链主币符号，例如 ETH、BNB、SOL
	RpcUrl      string    `json:"rpc_url" `      // 默认 RPC，非必填
	ExplorerUrl string    `json:"explorer_url" ` // 浏览器地址，例如 https://bscscan.com
	Logo        string    `json:"logo" `         // 图标
	Status      int64     `json:"status" `       // 状态 1启用 2禁用
	Sort        int64     `json:"sort" `         // 排序
	CreatedAt   time.Time `json:"created_at" `
}

type ManageNetworkOperationReq struct {
	OperationType int64  `json:"operation_type" ` // 操作类型：1-新增，2-编辑，3-删除
	Id            int64  `json:"id" `             // 主键ID
	Name          string `json:"name" `           // 网络名称，如 Ethereum、BSC、Solana
	ChainId       string `json:"chain_id" `       // 链ID，如 1、56、solana
	Symbol        string `json:"symbol" `         // 链主币符号，例如 ETH、BNB、SOL
	RpcUrl        string `json:"rpc_url" `        // 默认 RPC，非必填
	ExplorerUrl   string `json:"explorer_url" `   // 浏览器地址，例如 https://bscscan.com
	Logo          string `json:"logo" `           // 图标
	Status        int64  `json:"status" `         // 状态 1启用 2禁用
	Sort          int64  `json:"sort" `           // 排序
}
type ManageNetworkOperationResp struct {
}
type GetManageDiscoverToolInfoListReq struct {
	CategoriesId int64 `json:"categories_id"` //分类ID
	Pagination
}
type ManageCategoryInfoItem struct {
	Id      int64  `json:"id" `       // 分类ID
	Title   string `json:"title" `    // 分类标题
	ImgHref string `json:"img_href" ` // 图片地址
}
type ManageNetWorkInfoItem struct {
	Id          int64  `json:"id" `
	Name        string `json:"name" `
	ChainId     string `json:"chain_id" `
	Symbol      string `json:"symbol" `
	RpcUrl      string `json:"rpc_url" `      // 默认 RPC，非必填
	ExplorerUrl string `json:"explorer_url" ` // 浏览器地址，例如 https://bscscan.com
	Logo        string `json:"logo" `         // 图标
}
type GetManageDiscoverToolInfoListResp struct {
	Id                    int64                    `json:"id" `
	CategoryInfoIds       []ManageCategoryInfoItem `json:"category_info_ids" `        // 关联分类ID列表
	Title                 string                   `json:"title" `                    // 标题
	ShortDesc             string                   `json:"short_desc" `               // 简短描述
	Thumbnail             string                   `json:"thumbnail" `                // 缩略图 URL
	ImgHref               string                   `json:"img_href" `                 // 图片地址
	SkipUrl               string                   `json:"skip_url" `                 // 跳转URL
	ServiceSupport        string                   `json:"service_support" `          // 服务支持内容
	SupportNetworkInfoIds []ManageNetWorkInfoItem  `json:"support_network_info_ids" ` // 支持的网络ID列表
	Token                 string                   `json:"token" `                    // token跳转URL
	CommunityTutorial     string                   `json:"community_tutorial" `       // 社区教程跳转URL
	Status                int64                    `json:"status" `                   // 1-显示 2-隐藏
	Sort                  int64                    `json:"sort" `                     // 排序
	Tags                  []string                 `json:"tags" `                     // 标签
	CreatedAt             time.Time                `json:"created_at" `
	UpdatedAt             time.Time                `json:"updated_at" `
}

type ManageDiscoverToolInfoOperationReq struct {
	OperationType     int64    `json:"operation_type" ` // 操作类型：1-新增，2-编辑，3-删除
	Id                int64    `json:"id" `
	CategoryIds       []int64  `json:"category_ids" `        // 关联分类ID列表
	Title             string   `json:"title" `               // 标题
	ShortDesc         string   `json:"short_desc" `          // 简短描述
	Thumbnail         string   `json:"thumbnail" `           // 缩略图 URL
	ImgHref           string   `json:"img_href" `            // 图片地址
	SkipUrl           string   `json:"skip_url" `            // 跳转URL
	ServiceSupport    string   `json:"service_support" `     // 服务支持内容
	SupportNetworkIds []int64  `json:"support_network_ids" ` // 支持的网络ID列表
	Token             string   `json:"token" `               // token跳转URL
	CommunityTutorial string   `json:"community_tutorial" `  // 社区教程跳转URL
	Status            int64    `json:"status" `              // 1-显示 2-隐藏
	Sort              int64    `json:"sort" `                // 排序
	Tags              []string `json:"tags" `                // 标签: hot-热门 new-最新 top - 置顶
}
type ManageDiscoverToolInfoOperationResp struct {
}

type GetManageDiscovertoolbarListReq struct {
	Pagination
}

type GetManageDiscovertoolbarListResp struct {
	Id      int64  `json:"id" `
	Title   string `json:"title" `    // 标题
	Tag     string `json:"tag" `      // 标签: hot-热门 new-最新
	ImgHref string `json:"img_href" ` // 图片地址
}

type ManageDiscovertoolbarOperationReq struct {
	OperationType int64  `json:"operation_type" ` // 操作类型：1-新增，2-编辑，3-删除
	Id            int64  `json:"id" `
	Title         string `json:"title" `
	Tag           string `json:"tag" `
	ImgHref       string `json:"img_href" `
	Sort          int64  `json:"sort" `
}
type ManageDiscovertoolbarOperationResp struct {
}
