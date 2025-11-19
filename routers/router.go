package routers

import (
	"ImSdk/controllers"
	"ImSdk/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	router.Use(middleware.GinCors())
	api := router.Group("/api/v1/public")
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////dapp发现页面管理接口//////////////////////////////////////////////////////////////
	dappAdmin := api.Group("/manage/dapp")
	dappAdmin.POST("/banner/list", controllers.GetManageBannerList)
	dappAdmin.POST("/banner/operation", controllers.ManageBannerOperation)

	dappAdmin.POST("/discover/tool/categories/list", controllers.GetManageDiscoverToolCategoriesList)
	dappAdmin.POST("/discover/tool/categories/operation", controllers.ManageDiscoverToolCategoriesOperation)

	dappAdmin.POST("/network/list", controllers.GetManageNetworkList)
	dappAdmin.POST("/network/operation", controllers.ManageNetworkOperation)

	dappAdmin.POST("/discover/tool/info/list", controllers.GetManageDiscoverToolInfoList)
	dappAdmin.POST("/discover/tool/info/operation", controllers.ManageDiscoverToolInfoOperation)

	dappAdmin.POST("/discover/toolbar/list", controllers.GetManageDiscovertoolbarList)
	dappAdmin.POST("/discover/toolbar/operation", controllers.ManageDiscovertoolbarOperation)
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////dapp发现页面管理接口//////////////////////////////////////////////////////////////
	dapp := api.Group("/dapp")
	//dapp.Use(middleware.GinCors())
	dapp.POST("/banner/list", controllers.GetBannerList)

	dapp.POST("/discover/tool/categories/list", controllers.GetDiscoverToolCategoriesList)

	dapp.POST("/network/list", controllers.GetNetworkList)

	dapp.POST("/discover/tool/info/list", controllers.GetDiscoverToolInfoList)

	dapp.POST("/discover/toolbar/list", controllers.GetDiscoverToolbarList)

	dapp.POST("/discover/tool/favorites/list", controllers.GetDiscoverToolFavoritesList)
	dapp.POST("/discover/tool/favorites/operation", controllers.DiscoverToolFavoritesOperation)

	dapp.POST("/discover/tool/event/list", controllers.GetDiscoverToolEventList)
	dapp.POST("/discover/tool/event/operation", controllers.DiscoverToolEventOperation)
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////前端日志上报//////////////////////////////////////////////////////////////
	clientLog := router.Group("/api/v1/public/client")
	//clientLog.Use(middleware.GinCors())
	clientLog.POST("/log", controllers.GetClientLogReporting)
	clientLog.POST("/log/reporting", controllers.ClientLogReporting)
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	//////////////////////////////////////im开放接口////////////////////////////////////////////////////////////////
	im := router.Group("/api/v1/public")
	//im.Use(middleware.GinCors())
	im.POST("/authenticate", controllers.Authenticate)
	//IM检查服务状态
	im.GET("/api/circle/version", controllers.Version)
	//统计回调方法
	im.POST("/receivePayResult", controllers.ReceivePayResult)
	//IM
	// 支付相关-订单退款
	im.POST("/order/refund", controllers.OrderRefund)
	// 支付相关-订单提现
	im.POST("/order/withdrawal", controllers.OrderWithdrawal)
	// 支付相关-创建订单
	im.POST("/create/order", controllers.CreateOrder)
	// 支付相关-查询订单
	im.POST("/order/info", controllers.GetOrderInfo)
	// 支付相关-查询openid
	im.POST("/openid/user/info", controllers.GetOpenIdUserInfo)

	return router

}
