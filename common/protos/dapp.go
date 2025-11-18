package protos

import "time"

type GetBannerListReq struct {
	Pagination
}
type GetBannerListResp struct {
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

type GetDiscoverToolCategoriesListReq struct {
	Pagination
}
type GetDiscoverToolCategoriesListResp struct {
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

type GetNetworkListReq struct {
	Pagination
}
type GetNetworkListResp struct {
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

type GetDiscoverToolInfoListReq struct {
	UserId       string `json:"user_id"`
	ToolName     string `json:"tool_name"`     //工具名称
	NetWorkId    int64  `json:"net_work_id"`   //网络ID
	CategoriesId int64  `json:"categories_id"` //分类ID
	Tag          string `json:"tag"`           //标签
	Pagination
}
type CategoryInfoItem struct {
	Id      int64  `json:"id" `       // 分类ID
	Title   string `json:"title" `    // 分类标题
	ImgHref string `json:"img_href" ` // 图片地址
}
type NetWorkInfoItem struct {
	Id          int64  `json:"id" `
	Name        string `json:"name" `
	ChainId     string `json:"chain_id" `
	Symbol      string `json:"symbol" `
	RpcUrl      string `json:"rpc_url" `      // 默认 RPC，非必填
	ExplorerUrl string `json:"explorer_url" ` // 浏览器地址，例如 https://bscscan.com
	Logo        string `json:"logo" `         // 图标
}
type GetDiscoverToolInfoListResp struct {
	Id                    int64              `json:"id" `
	CategoryInfoIds       []CategoryInfoItem `json:"category_info_ids" `        // 关联分类ID列表
	Title                 string             `json:"title" `                    // 标题
	ShortDesc             string             `json:"short_desc" `               // 简短描述
	Thumbnail             string             `json:"thumbnail" `                // 缩略图 URL
	ImgHref               string             `json:"img_href" `                 // 图片地址
	SkipUrl               string             `json:"skip_url" `                 // 跳转URL
	ServiceSupport        string             `json:"service_support" `          // 服务支持内容
	SupportNetworkInfoIds []NetWorkInfoItem  `json:"support_network_info_ids" ` // 支持的网络ID列表
	Token                 string             `json:"token" `                    // token跳转URL
	CommunityTutorial     string             `json:"community_tutorial" `       // 社区教程跳转URL
	Status                int64              `json:"status" `                   // 1-显示 2-隐藏
	Sort                  int64              `json:"sort" `                     // 排序
	Tags                  []string           `json:"tags" `                     // 标签: hot-热门 new-最新 top - 置顶
	IsFavorites           int64              `json:"is_favorites"`              // 收藏状态：1收藏，2未收藏
	CreatedAt             time.Time          `json:"created_at" `
	UpdatedAt             time.Time          `json:"updated_at" `
}

type GetDiscoverToolbarListReq struct {
	Pagination
}
type GetDiscoverToolbarListResp struct {
	Id        int64     `json:"id" `
	Tag       string    `json:"tag" `      // 标签
	Title     string    `json:"title" `    // 标题
	ImgHref   string    `json:"img_href" ` // 图片地址
	Sort      int64     `json:"sort" `     // 排序
	CreatedAt time.Time `json:"created_at" `
}
type GetDiscoverToolFavoritesListReq struct {
	NetWorkId int64 `json:"net_work_id"` // 网络ID
	Pagination
}
type FavoritesCategoryInfoItem struct {
	Id      int64  `json:"id" `       // 分类ID
	Title   string `json:"title" `    // 分类标题
	ImgHref string `json:"img_href" ` // 图片地址
}
type FavoritesNetWorkInfoItem struct {
	Id          int64  `json:"id" `
	Name        string `json:"name" `
	ChainId     string `json:"chain_id" `
	Symbol      string `json:"symbol" `
	RpcUrl      string `json:"rpc_url" `      // 默认 RPC，非必填
	ExplorerUrl string `json:"explorer_url" ` // 浏览器地址，例如 https://bscscan.com
	Logo        string `json:"logo" `         // 图标
}

type GetDiscoverToolFavoritesListResp struct {
	Id                    int64                       `json:"id" `
	ToolId                int64                       `json:"tool_id" `
	CategoryInfoIds       []FavoritesCategoryInfoItem `json:"category_info_ids" `        // 关联分类ID列表
	Title                 string                      `json:"title" `                    // 标题
	ShortDesc             string                      `json:"short_desc" `               // 简短描述
	Thumbnail             string                      `json:"thumbnail" `                // 缩略图 URL
	ImgHref               string                      `json:"img_href" `                 // 图片地址
	SkipUrl               string                      `json:"skip_url" `                 // 跳转URL
	ServiceSupport        string                      `json:"service_support" `          // 服务支持内容
	SupportNetworkInfoIds []FavoritesNetWorkInfoItem  `json:"support_network_info_ids" ` // 支持的网络ID列表
	Token                 string                      `json:"token" `                    // token跳转URL
	CommunityTutorial     string                      `json:"community_tutorial" `       // 社区教程跳转URL
	Status                int64                       `json:"status" `                   // 1-显示 2-隐藏
	Sort                  int64                       `json:"sort" `                     // 排序
	CreatedAt             time.Time                   `json:"created_at" `
}

type DiscoverToolFavoritesOperationReq struct {
	ToolId int64 `json:"tool_id" ` //
	Status int64 `json:"status" `  // 1-显示 2-隐藏
}
type DiscoverToolFavoritesOperationResp struct{}

type GetDiscoverToolEventListReq struct {
	EventType string `json:"event_type" ` // view/click/share/fav/comment
	Pagination
}

type EventCategoryInfoItem struct {
	Id      int64  `json:"id" `       // 分类ID
	Title   string `json:"title" `    // 分类标题
	ImgHref string `json:"img_href" ` // 图片地址
}
type EventNetWorkInfoItem struct {
	Id          int64  `json:"id" `
	Name        string `json:"name" `
	ChainId     string `json:"chain_id" `
	Symbol      string `json:"symbol" `
	RpcUrl      string `json:"rpc_url" `      // 默认 RPC，非必填
	ExplorerUrl string `json:"explorer_url" ` // 浏览器地址，例如 https://bscscan.com
	Logo        string `json:"logo" `         // 图标
}

type GetDiscoverToolEventListResp struct {
	Id                    int64                   `json:"id" `
	ToolId                int64                   `json:"tool_id" `
	CategoryInfoIds       []EventCategoryInfoItem `json:"category_info_ids" `        // 关联分类ID列表
	Title                 string                  `json:"title" `                    // 标题
	ShortDesc             string                  `json:"short_desc" `               // 简短描述
	Thumbnail             string                  `json:"thumbnail" `                // 缩略图 URL
	ImgHref               string                  `json:"img_href" `                 // 图片地址
	SkipUrl               string                  `json:"skip_url" `                 // 跳转URL
	ServiceSupport        string                  `json:"service_support" `          // 服务支持内容
	SupportNetworkInfoIds []EventNetWorkInfoItem  `json:"support_network_info_ids" ` // 支持的网络ID列表
	Token                 string                  `json:"token" `                    // token跳转URL
	CommunityTutorial     string                  `json:"community_tutorial" `       // 社区教程跳转URL
	Status                int64                   `json:"status" `                   // 1-显示 2-隐藏
	Sort                  int64                   `json:"sort" `                     // 排序
	CreatedAt             time.Time               `json:"created_at" `
}

type DiscoverToolEventOperationReq struct {
	EventType string `json:"event_type" ` // view/click/share/fav/comment
	ToolId    int64  `json:"tool_id" `
}

type DiscoverToolEventOperationResp struct{}
