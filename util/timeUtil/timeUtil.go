package timeUtil

import "time"

const TIME_FORMAT_YMDHMS = "2006-01-02 15:04:05"
const TIME_FORMAT_YMD = "2006-01-02"
const TIME_FORMAT_SHORT_YMDHMS = "01/02/2006 15:04:05"

func GetNowTimeStamp() int64 {
	return time.Now().Unix()
}

func GetTimeStamp(timeStr, formatStr string) int64 {
	t, err := time.Parse(formatStr, timeStr)
	if err != nil {
		return 0
	}

	return t.Unix()
}

func GetStringTime(timeStamp int64, formatStr string) string {
	return time.Unix(timeStamp, 0).Format(formatStr)
}

func GetNowStringTime(formatStr string) string {
	return time.Now().Format(formatStr)
}