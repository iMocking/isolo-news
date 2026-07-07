package schema

import "time"

// nowUnix 返回当前 UTC Unix 时间戳
func nowUnix() int64 {
	return time.Now().Unix()
}

// nowUnixDate 返回当前 UTC 日期的时间戳（当天 00:00:00 UTC）
func nowUnixDate() int64 {
	now := time.Now().UTC()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).Unix()
}
