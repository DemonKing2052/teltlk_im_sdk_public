package model

import (
	"ImSdk/common/utils"
	"context"
	"gorm.io/gorm"
)

var _ DappDiscoverToolbarModel = (*customDappDiscoverToolbarModel)(nil)

type (
	// DappDiscoverToolbarModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDappDiscoverToolbarModel.
	DappDiscoverToolbarModel interface {
		dappDiscoverToolbarModel
	}

	customDappDiscoverToolbarModel struct {
		*defaultDappDiscoverToolbarModel
	}
)

// NewDappDiscoverToolbarModel returns a model for the database table.
func NewDappDiscoverToolbarModel(conn *gorm.DB) DappDiscoverToolbarModel {
	return &customDappDiscoverToolbarModel{
		defaultDappDiscoverToolbarModel: newDappDiscoverToolbarModel(conn),
	}
}

func (m *defaultDappDiscoverToolbarModel) DiscoverToolbarFindListPage(ctx context.Context, page *utils.PageData) (*[]DappDiscoverToolbar, int64, error) {
	var resp []DappDiscoverToolbar
	var count int64
	err := m.conn.WithContext(ctx).Table(m.table).Count(&count).Limit(page.PageSize).Offset(page.Offset).Order("`sort` DESC").Find(&resp).Error
	switch err {
	case nil:
		return &resp, count, nil
	case gorm.ErrRecordNotFound:
		return nil, 0, gorm.ErrRecordNotFound
	default:
		return nil, 0, err
	}
}
