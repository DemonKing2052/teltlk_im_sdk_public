package routers

import (
	"ImSdk/controllers"
	"ImSdk/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	//router.Use(middleware.GinCors())
	api := router.Group("/api/v1/public")
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////前端日志上报//////////////////////////////////////////////////////////////
	clientLog := api.Group("/client")
	clientLog.Use(middleware.GinCors())
	clientLog.Any("/log", controllers.GetClientLogReporting)
	clientLog.Any("/log/reporting", controllers.ClientLogReporting)
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	//////////////////////////////////////im开放接口////////////////////////////////////////////////////////////////
	im := api.Group("")
	im.Use(middleware.GinCors())
	im.Any("/authenticate", controllers.Authenticate)
	//IM检查服务状态
	im.Any("/api/circle/version", controllers.Version)
	//统计回调方法
	im.Any("/receivePayResult", controllers.ReceivePayResult)
	//IMAny
	// 支付相关-订单退款
	im.Any("/order/refund", controllers.OrderRefund)
	// 支付相关-订单提现
	im.Any("/order/withdrawal", controllers.OrderWithdrawal)
	// 支付相关-创建订单
	im.Any("/create/order", controllers.CreateOrder)
	// 支付相关-查询订单
	im.Any("/order/info", controllers.GetOrderInfo)
	// 支付相关-查询openid
	im.Any("/openid/user/info", controllers.GetOpenIdUserInfo)

	return router

}
