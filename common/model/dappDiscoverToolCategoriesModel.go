package model

import (
	"ImSdk/common/utils"
	"context"
	"gorm.io/gorm"
)

var _ DappDiscoverToolCategoriesModel = (*customDappDiscoverToolCategoriesModel)(nil)

type (
	// DappDiscoverToolCategoriesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDappDiscoverToolCategoriesModel.
	DappDiscoverToolCategoriesModel interface {
		dappDiscoverToolCategoriesModel
	}

	customDappDiscoverToolCategoriesModel struct {
		*defaultDappDiscoverToolCategoriesModel
	}
)

// NewDappDiscoverToolCategoriesModel returns a model for the database table.
func NewDappDiscoverToolCategoriesModel(conn *gorm.DB) DappDiscoverToolCategoriesModel {
	return &customDappDiscoverToolCategoriesModel{
		defaultDappDiscoverToolCategoriesModel: newDappDiscoverToolCategoriesModel(conn),
	}
}

func (m *defaultDappDiscoverToolCategoriesModel) DiscoverToolCategoriesFindListPage(ctx context.Context, page *utils.PageData) (*[]DappDiscoverToolCategories, int64, error) {
	var resp []DappDiscoverToolCategories
	var count int64
	err := m.conn.WithContext(ctx).Model(&DappDiscoverToolCategories{}).Count(&count).Limit(page.PageSize).Offset(page.Offset).Find(&resp).Error
	switch err {
	case nil:
		return &resp, count, nil
	case gorm.ErrRecordNotFound:
		return nil, 0, gorm.ErrRecordNotFound
	default:
		return nil, 0, err
	}
}
