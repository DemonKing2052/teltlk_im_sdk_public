package model

import (
	"ImSdk/common/utils"
	"context"
	"fmt"
	"gorm.io/gorm"
)

var _ DappDiscoverToolFavoritesModel = (*customDappDiscoverToolFavoritesModel)(nil)

type (
	// DappDiscoverToolFavoritesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDappDiscoverToolFavoritesModel.
	DappDiscoverToolFavoritesModel interface {
		dappDiscoverToolFavoritesModel
	}

	customDappDiscoverToolFavoritesModel struct {
		*defaultDappDiscoverToolFavoritesModel
	}
)

// NewDappDiscoverToolFavoritesModel returns a model for the database table.
func NewDappDiscoverToolFavoritesModel(conn *gorm.DB) DappDiscoverToolFavoritesModel {
	return &customDappDiscoverToolFavoritesModel{
		defaultDappDiscoverToolFavoritesModel: newDappDiscoverToolFavoritesModel(conn),
	}
}

func (m *defaultDappDiscoverToolFavoritesModel) DiscoverToolFavoritesFindListPage(ctx context.Context, userId string, networkId int64, page *utils.PageData) (*[]DappDiscoverToolFavorites, int64, error) {
	where := "1 = 1"
	var args []interface{}

	if userId != "" {
		where += " AND user_id = ? "
		args = append(args, userId)
	}
	if networkId > 0 {
		where += " AND JSON_CONTAINS(support_network_ids, ?, '$')"
		args = append(args, fmt.Sprintf("[%d]", networkId))
	}
	var resp []DappDiscoverToolFavorites
	var count int64
	err := m.conn.WithContext(ctx).Model(&DappDiscoverToolFavorites{}).Where(where, args...).Count(&count).Limit(page.PageSize).Offset(page.Offset).Find(&resp).Error
	switch err {
	case nil:
		return &resp, count, nil
	case gorm.ErrRecordNotFound:
		return nil, 0, gorm.ErrRecordNotFound
	default:
		return nil, 0, err
	}
}
