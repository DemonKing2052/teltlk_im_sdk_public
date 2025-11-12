package utils

import (
	"strings"
	"time"
)

const (
	timeFormat      = "2006-01-02 15:04:05"
	TimeDayFormat   = "2006-01-02"
	TimestampFormat = "20060102150405.000"
	TimeFormat      = "15:04:05"
)

func GetTimeStampStr() string {
	s := time.Now().Format(TimestampFormat)
	res := strings.Split(s, ".")
	return res[0] + res[1]
}

func GetCurruntTimerBefore(day int) string {
	return time.Now().AddDate(0, 0, day).Format(TimeDayFormat)
}

// StrtoTime 日期字符轉時間戳
func StrToTime(s string) (time.Time, error) {
	res, err := time.ParseInLocation(timeFormat, s, time.Local)
	return res, err
}

// TimeToString 格式化日期字符
func TimeToString(t time.Time) string {
	return t.Format(timeFormat)
}

func TimeToStringPtr(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(timeFormat)
}

// 秒
func Timestamp() int64 {
	return time.Now().Unix()
}
func GetDaysBetween(t1 *time.Time, t2 *time.Time) float64 {
	if t1 == nil || t2 == nil {
		return 0
	}
	//计算两个时间过去了多少天
	return t2.Sub(*t1).Hours() / 24
}

// 时间戳到时间秒级
func StampToDateTimeS(s int64) time.Time {
	t := time.Unix(s, 0)
	return t
}

// 时间戳到时间毫秒级
func StampToDateTimeN(s int64) time.Time {
	t := time.Unix(0, s*int64(time.Millisecond))
	return t
}

func TimeStringToTime(timeString string) (time.Time, error) {
	t, err := time.ParseInLocation(TimeDayFormat, timeString, time.Local)
	return t, err
}
