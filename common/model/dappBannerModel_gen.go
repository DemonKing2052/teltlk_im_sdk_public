package model

import (
	"ImSdk/common/utils"
	"context"
	"time"

	"gorm.io/gorm"
)

var ()

type (
	dappBannerModel interface {
		Insert(ctx context.Context, data *DappBanner) error
		FindOne(ctx context.Context, id int64) (*DappBanner, error)
		Update(ctx context.Context, session *gorm.DB, data *DappBanner) error
		Delete(ctx context.Context, session *gorm.DB, id int64) error
		BannerFindListPage(ctx context.Context, page *utils.PageData) (*[]DappBanner, int64, error)
	}

	defaultDappBannerModel struct {
		conn  *gorm.DB
		table string
	}

	DappBanner struct {
		Id             int64      `json:"id" gorm:"column:id"`                           // 主键ID
		Position       string     `json:"position" gorm:"column:position"`               // 位置 key
		Title          string     `json:"title" gorm:"column:title"`                     // 标题
		Sort           int64      `json:"sort" gorm:"column:sort"`                       // 顺序
		SkipUrl        string     `json:"skip_url" gorm:"column:skip_url"`               // 跳转URL
		ImgHref        string     `json:"img_href" gorm:"column:img_href"`               // 图片地址
		Status         uint64     `json:"status" gorm:"column:status"`                   // 状态：1-上架，2-下架
		NeedLogin      uint64     `json:"need_login" gorm:"column:need_login"`           // 是否需要登录：1-不需要，2-需要
		SkipTarget     uint64     `json:"skip_target" gorm:"column:skip_target"`         // 跳转目标：1-H5,2-客服
		StartAt        *time.Time `json:"start_at" gorm:"column:start_at"`               // 开始时间
		ExpirationTime *time.Time `json:"expiration_time" gorm:"column:expiration_time"` // 过期时间
		CreatedTime    time.Time  `json:"created_time" gorm:"column:created_time"`       // 创建时间
		UpdatedTime    time.Time  `json:"updated_time" gorm:"column:updated_time"`       // 更新时间

	}
)

func newDappBannerModel(conn *gorm.DB) *defaultDappBannerModel {
	return &defaultDappBannerModel{
		conn:  conn,
		table: "`dapp_banner`",
	}
}

func (m *defaultDappBannerModel) Insert(ctx context.Context, data *DappBanner) error {
	err := m.conn.WithContext(ctx).Create(&data).Error
	return err
}

func (m *defaultDappBannerModel) FindOne(ctx context.Context, id int64) (*DappBanner, error) {
	var resp DappBanner
	err := m.conn.WithContext(ctx).Model(&DappBanner{}).Where("`id` = ?", id).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case gorm.ErrRecordNotFound:
		return nil, gorm.ErrRecordNotFound
	default:
		return nil, err
	}
}

func (m *defaultDappBannerModel) Update(ctx context.Context, session *gorm.DB, data *DappBanner) error {
	err := m.conn.WithContext(ctx).Updates(data).Error
	return err
}

func (m *defaultDappBannerModel) Delete(ctx context.Context, session *gorm.DB, id int64) error {
	err := m.conn.WithContext(ctx).Delete(&DappBanner{}, id).Error

	return err
}

func (m *defaultDappBannerModel) tableName() string {
	return m.table
}
