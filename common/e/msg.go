package e

var MsgFlags = map[int64]string{
	SUCCESS: "success",
	ERROR:   "fail",
	//
	ErrorInvalidParam:  "请求参数错误",
	ErrorTokenNotExist: "请登录后再操作",
}

type MyError struct {
	Code    string
	Message string
}

// GetMsg get warn information based on Code
func GetMsg(code int64) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
