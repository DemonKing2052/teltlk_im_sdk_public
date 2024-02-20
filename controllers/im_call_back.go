package controllers

import (
	"ImSdk/common/e"
	"ImSdk/common/protos"
	"ImSdk/svc"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

/* IM 定义的状态
ORDER_STATUS_DELETED  = -2
ORDER_STATUS_CLOSED   = -1
ORDER_STATUS_CREATED  = 0
ORDER_STATUS_PAYED    = 1
ORDER_STATUS_REFUNDED = 2
*/
// 接收支付结果回调参数
func ReceivePayResult(c *gin.Context) {

	fmt.Println("返回结果！")
	// 读取请求内容
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("返回结果err :", err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": err.Error()})
		return
	}
	// 解析JSON
	fmt.Println("参数为：" + string(body))
	var data protos.ReceivePayResultResp
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("解释结果错误！", err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": "JSON unmarshal error"})
		return
	}
	//验证签名是否正确 查询订单id是否存在 计算 MD5 哈希值
	hash := md5.Sum([]byte(svc.Ctx.Config.Token.Secret + data.Nonce))
	md5Hash := hex.EncodeToString(hash[:])
	fmt.Println(md5Hash)
	if md5Hash != data.Sign {
		c.JSON(500, gin.H{"code": 500, "message": errors.New("Sign error").Error(), "result": ""})
		return
	}

	c.String(http.StatusOK, e.GetMsg(e.SUCCESS))
}
