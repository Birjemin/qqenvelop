package utils

import (
	"time"
)

// GetCurrTime get expire timestamp
func GetCurrTime() int {
	return int(time.Now().Unix())
}

// GetDateNum get date num
func GetDateNum() string {
	loc, _ := time.LoadLocation("Asia/Chongqing")
	return time.Now().In(loc).Format("20060102")
}
