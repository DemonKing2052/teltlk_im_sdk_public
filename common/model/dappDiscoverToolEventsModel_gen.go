package model

import (
	"ImSdk/common/utils"
	"context"
	"time"

	"gorm.io/gorm"
)

var ()

type (
	dappDiscoverToolEventsModel interface {
		Insert(ctx context.Context, data *DappDiscoverToolEvents) error
		FindOne(ctx context.Context, id int64) (*DappDiscoverToolEvents, error)
		FindOneByUserIdEventTypeToolId(ctx context.Context, userId, eventType string, toolId int64) (*DappDiscoverToolEvents, error)
		Update(ctx context.Context, session *gorm.DB, data *DappDiscoverToolEvents) error
		Delete(ctx context.Context, session *gorm.DB, id int64) error
		DiscoverToolEventsFindListPage(ctx context.Context, userId, eventType string, page *utils.PageData) (*[]DappDiscoverToolEvents, int64, error)
	}

	defaultDappDiscoverToolEventsModel struct {
		conn  *gorm.DB
		table string
	}

	DappDiscoverToolEvents struct {
		Id        int64     `json:"id" gorm:"column:id"`
		UserId    string    `json:"user_id" gorm:"column:user_id"`
		ItemId    int64     `json:"item_id" gorm:"column:item_id"`
		EventType string    `json:"event_type" gorm:"column:event_type"` // view/click/share/fav/comment
		Ip        string    `json:"ip" gorm:"column:ip"`                 // IP
		Ua        string    `json:"ua" gorm:"column:ua"`                 // 设备 UA
		Meta      string    `json:"meta" gorm:"column:meta"`             // 额外数据，如按钮、入口、来源
		CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	}
)

func newDappDiscoverToolEventsModel(conn *gorm.DB) *defaultDappDiscoverToolEventsModel {
	return &defaultDappDiscoverToolEventsModel{
		conn:  conn,
		table: "`dapp_discover_tool_events`",
	}
}

func (m *defaultDappDiscoverToolEventsModel) Insert(ctx context.Context, data *DappDiscoverToolEvents) error {
	err := m.conn.WithContext(ctx).Create(&data).Error
	return err
}

func (m *defaultDappDiscoverToolEventsModel) FindOne(ctx context.Context, id int64) (*DappDiscoverToolEvents, error) {
	var resp DappDiscoverToolEvents
	err := m.conn.WithContext(ctx).Model(&DappDiscoverToolEvents{}).Where("`id` = ?", id).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case gorm.ErrRecordNotFound:
		return nil, gorm.ErrRecordNotFound
	default:
		return nil, err
	}
}

func (m *defaultDappDiscoverToolEventsModel) Update(ctx context.Context, session *gorm.DB, data *DappDiscoverToolEvents) error {
	err := m.conn.WithContext(ctx).Updates(data).Error
	return err
}

func (m *defaultDappDiscoverToolEventsModel) Delete(ctx context.Context, session *gorm.DB, id int64) error {
	err := m.conn.WithContext(ctx).Delete(&DappDiscoverToolEvents{}, id).Error

	return err
}

func (m *defaultDappDiscoverToolEventsModel) tableName() string {
	return m.table
}
