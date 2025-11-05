package model

import (
	"ImSdk/common/utils"
	"context"
	"gorm.io/gorm"
)

var _ DappBannerModel = (*customDappBannerModel)(nil)

type (
	// DappBannerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDappBannerModel.
	DappBannerModel interface {
		dappBannerModel
	}

	customDappBannerModel struct {
		*defaultDappBannerModel
	}
)

// NewDappBannerModel returns a model for the database table.
func NewDappBannerModel(conn *gorm.DB) DappBannerModel {
	return &customDappBannerModel{
		defaultDappBannerModel: newDappBannerModel(conn),
	}
}

func (m *defaultDappBannerModel) BannerFindListPage(ctx context.Context, page *utils.PageData) (*[]DappBanner, int64, error) {
	var resp []DappBanner
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
