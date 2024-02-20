package routers

import (
	"ImSdk/controllers"
	"ImSdk/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	router.Use(middleware.GinCors())

	//////////////////////////////////////////////////////////////////////////////////////////////////////
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	router.POST("/authenticate", controllers.Authenticate)

	//IM检查服务状态
	router.GET("/api/circle/version", controllers.Version)
	//统计回调方法
	router.POST("/receivePayResult", controllers.ReceivePayResult)
	//IM

	// 支付相关-订单退款
	router.POST("/order/refund", controllers.OrderRefund)
	// 支付相关-订单提现
	router.POST("/order/withdrawal", controllers.OrderWithdrawal)
	// 支付相关-创建订单
	router.POST("/create/order", controllers.CreateOrder)
	// 支付相关-查询订单
	router.POST("/order/info", controllers.GetOrderInfo)
	// 支付相关-查询openid
	router.POST("/openid/user/info", controllers.GetOpenIdUserInfo)

	return router

}
