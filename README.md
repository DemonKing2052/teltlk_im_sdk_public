# Teltlk-Im-Sdk

#### 介绍
{**以下是 Teltlk-Im-Sdk 后端说明**
后端演示对接Teltlk SDK的使用方法，目前所支持对接的接口协议有，创建订单，查询订单， 订单退款， 提现
}

#### 软件架构

golang gin


#### 安装教程

1.  go mod tidy  获取依赖
2.  go run  运行
3.  使用goland软件对go源码进行开发，编译

#### 接口使用例子

1.  创建订单
--定义结构
type (
	CreateOrderReq struct {
		Amount     string `form:"amount" json:"amount" binding:"required"`           //订单金额
		CurrencyId int64  `form:"currency_id" json:"currency_id" binding:"required"` //币种
		Title      string `form:"title" json:"title" binding:"required"`             //标识
	}

	CreateOrderResp struct {
		OrderId        string `json:"order_id"`          //订单id
		PayCallBackUrl string `json:"pay_call_back_url"` //回调
	}
)
--请求方法
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
--sdk接口编写，实现调用
// 支付相关-创建订单
func CreateOrder(title string, coinType int64, amount string, nonce string) (*protos.ImCommonResp, string, error) {
	url := configs.Conf.ImUrl + "/wallet/create_order"
	callBackUrl := configs.Conf.CallBackUrl + "/ReceivePayResult"
	// 构建请求参数
	params := map[string]any{
		"title":      title,
		"currencyID": coinType,
		"amount":     amount,
		"desc":       title,
		"outOrderID": nonce,
		"appID":      svc.Ctx.Config.Token.Appid,
		"nonce":      nonce,
		"sign":       utils.Md5(svc.Ctx.Config.Token.Secret + nonce),
	}
	// 发起POST请求
	body, err := utils.PostJsonData(url, params)
	if err != nil {
		fmt.Println("Post 创建订单， %s ", err.Error())
		return nil, "", err
	}
	// 解析响应JSON
	fmt.Println("参数为：" + string(body))
	res := protos.ImCommonResp{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("Post 创建订单 解析响应JSON， %s ", err.Error())
		return nil, "", err
	}
	// 处理响应结果
	if res.Data != nil {
		return &res, callBackUrl, nil
	}
	return &res, "", errors.New(fmt.Sprintf("创建订单号未知错误 ：%s", string(body)))
}

