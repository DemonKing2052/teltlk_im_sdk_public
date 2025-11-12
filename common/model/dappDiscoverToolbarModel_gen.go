package model

import (
	"ImSdk/common/utils"
	"context"
	"time"

	"gorm.io/gorm"
)

var ()

type (
	dappDiscoverToolbarModel interface {
		Insert(ctx context.Context, data *DappDiscoverToolbar) error
		FindOne(ctx context.Context, id int64) (*DappDiscoverToolbar, error)
		Update(ctx context.Context, session *gorm.DB, data *DappDiscoverToolbar) error
		Delete(ctx context.Context, session *gorm.DB, id int64) error
		DiscoverToolbarFindListPage(ctx context.Context, page *utils.PageData) (*[]DappDiscoverToolbar, int64, error)
	}

	defaultDappDiscoverToolbarModel struct {
		conn  *gorm.DB
		table string
	}

	DappDiscoverToolbar struct {
		Id        int64     `json:"id" gorm:"column:id"`
		Title     string    `json:"title" gorm:"column:title"`       // 标题
		Tag       string    `json:"tag" gorm:"column:tag"`           // hot
		ImgHref   string    `json:"img_href" gorm:"column:img_href"` // 图片地址
		Sort      int64     `json:"sort" gorm:"column:sort"`         //排序
		CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	}
)

func newDappDiscoverToolbarModel(conn *gorm.DB) *defaultDappDiscoverToolbarModel {
	return &defaultDappDiscoverToolbarModel{
		conn:  conn,
		table: "`dapp_discover_toolbar`",
	}
}

func (m *defaultDappDiscoverToolbarModel) Insert(ctx context.Context, data *DappDiscoverToolbar) error {
	err := m.conn.WithContext(ctx).Table(m.table).Create(&data).Error
	return err
}

func (m *defaultDappDiscoverToolbarModel) FindOne(ctx context.Context, id int64) (*DappDiscoverToolbar, error) {
	var resp DappDiscoverToolbar
	err := m.conn.WithContext(ctx).Table(m.table).Where("`id` = ?", id).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDappDiscoverToolbarModel) Update(ctx context.Context, session *gorm.DB, data *DappDiscoverToolbar) error {
	err := m.conn.WithContext(ctx).Table(m.table).Updates(data).Error
	return err
}

func (m *defaultDappDiscoverToolbarModel) Delete(ctx context.Context, session *gorm.DB, id int64) error {
	err := m.conn.WithContext(ctx).Table(m.table).Where("`id` = ?", id).Delete(nil).Error

	return err
}

func (m *defaultDappDiscoverToolbarModel) tableName() string {
	return m.table
}
