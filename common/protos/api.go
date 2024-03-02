package protos

type ReceivePayResultResp struct {
	Actual     string `json:"actual"`
	Amount     string `json:"amount"`
	CurrencyID int    `json:"currencyID"`
	FeeAmount  string `json:"feeAmount"`
	Nonce      string `json:"nonce"`
	OpenID     string `json:"openID"`
	OrderID    string `json:"orderID"`
	OrderTitle string `json:"orderTitle"`
	OutOrderID string `json:"outOrderID"`
	Sign       string `json:"sign"`
	Status     int    `json:"status"`
}

type (
	QueryOrderReq struct {
		OrderId string `form:"order_id" json:"order_id" binding:"required"` //订单id
		IsModel bool   `form:"is_model" json:"is_model"`                    //调用是否是正式，测试
	}
)

type (
	OrderRefundReq struct {
		OrderId string `form:"order_id" json:"order_id" binding:"required"` //订单id
		Amount  string `form:"amount" json:"amount" binding:"required"`     //订单金额
		IsModel bool   `form:"is_model" json:"is_model"`                    //调用是否是正式，测试
	}
)

type (
	OrderWithdrawalReq struct {
		OpenId     string `form:"open_id" json:"open_id" binding:"required"`         //openId
		Amount     string `form:"amount" json:"amount" binding:"required"`           //订单金额
		CurrencyId string `form:"currency_id" json:"currency_id" binding:"required"` //币种
		IsModel    bool   `form:"is_model" json:"is_model"`                          //调用是否是正式，测试

	}
)

type (
	CreateOrderReq struct {
		Amount     string `form:"amount" json:"amount" binding:"required"`           //订单金额
		CurrencyId int64  `form:"currency_id" json:"currency_id" binding:"required"` //币种
		Title      string `form:"title" json:"title" binding:"required"`             //标识
		IsModel    bool   `form:"is_model" json:"is_model"`                          //调用是否是正式，测试
	}

	CreateOrderResp struct {
		OrderId        string `json:"order_id"`          //订单id
		PayCallBackUrl string `json:"pay_call_back_url"` //回调
	}
)

type (
	GetOpenIdUserInfoReq struct {
		TeltlkId string `form:"teltlk_id" json:"teltlk_id" binding:"required"` //im id
		IsModel  bool   `form:"is_model" json:"is_model"`                      //调用是否是正式，测试
	}
)
