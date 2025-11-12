package model

import (
	"ImSdk/common/utils"
	"context"
	"time"

	"gorm.io/gorm"
)

var ()

type (
	dappDiscoverToolFavoritesModel interface {
		Insert(ctx context.Context, data *DappDiscoverToolFavorites) error
		FindOne(ctx context.Context, id int64) (*DappDiscoverToolFavorites, error)
		FindOneByUserIdToolId(ctx context.Context, userId string, toolId int64) (*DappDiscoverToolFavorites, error)
		Update(ctx context.Context, session *gorm.DB, data *DappDiscoverToolFavorites) error
		Delete(ctx context.Context, session *gorm.DB, id int64) error
		DiscoverToolFavoritesFindListPage(ctx context.Context, userId string, networkId int64, page *utils.PageData) (*[]DappDiscoverToolFavorites, int64, error)
	}

	defaultDappDiscoverToolFavoritesModel struct {
		conn  *gorm.DB
		table string
	}

	DappDiscoverToolFavorites struct {
		Id                int64     `json:"id" gorm:"column:id"`
		UserId            string    `json:"user_id" gorm:"column:user_id"`
		ToolId            int64     `json:"tool_id" gorm:"column:tool_id"`
		SupportNetworkIds []int64   `json:"support_network_ids" gorm:"serializer:json;column:support_network_ids"` // 支持的网络ID列表
		Status            int64     `json:"status" gorm:"column:status"`                                           // 1-显示 2-隐藏
		CreatedAt         time.Time `json:"created_at" gorm:"column:created_at"`
	}
)

func newDappDiscoverToolFavoritesModel(conn *gorm.DB) *defaultDappDiscoverToolFavoritesModel {
	return &defaultDappDiscoverToolFavoritesModel{
		conn:  conn,
		table: "`dapp_discover_tool_favorites`",
	}
}

func (m *defaultDappDiscoverToolFavoritesModel) Insert(ctx context.Context, data *DappDiscoverToolFavorites) error {
	err := m.conn.WithContext(ctx).Table(m.table).Create(&data).Error
	return err
}

func (m *defaultDappDiscoverToolFavoritesModel) FindOne(ctx context.Context, id int64) (*DappDiscoverToolFavorites, error) {
	var resp DappDiscoverToolFavorites
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

func (m *defaultDappDiscoverToolFavoritesModel) FindOneByUserIdToolId(ctx context.Context, userId string, toolId int64) (*DappDiscoverToolFavorites, error) {
	var resp DappDiscoverToolFavorites
	err := m.conn.WithContext(ctx).Table(m.table).Where("`user_id` = ? and `tool_id` = ?", userId, toolId).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case gorm.ErrRecordNotFound:
		return nil, gorm.ErrRecordNotFound
	default:
		return nil, err
	}
}

func (m *defaultDappDiscoverToolFavoritesModel) Update(ctx context.Context, session *gorm.DB, data *DappDiscoverToolFavorites) error {
	err := m.conn.WithContext(ctx).Table(m.table).Updates(data).Error
	return err
}

func (m *defaultDappDiscoverToolFavoritesModel) Delete(ctx context.Context, session *gorm.DB, id int64) error {
	err := m.conn.WithContext(ctx).Table(m.table).Where("`id` = ?", id).Delete(nil).Error

	return err
}

func (m *defaultDappDiscoverToolFavoritesModel) tableName() string {
	return m.table
}
