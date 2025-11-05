package interfaces

import (
	"ImSdk/common"
	"ImSdk/common/protos"
	"ImSdk/common/utils"
	"ImSdk/configs"
	"ImSdk/svc"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// 支付相关-订单退款
func RefundOrder(orderId, amount string, model string) error {
	url := configs.Conf.Teltlk.TestImUrl + "/wallet/refund_order"
	if model == "model" {
		url = configs.Conf.Teltlk.ImUrl + "/wallet/refund_order"
	}
	// 生成随机订单号
	nonce := utils.GenerateRandomOrderID()
	// 构建请求参数
	params := map[string]any{
		"orderID": orderId,
		"amount":  amount,
		"appID":   svc.Ctx.Config.Token.Appid,
		"nonce":   nonce,
		"sign":    utils.Md5(svc.Ctx.Config.Token.Secret + nonce),
	}
	// 发起POST请求
	body, err := utils.PostJsonData(url, params)
	if err != nil {
		fmt.Println("Post 订单退款， %s ", err.Error())
		return err
	}
	// 解析响应JSON
	fmt.Println("参数为：" + string(body))
	res := protos.ImCommonResp{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("Post 订单退款 解析响应JSON， %s ", err.Error())
		return err
	}
	// 处理响应结果
	if res.Data != nil {
		return nil
	}
	return errors.New(fmt.Sprintf("%s", string(body)))
}

// 支付相关-订单提现
func PayOutCurrencyForOfficial(userCard, currencyid, amount string, model string) error {
	url := configs.Conf.Teltlk.TestImUrl + "/wallet/pay_out_currency_for_official"
	if model == "model" {
		url = configs.Conf.Teltlk.ImUrl + "/wallet/pay_out_currency_for_official"
	}
	outOrderID := utils.GenerateRandomOrderID()
	// 构建请求参数
	params := map[string]any{
		"payPassword": "",
		"currencyID":  utils.StrToInt(currencyid),
		"amount":      utils.AnyToStr(amount),
		"toOpenID":    userCard,
		"reference":   outOrderID,
		"title":       common.Create_Witdrawal,
		"appID":       svc.Ctx.Config.Token.Appid,
		"nonce":       outOrderID,
		"sign":        utils.Md5(svc.Ctx.Config.Token.Secret + outOrderID),
	}
	// 发起POST请求
	body, err := utils.PostJsonData(url, params)
	if err != nil {
		fmt.Println("Post 订单提现， %s ", err.Error())
		return err
	}
	// 解析响应JSON
	fmt.Println("参数为：" + string(body))
	res := protos.ImCommonResp{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("Post 订单提现 解析响应JSON， %s ", err.Error())
		return err
	}
	// 处理响应结果
	if res.Data != nil {
		dataValue, ok := res.Data.(map[string]interface{})
		if !ok {
			return errors.New(fmt.Sprintf("%+v", body))
		}
		fmt.Println(dataValue)
		if utils.StrToInt(utils.AnyToStr(dataValue["result"])) == common.Payment_Status_Success {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("%s", string(body)))
}

// 支付相关-创建订单
func CreateOrder(title string, coinType int64, amount string, nonce string, model string) (*protos.ImCommonResp, string, error) {
	url := configs.Conf.Teltlk.TestImUrl + "/wallet/create_order"
	if model == "model" {
		url = configs.Conf.Teltlk.ImUrl + "/wallet/create_order"
	}
	callBackUrl := configs.Conf.Teltlk.CallBackUrl + "/ReceivePayResult"
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

// 支付相关-查询订单
func QueryOrder(orderId string, model string) (*protos.ImCommonResp, error) {
	url := configs.Conf.Teltlk.TestImUrl + "/wallet/query_order"
	if model == "model" {
		url = configs.Conf.Teltlk.ImUrl + "/wallet/query_order"
	}
	nonce := utils.GenerateRandomOrderID()
	// 构建请求参数
	params := map[string]interface{}{
		"orderID": orderId,
		"appID":   svc.Ctx.Config.Token.Appid,
		"nonce":   nonce,
		"sign":    utils.Md5(svc.Ctx.Config.Token.Secret + nonce),
	}
	// 发起POST请求
	body, err := utils.PostJsonData(url, params)
	if err != nil {
		fmt.Println("Post 查询订单， %s ", err.Error())
		return nil, err
	}
	// 解析响应JSON
	fmt.Println("参数为：" + string(body))
	res := protos.ImCommonResp{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("Post 查询订单 解析响应JSON， %s ", err.Error())
		return nil, err
	}
	// 处理响应结果
	if res.Data != nil {
		return &res, nil
	}
	return &res, errors.New(fmt.Sprintf("查询订单号未知错误 ：%s", string(body)))
}

// 支付相关-查询openid
func GetOpenId(account string, model string) (*protos.ImCommonResp, error) {
	url := configs.Conf.Teltlk.TestImUrl + "/official/get_openid"
	if model == "model" {
		url = configs.Conf.Teltlk.ImUrl + "/official/get_openid"
	}
	nonce := utils.GenerateRandomOrderID()
	// 构建请求参数
	params := map[string]interface{}{
		"account": account,
		"appID":   svc.Ctx.Config.Token.Appid,
		"nonce":   nonce,
		"sign":    utils.Md5(svc.Ctx.Config.Token.Secret + nonce),
	}
	// 发起POST请求
	body, err := utils.PostJsonData(url, params)
	if err != nil {
		fmt.Println("Post 查询用户信息， %s ", err.Error())
		return nil, err
	}
	// 解析响应JSON
	fmt.Println("参数为：" + string(body))
	res := protos.ImCommonResp{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("Post 查询用户信息 解析响应JSON， %s ", err.Error())
		return nil, err
	}
	// 处理响应结果
	if res.Data != nil {
		dataValue, ok := res.Data.(map[string]interface{})
		if !ok {
			return nil, errors.New(fmt.Sprintf("%+v", body))
		}
		fmt.Println(dataValue)
		if utils.StrToInt(utils.AnyToStr(dataValue["result"])) == common.Payment_Status_Success {

			return &res, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("查询用户信息未知错误 ：%s", string(body)))
}

type GetSelfUserInfoReq struct {
	OperationID string `json:"operationID" binding:"required"`
	UserID      string `json:"userID"`
}
type CommResp struct {
	ErrCode int32  `json:"errCode" example:"0" description:"状态码，0 表示成功"`
	ErrMsg  string `json:"errMsg" example:"ok" description:"响应消息"`
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID           string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Nickname         string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	FaceURL          string `protobuf:"bytes,3,opt,name=faceURL,proto3" json:"faceURL,omitempty"`
	Gender           int32  `protobuf:"varint,4,opt,name=gender,proto3" json:"gender,omitempty"`
	PhoneNumber      string `protobuf:"bytes,5,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Birth            string `protobuf:"bytes,6,opt,name=birth,proto3" json:"birth,omitempty"`
	Email            string `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	Uname            string `protobuf:"bytes,8,opt,name=uname,proto3" json:"uname,omitempty"`            //add by eric
	Utype            int32  `protobuf:"varint,9,opt,name=utype,proto3" json:"utype,omitempty"`           //add by eric
	Voiceid          string `protobuf:"bytes,10,opt,name=voiceid,proto3" json:"voiceid,omitempty"`       //add by eric
	Countryid        string `protobuf:"bytes,11,opt,name=countryid,proto3" json:"countryid,omitempty"`   //add by eric
	Piid             string `protobuf:"bytes,12,opt,name=piid,proto3" json:"piid,omitempty"`             //add by eric
	Wechatid         string `protobuf:"bytes,13,opt,name=wechatid,proto3" json:"wechatid,omitempty"`     //add by eric
	Twitterid        string `protobuf:"bytes,14,opt,name=twitterid,proto3" json:"twitterid,omitempty"`   //add by eric
	Googleid         string `protobuf:"bytes,15,opt,name=googleid,proto3" json:"googleid,omitempty"`     //add by eric
	Facebookid       string `protobuf:"bytes,16,opt,name=facebookid,proto3" json:"facebookid,omitempty"` //add by eric
	Whatsappid       string `protobuf:"bytes,17,opt,name=whatsappid,proto3" json:"whatsappid,omitempty"` //add by eric
	Appleid          string `protobuf:"bytes,18,opt,name=appleid,proto3" json:"appleid,omitempty"`
	Huaweiid         string `protobuf:"bytes,19,opt,name=huaweiid,proto3" json:"huaweiid,omitempty"`
	Qqid             string `protobuf:"bytes,20,opt,name=qqid,proto3" json:"qqid,omitempty"`
	Identity         int32  `protobuf:"varint,21,opt,name=identity,proto3" json:"identity,omitempty"` //add by eric
	Credit           int32  `protobuf:"varint,22,opt,name=credit,proto3" json:"credit,omitempty"`     //add by eric
	Status           int32  `protobuf:"varint,23,opt,name=status,proto3" json:"status,omitempty"`     //add by eric
	Ex               string `protobuf:"bytes,24,opt,name=ex,proto3" json:"ex,omitempty"`
	CreateTime       uint32 `protobuf:"varint,25,opt,name=createTime,proto3" json:"createTime,omitempty"`
	AppMangerLevel   int32  `protobuf:"varint,26,opt,name=appMangerLevel,proto3" json:"appMangerLevel,omitempty"`
	GlobalRecvMsgOpt int32  `protobuf:"varint,27,opt,name=globalRecvMsgOpt,proto3" json:"globalRecvMsgOpt,omitempty"`
}

type GetSelfUserInfoResp struct {
	CommResp
	UserInfo *UserInfo              `json:"-"`
	Data     map[string]interface{} `json:"data" swaggerignore:"true"`
}

func GetUserInfo(token string, model string) (*GetSelfUserInfoResp, error) {
	url := configs.Conf.Teltlk.TestImUrl + "/user/get_self_user_info"
	if model == "model" {
		url = configs.Conf.Teltlk.ImUrl + "/user/get_self_user_info"
	}
	nonce := utils.GenerateRandomOrderID()
	// 构建请求参数
	params := map[string]interface{}{
		"operationID": nonce,
	}
	// 发起POST请求
	body, err := utils.RequestsJSON(url, "", "", params, map[string]string{"token": token})
	if err != nil {
		fmt.Println("Post 查询用户信息， %s ", err.Error())
		return nil, err
	}
	// 解析响应JSON
	fmt.Println("参数为：" + string(body))
	res := GetSelfUserInfoResp{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("Post 查询用户信息 解析响应JSON， %s ", err.Error())
		return nil, err
	}
	// 处理响应结果
	return &res, nil
	//return nil, errors.New(fmt.Sprintf("查询用户信息未知错误 ：%s", string(body)))
}
