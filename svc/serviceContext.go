package svc

import (
	"ImSdk/common/model"
	"ImSdk/configs"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var Ctx *ServiceContext

type ServiceContext struct {
	Config       configs.Config
	DbClient     *gorm.DB
	ClientLogReq struct {
		FhMutex sync.Mutex
		Fh      *os.File
	}
	//
	DappBannerModel                 model.DappBannerModel
	DappNetworkModel                model.DappNetworkModel
	DappDiscoverToolbarModel        model.DappDiscoverToolbarModel
	DappDiscoverToolCategoriesModel model.DappDiscoverToolCategoriesModel
	DappDiscoverToolInfoModel       model.DappDiscoverToolInfoModel
	DappDiscoverToolFavoritesModel  model.DappDiscoverToolFavoritesModel
	DappDiscoverToolEventsModel     model.DappDiscoverToolEventsModel
}

func NewServiceContext(c configs.Config) *ServiceContext {
	db := InitDbClient(c)
	return &ServiceContext{
		Config:   c,
		DbClient: db,
		ClientLogReq: struct {
			FhMutex sync.Mutex
			Fh      *os.File
		}{
			Fh: ensureLogFile(c.ClientLogReq.Path),
		},
		//
		DappBannerModel:                 model.NewDappBannerModel(db),
		DappNetworkModel:                model.NewDappNetworkModel(db),
		DappDiscoverToolbarModel:        model.NewDappDiscoverToolbarModel(db),
		DappDiscoverToolCategoriesModel: model.NewDappDiscoverToolCategoriesModel(db),
		DappDiscoverToolInfoModel:       model.NewDappDiscoverToolInfoModel(db),
		DappDiscoverToolFavoritesModel:  model.NewDappDiscoverToolFavoritesModel(db),
		DappDiscoverToolEventsModel:     model.NewDappDiscoverToolEventsModel(db),
	}
}

func ensureLogFile(p string) *os.File {
	dir := filepath.Dir(p)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil
	}
	f, err := os.OpenFile(p, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil
	}
	return f
}

func InitDbClient(c configs.Config) *gorm.DB {
	fmt.Printf("init mysql start...")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=30s",
		c.MySQL.User, c.MySQL.Password, c.MySQL.Host, c.MySQL.Port, c.MySQL.Database)
	dB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(Writer{}, logger.Config{
			SlowThreshold:             time.Duration(c.MySQL.SlowThreshold) * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.LogLevel(c.MySQL.LogLevel),                       // Log level
			IgnoreRecordNotFoundError: true,                                                    // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                                                    // Disable color
		},
		),
	})
	if err != nil {
		panic(fmt.Sprintf("init mysql failed,err:%v", err))
	}

	sqlDB, err := dB.DB()
	if err != nil {
		panic(fmt.Sprintf("init gorm failed,err:%v", err))
	}
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(c.MySQL.DBMaxLifeTime))
	sqlDB.SetMaxOpenConns(c.MySQL.DBMaxOpenConns)
	sqlDB.SetMaxIdleConns(c.MySQL.DBMaxIdleConns)

	dB.Set("gorm:table_options", "CHARSET=utf8mb4")
	dB.Set("gorm:table_options", "collation=utf8mb4_unicode_ci")
	fmt.Printf("init mysql ok")
	return dB
}

type Writer struct{}

func (w Writer) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)

}
