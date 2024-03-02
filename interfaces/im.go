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
)

// 支付相关-订单退款
func RefundOrder(orderId, amount string, isModel bool) error {
	url := configs.Conf.TestImUrl + "/wallet/refund_order"
	if isModel {
		url = configs.Conf.ImUrl + "/wallet/refund_order"
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
func PayOutCurrencyForOfficial(userCard, currencyid, amount string, isModel bool) error {
	url := configs.Conf.TestImUrl + "/wallet/pay_out_currency_for_official"
	if isModel {
		url = configs.Conf.ImUrl + "/wallet/pay_out_currency_for_official"
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
func CreateOrder(title string, coinType int64, amount string, nonce string, isModel bool) (*protos.ImCommonResp, string, error) {
	url := configs.Conf.TestImUrl + "/wallet/create_order"
	if isModel {
		url = configs.Conf.ImUrl + "/wallet/create_order"
	}
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

// 支付相关-查询订单
func QueryOrder(orderId string, isModel bool) (*protos.ImCommonResp, error) {
	url := configs.Conf.TestImUrl + "/wallet/query_order"
	if isModel {
		url = configs.Conf.ImUrl + "/wallet/query_order"
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
func GetOpenId(account string, isModel bool) (*protos.ImCommonResp, error) {
	url := configs.Conf.TestImUrl + "/official/get_openid"
	if isModel {
		url = configs.Conf.ImUrl + "/official/get_openid"
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
