package controllers

import (
	"ImSdk/common/e"
	"ImSdk/common/protos"
	"ImSdk/common/utils"
	"ImSdk/interfaces"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // 导入MySQL驱动
	"net/http"
)

//////////////////////////////////////////////////////////////////////////////////////

func OrderRefund(c *gin.Context) {
	var req protos.OrderRefundReq
	if err := c.ShouldBind(&req); err != nil {
		fmt.Printf("获取请求参数：%s\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error(), "result": nil})
		return
	}

	err := interfaces.RefundOrder(req.OrderId, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": e.ERROR, "message": err.Error()})
		return
	}
	// 返回JSON响应给客户端
	c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": e.GetMsg(e.SUCCESS), "data": nil})
}

func OrderWithdrawal(c *gin.Context) {
	var req protos.OrderWithdrawalReq
	if err := c.ShouldBind(&req); err != nil {
		fmt.Printf("获取请求参数：%s\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error(), "result": nil})
		return
	}

	err := interfaces.PayOutCurrencyForOfficial(req.OpenId, req.CurrencyId, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": e.ERROR, "message": err.Error()})
		return
	}
	// 返回JSON响应给客户端
	c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": e.GetMsg(e.SUCCESS), "data": nil})
}

func CreateOrder(c *gin.Context) {

	var req protos.CreateOrderReq
	if err := c.ShouldBind(&req); err != nil {
		fmt.Printf("获取请求参数：%s\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "http.StatusBadRequest": err.Error(), "result": nil})
		return
	}
	// 生成随机订单号
	outOrderID := utils.GenerateRandomOrderID()
	res, callBackUrl, err := interfaces.CreateOrder(req.Title, req.CurrencyId, req.Amount, outOrderID)
	// 处理响应结果
	if res.Data != nil {
		dataValue, _ := res.Data.(map[string]interface{})
		orderId := utils.AnyToStr(dataValue["orderID"])
		fmt.Printf("订单号：%s\n", orderId)

		resp := protos.CreateOrderResp{
			OrderId:        orderId,
			PayCallBackUrl: callBackUrl,
		}
		// 返回JSON响应给客户端
		c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": e.GetMsg(e.SUCCESS), "data": resp})
		return
	}
	er := ""
	if err != nil {
		er = err.Error()
	}
	c.JSON(http.StatusInternalServerError, gin.H{"code": e.ERROR, "message": er})
	return
}

func GetOrderInfo(c *gin.Context) {
	var req protos.QueryOrderReq
	if err := c.ShouldBind(&req); err != nil {
		fmt.Printf("获取请求参数：%s\n", err.Error())
		c.JSON(400, gin.H{"code": 400, "message": err.Error(), "result": nil})
		return
	}
	res, err := interfaces.QueryOrder(req.OrderId)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"code": e.ERROR, "message": err.Error(), "data": res})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": e.GetMsg(e.SUCCESS), "data": res})
	return
}

func GetOpenIdUserInfo(c *gin.Context) {
	var req protos.GetOpenIdUserInfoReq
	if err := c.ShouldBind(&req); err != nil {
		fmt.Printf("获取请求参数：%s\n", err.Error())
		c.JSON(400, gin.H{"code": 400, "message": err.Error(), "result": nil})
		return
	}
	res, err := interfaces.GetOpenId(req.TeltlkId)
	if res.Data != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"code": e.ERROR, "message": err.Error(), "data": res})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "message": e.GetMsg(e.SUCCESS), "data": res})
	return
}
