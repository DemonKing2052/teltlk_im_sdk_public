package e

var MsgFlags = map[int64]string{
	SUCCESS: "success",
	ERROR:   "fail",
	//
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
