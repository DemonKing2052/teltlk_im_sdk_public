package common

var (
	JwtKey = []byte("PIProject")
)

const (

	//支付状态
	Payment_Status_Not     = 0 //0未支付1已支付2支付失败3已退款
	Payment_Status_Success = 1 //0未支付1已支付2支付失败3已退款
	Payment_Status_Fail    = 2 //0未支付1已支付2支付失败3已退款
	Payment_Status_Refund  = 3 //0未支付1已支付2支付失败3已退款

	Create_Witdrawal = "提现"
)
