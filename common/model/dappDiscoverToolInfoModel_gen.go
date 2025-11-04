package model

import (
	"ImSdk/common/utils"
	"context"
	"time"

	"gorm.io/gorm"
)

var ()

type (
	dappDiscoverToolInfoModel interface {
		Insert(ctx context.Context, data *DappDiscoverToolInfo) error
		FindOne(ctx context.Context, id int64) (*DappDiscoverToolInfo, error)
		Update(ctx context.Context, session *gorm.DB, data *DappDiscoverToolInfo) error
		Delete(ctx context.Context, session *gorm.DB, id int64) error
		DiscoverToolInfoFindListPage(ctx context.Context, toolName string, netWorkId, categoriesId int64, tag string, page *utils.PageData) (*[]DappDiscoverToolInfo, int64, error)
	}

	defaultDappDiscoverToolInfoModel struct {
		conn  *gorm.DB
		table string
	}

	DappDiscoverToolInfo struct {
		Id                int64     `json:"id" gorm:"column:id"`
		CategoryIds       []int64   `json:"category_ids" gorm:"column:category_ids"`               // 关联分类ID列表
		Title             string    `json:"title" gorm:"column:title"`                             // 标题
		ShortDesc         string    `json:"short_desc" gorm:"column:short_desc"`                   // 简短描述
		Thumbnail         string    `json:"thumbnail" gorm:"column:thumbnail"`                     // 缩略图 URL
		ImgHref           string    `json:"img_href" gorm:"column:img_href"`                       // 图片地址
		SkipUrl           string    `json:"skip_url" gorm:"column:skip_url"`                       // 跳转URL
		ServiceSupport    string    `json:"service_support" gorm:"column:service_support"`         // 服务支持内容
		SupportNetworkIds []int64   `json:"support_network_ids" gorm:"column:support_network_ids"` // 支持的网络ID列表
		Token             string    `json:"token" gorm:"column:token"`                             // token跳转URL
		CommunityTutorial string    `json:"community_tutorial" gorm:"column:community_tutorial"`   // 社区教程跳转URL
		Status            int64     `json:"status" gorm:"column:status"`                           // 1-显示 2-隐藏
		Sort              int64     `json:"sort" gorm:"column:sort"`                               // 排序
		Tags              []string  `json:"tags" gorm:"column:tags"`                               // 标签: hot-热门 new-最新 top - 置顶
		CreatedAt         time.Time `json:"created_at" gorm:"column:created_at"`
		UpdatedAt         time.Time `json:"updated_at" gorm:"column:updated_at"`
	}
)

func newDappDiscoverToolInfoModel(conn *gorm.DB) *defaultDappDiscoverToolInfoModel {
	return &defaultDappDiscoverToolInfoModel{
		conn:  conn,
		table: "`dapp_discover_tool_info`",
	}
}

func (m *defaultDappDiscoverToolInfoModel) Insert(ctx context.Context, data *DappDiscoverToolInfo) error {
	err := m.conn.WithContext(ctx).Create(&data).Error
	return err
}

func (m *defaultDappDiscoverToolInfoModel) FindOne(ctx context.Context, id int64) (*DappDiscoverToolInfo, error) {
	var resp DappDiscoverToolInfo
	err := m.conn.WithContext(ctx).Model(&DappDiscoverToolInfo{}).Where("`id` = ?", id).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case gorm.ErrRecordNotFound:
		return nil, gorm.ErrRecordNotFound
	default:
		return nil, err
	}
}

func (m *defaultDappDiscoverToolInfoModel) Update(ctx context.Context, session *gorm.DB, data *DappDiscoverToolInfo) error {
	err := m.conn.WithContext(ctx).Updates(data).Error
	return err
}

func (m *defaultDappDiscoverToolInfoModel) Delete(ctx context.Context, session *gorm.DB, id int64) error {
	err := m.conn.WithContext(ctx).Delete(&DappDiscoverToolInfo{}, id).Error

	return err
}

func (m *defaultDappDiscoverToolInfoModel) tableName() string {
	return m.table
}
