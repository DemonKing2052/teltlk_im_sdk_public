package model

import (
	"ImSdk/common/utils"
	"context"
	"fmt"
	"gorm.io/gorm"
)

var _ DappDiscoverToolInfoModel = (*customDappDiscoverToolInfoModel)(nil)

type (
	// DappDiscoverToolInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDappDiscoverToolInfoModel.
	DappDiscoverToolInfoModel interface {
		dappDiscoverToolInfoModel
	}

	customDappDiscoverToolInfoModel struct {
		*defaultDappDiscoverToolInfoModel
	}
)

// NewDappDiscoverToolInfoModel returns a model for the database table.
func NewDappDiscoverToolInfoModel(conn *gorm.DB) DappDiscoverToolInfoModel {
	return &customDappDiscoverToolInfoModel{
		defaultDappDiscoverToolInfoModel: newDappDiscoverToolInfoModel(conn),
	}
}

func (m *defaultDappDiscoverToolInfoModel) DiscoverToolInfoFindListPage(ctx context.Context, toolName string, netWorkId, categoriesId int64, tag string, page *utils.PageData) (*[]DappDiscoverToolInfo, int64, error) {
	where := "1 = 1"
	var args []interface{}
	if toolName != "" {
		where += " and title like ?"
		args = append(args, "%"+toolName+"%")
	}
	if netWorkId > 0 {
		where += " AND JSON_CONTAINS(support_network_ids, ?, '$')"
		args = append(args, fmt.Sprintf("[%d]", netWorkId))
	}
	if categoriesId > 0 {
		where += " AND JSON_CONTAINS(category_ids, ?, '$')"
		args = append(args, fmt.Sprintf("[%d]", categoriesId))
	}
	if tag != "" {
		where += " AND JSON_CONTAINS(tags, ?, '$')"
		args = append(args, fmt.Sprintf("[%s]", tag))
	}
	var resp []DappDiscoverToolInfo
	var count int64
	err := m.conn.WithContext(ctx).Model(&DappDiscoverToolInfo{}).Count(&count).Limit(page.PageSize).Offset(page.Offset).Find(&resp).Error
	switch err {
	case nil:
		return &resp, count, nil
	case gorm.ErrRecordNotFound:
		return nil, 0, gorm.ErrRecordNotFound
	default:
		return nil, 0, err
	}
}
