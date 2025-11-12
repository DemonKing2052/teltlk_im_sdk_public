package model

import (
	"ImSdk/common/utils"
	"context"
	"time"

	"gorm.io/gorm"
)

var ()

type (
	dappDiscoverToolCategoriesModel interface {
		Insert(ctx context.Context, data *DappDiscoverToolCategories) error
		FindOne(ctx context.Context, id int64) (*DappDiscoverToolCategories, error)
		Update(ctx context.Context, session *gorm.DB, data *DappDiscoverToolCategories) error
		Delete(ctx context.Context, session *gorm.DB, id int64) error
		DiscoverToolCategoriesFindListPage(ctx context.Context, page *utils.PageData) (*[]DappDiscoverToolCategories, int64, error)
	}

	defaultDappDiscoverToolCategoriesModel struct {
		conn  *gorm.DB
		table string
	}

	DappDiscoverToolCategories struct {
		Id          int64     `json:"id" gorm:"column:id"`                     // 主键ID
		Title       string    `json:"title" gorm:"column:title"`               // 标题
		Sort        int64     `json:"sort" gorm:"column:sort"`                 // 顺序
		SkipUrl     string    `json:"skip_url" gorm:"column:skip_url"`         // 跳转URL
		ImgHref     string    `json:"img_href" gorm:"column:img_href"`         // 图片地址
		Description string    `json:"description" gorm:"column:description"`   // 说明
		Status      uint64    `json:"status" gorm:"column:status"`             // 状态：1-上架，2-下架
		IsShow      int64     `json:"is_show" gorm:"column:is_show"`           // 1-显示，2-隐藏
		CreatedTime time.Time `json:"created_time" gorm:"column:created_time"` // 创建时间
		UpdatedTime time.Time `json:"updated_time" gorm:"column:updated_time"` // 更新时间
	}
)

func newDappDiscoverToolCategoriesModel(conn *gorm.DB) *defaultDappDiscoverToolCategoriesModel {
	return &defaultDappDiscoverToolCategoriesModel{
		conn:  conn,
		table: "`dapp_discover_tool_categories`",
	}
}

func (m *defaultDappDiscoverToolCategoriesModel) Insert(ctx context.Context, data *DappDiscoverToolCategories) error {
	err := m.conn.WithContext(ctx).Table(m.table).Create(&data).Error
	return err
}

func (m *defaultDappDiscoverToolCategoriesModel) FindOne(ctx context.Context, id int64) (*DappDiscoverToolCategories, error) {
	var resp DappDiscoverToolCategories
	err := m.conn.WithContext(ctx).Table(m.table).Where("`id` = ?", id).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case gorm.ErrRecordNotFound:
		return nil, gorm.ErrRecordNotFound
	default:
		return nil, err
	}
}

func (m *defaultDappDiscoverToolCategoriesModel) Update(ctx context.Context, session *gorm.DB, data *DappDiscoverToolCategories) error {
	err := m.conn.WithContext(ctx).Table(m.table).Updates(data).Error
	return err
}

func (m *defaultDappDiscoverToolCategoriesModel) Delete(ctx context.Context, session *gorm.DB, id int64) error {
	err := m.conn.WithContext(ctx).Table(m.table).Where("`id` = ?", id).Delete(nil).Error

	return err
}

func (m *defaultDappDiscoverToolCategoriesModel) tableName() string {
	return m.table
}
