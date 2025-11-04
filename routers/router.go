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
	////////////////////////////////////////dapp发现页面接口//////////////////////////////////////////////////////////////
	dapp := api.Group("/dapp")
	dapp.POST("/manage/banner/list", controllers.GetManageBannerList)
	dapp.POST("/manage/banner/operation", controllers.ManageBannerOperation)
	dapp.POST("/banner/list", controllers.GetBannerList)

	dapp.POST("/manage/discover/tool/categories/list", controllers.GetManageDiscoverToolCategoriesList)
	dapp.POST("/manage/discover/tool/categories/operation", controllers.ManageDiscoverToolCategoriesOperation)
	dapp.POST("/discover/tool/categories/list", controllers.GetDiscoverToolCategoriesList)

	dapp.POST("/manage/network/list", controllers.GetManageNetworkList)
	dapp.POST("/manage/network/operation", controllers.ManageNetworkOperation)
	dapp.POST("/network/list", controllers.GetNetworkList)

	dapp.POST("/manage/discover/tool/info/list", controllers.GetManageDiscoverToolInfoList)
	dapp.POST("/manage/discover/tool/info/operation", controllers.ManageDiscoverToolInfoOperation)
	dapp.POST("/discover/tool/info/list", controllers.GetDiscoverToolInfoList)

	dapp.POST("/manage/discover/toolbar/list", controllers.GetManageDiscovertoolbarList)
	dapp.POST("/manage/discover/toolbar/operation", controllers.ManageDiscovertoolbarOperation)
	dapp.POST("/discover/toolbar/list", controllers.GetDiscoverToolbarList)

	dapp.POST("/discover/tool/favorites/list", controllers.GetDiscoverToolFavoritesList)
	dapp.POST("/discover/tool/favorites/operation", controllers.DiscoverToolFavoritesOperation)

	dapp.POST("/discover/tool/event/list", controllers.GetDiscoverToolEventList)
	dapp.POST("/discover/tool/event/operation", controllers.DiscoverToolEventOperation)
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////前端日志上报//////////////////////////////////////////////////////////////
	api.POST("/client/log/reporting", controllers.ClientLogReporting)
	api.POST("/client/log", controllers.GetClientLogReporting)
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	//////////////////////////////////////im开放接口////////////////////////////////////////////////////////////////
	api.POST("/authenticate", controllers.Authenticate)

	//IM检查服务状态
	api.GET("/api/circle/version", controllers.Version)
	//统计回调方法
	api.POST("/receivePayResult", controllers.ReceivePayResult)
	//IM

	// 支付相关-订单退款
	api.POST("/order/refund", controllers.OrderRefund)
	// 支付相关-订单提现
	api.POST("/order/withdrawal", controllers.OrderWithdrawal)
	// 支付相关-创建订单
	api.POST("/create/order", controllers.CreateOrder)
	// 支付相关-查询订单
	api.POST("/order/info", controllers.GetOrderInfo)
	// 支付相关-查询openid
	api.POST("/openid/user/info", controllers.GetOpenIdUserInfo)

	return router

}
