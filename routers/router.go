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
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	api.POST("/client/log/reporting", controllers.ClientLogReporting)
	api.POST("/client/log", controllers.GetClientLogReporting)
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	//////////////////////////////////////////////////////////////////////////////////////////////////////
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
