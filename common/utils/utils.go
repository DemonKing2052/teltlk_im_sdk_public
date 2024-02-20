package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Md5(s string, salt ...string) string {
	h := md5.New()
	h.Write([]byte(s))
	if len(salt) > 0 {
		h.Write([]byte(salt[0]))
	}
	cipher := h.Sum(nil)
	return hex.EncodeToString(cipher)
}

func GenerateRandomOrderID() string {
	// 在这里生成随机的订单号，可以使用时间戳、随机数等方式
	// 例如，使用时间戳作为订单号
	orderID := strconv.FormatInt(time.Now().Unix(), 10)
	return orderID
}

func GenerateRandomStr() string {
	// 生成一个随机字符串作为文件名，可以根据需要自定义长度和字符集
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 20
	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = charset[rand.Intn(len(charset))]
	}
	return string(randomString)
}

func PostJsonData(url string, params interface{}) ([]byte, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	request, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	// 处理响应阶段：
	defer request.Body.Close()

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func StrToFloat64(value string) float64 {
	i, err := strconv.ParseFloat(value, 2)
	if err != nil {
		return 0
	}
	return i
}

func FloatToStr(value float64) string {
	s := strconv.FormatFloat(value, 'f', -1, 64)
	return s
}

func StrToInt(s string) int64 {
	v, _ := strconv.ParseInt(s, 10, 64)
	return v
}

func IntToStr(i int64) string {
	s := strconv.FormatInt(i, 10)
	return s
}

func AnyToStr(value any) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

// 转换并保留7位小数
func StrTo7Float(s string) float64 {
	newn := strings.Split(s, ".")
	l := len(newn)
	if l > 2 {
		return 0
	} else if l == 2 {
		if len(newn[1]) > 7 {
			v, _ := strconv.ParseFloat(newn[0]+"."+newn[1][:7], 64)
			return v
		} else {
			v, _ := strconv.ParseFloat(s, 64)
			return v
		}
	}
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

func FloatTo7Str(i float64) string {
	n := strconv.FormatFloat(i, 'f', -1, 64)
	if n == "" {
		return ""
	}
	newn := strings.Split(n, ".")
	if len(newn) < 2 {
		return n
	} else {
		if len(newn[1]) > 7 {
			return newn[0] + "." + newn[1][:7]
		}
		return n
	}
}
