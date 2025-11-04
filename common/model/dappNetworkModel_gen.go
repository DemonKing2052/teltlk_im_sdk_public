package model

import (
	"ImSdk/common/utils"
	"context"
	"time"

	"gorm.io/gorm"
)

var ()

type (
	dappNetworkModel interface {
		Insert(ctx context.Context, data *DappNetwork) error
		FindOne(ctx context.Context, id int64) (*DappNetwork, error)
		Update(ctx context.Context, session *gorm.DB, data *DappNetwork) error
		Delete(ctx context.Context, session *gorm.DB, id int64) error
		NetWorkFindListPage(ctx context.Context, page *utils.PageData) (*[]DappNetwork, int64, error)
	}

	defaultDappNetworkModel struct {
		conn  *gorm.DB
		table string
	}

	DappNetwork struct {
		Id          int64     `json:"id" gorm:"column:id"`
		Name        string    `json:"name" gorm:"column:name"`                 // 网络名称，如 Ethereum、BSC、Solana
		ChainId     string    `json:"chain_id" gorm:"column:chain_id"`         // 链ID，如 1、56、solana
		Symbol      string    `json:"symbol" gorm:"column:symbol"`             // 链主币符号，例如 ETH、BNB、SOL
		RpcUrl      string    `json:"rpc_url" gorm:"column:rpc_url"`           // 默认 RPC，非必填
		ExplorerUrl string    `json:"explorer_url" gorm:"column:explorer_url"` // 浏览器地址，例如 https://bscscan.com
		Logo        string    `json:"logo" gorm:"column:logo"`                 // 图标
		Status      int64     `json:"status" gorm:"column:status"`             // 状态 1启用 2禁用
		Sort        int64     `json:"sort" gorm:"column:sort"`                 // 排序
		CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	}
)

func newDappNetworkModel(conn *gorm.DB) *defaultDappNetworkModel {
	return &defaultDappNetworkModel{
		conn:  conn,
		table: "`dapp_network`",
	}
}

func (m *defaultDappNetworkModel) Insert(ctx context.Context, data *DappNetwork) error {
	err := m.conn.WithContext(ctx).Create(&data).Error
	return err
}

func (m *defaultDappNetworkModel) FindOne(ctx context.Context, id int64) (*DappNetwork, error) {
	var resp DappNetwork
	err := m.conn.WithContext(ctx).Model(&DappNetwork{}).Where("`id` = ?", id).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case gorm.ErrRecordNotFound:
		return nil, gorm.ErrRecordNotFound
	default:
		return nil, err
	}
}

func (m *defaultDappNetworkModel) Update(ctx context.Context, session *gorm.DB, data *DappNetwork) error {
	err := m.conn.WithContext(ctx).Updates(data).Error
	return err
}

func (m *defaultDappNetworkModel) Delete(ctx context.Context, session *gorm.DB, id int64) error {
	err := m.conn.WithContext(ctx).Delete(&DappNetwork{}, id).Error

	return err
}

func (m *defaultDappNetworkModel) tableName() string {
	return m.table
}
