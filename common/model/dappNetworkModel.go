package model

import (
	"ImSdk/common/utils"
	"context"
	"gorm.io/gorm"
)

var _ DappNetworkModel = (*customDappNetworkModel)(nil)

type (
	// DappNetworkModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDappNetworkModel.
	DappNetworkModel interface {
		dappNetworkModel
	}

	customDappNetworkModel struct {
		*defaultDappNetworkModel
	}
)

// NewDappNetworkModel returns a model for the database table.
func NewDappNetworkModel(conn *gorm.DB) DappNetworkModel {
	return &customDappNetworkModel{
		defaultDappNetworkModel: newDappNetworkModel(conn),
	}
}

func (m *defaultDappNetworkModel) NetWorkFindListPage(ctx context.Context, page *utils.PageData) (*[]DappNetwork, int64, error) {
	var resp []DappNetwork
	var count int64
	err := m.conn.WithContext(ctx).Table(m.table).Count(&count).Limit(page.PageSize).Offset(page.Offset).Find(&resp).Error
	switch err {
	case nil:
		return &resp, count, nil
	case gorm.ErrRecordNotFound:
		return nil, 0, gorm.ErrRecordNotFound
	default:
		return nil, 0, err
	}
}
