package model

import (
	"ImSdk/common/utils"
	"context"
	"gorm.io/gorm"
)

var _ DappDiscoverToolEventsModel = (*customDappDiscoverToolEventsModel)(nil)

type (
	// DappDiscoverToolEventsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDappDiscoverToolEventsModel.
	DappDiscoverToolEventsModel interface {
		dappDiscoverToolEventsModel
	}

	customDappDiscoverToolEventsModel struct {
		*defaultDappDiscoverToolEventsModel
	}
)

// NewDappDiscoverToolEventsModel returns a model for the database table.
func NewDappDiscoverToolEventsModel(conn *gorm.DB) DappDiscoverToolEventsModel {
	return &customDappDiscoverToolEventsModel{
		defaultDappDiscoverToolEventsModel: newDappDiscoverToolEventsModel(conn),
	}
}

func (m *defaultDappDiscoverToolEventsModel) FindOneByUserIdEventTypeToolId(ctx context.Context, userId, eventType string, toolId int64) (*DappDiscoverToolEvents, error) {
	var resp DappDiscoverToolEvents
	err := m.conn.WithContext(ctx).Table(m.table).Where("`user_id` = ? and event_type = ? and `tool_id` = ?", userId, eventType, toolId).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case gorm.ErrRecordNotFound:
		return nil, gorm.ErrRecordNotFound
	default:
		return nil, err
	}
}
func (m *defaultDappDiscoverToolEventsModel) DiscoverToolEventsFindListPage(ctx context.Context, userId, eventType string, page *utils.PageData) (*[]DappDiscoverToolEvents, int64, error) {
	var resp []DappDiscoverToolEvents
	var count int64
	err := m.conn.WithContext(ctx).Model(&DappDiscoverToolEvents{}).Where("`user_id` = ? and event_type = ?", userId, eventType).Count(&count).Limit(page.PageSize).Offset(page.Offset).Find(&resp).Error
	switch err {
	case nil:
		return &resp, count, nil
	case gorm.ErrRecordNotFound:
		return nil, 0, gorm.ErrRecordNotFound
	default:
		return nil, 0, err
	}
}
