package utils

import "time"

// Now 获取1970到现在的秒数
func Now() int {
	local, err := time.LoadLocation("Local") //服务器设置的时区
	if err != nil {
		return int(time.Now().Unix())
	}
	now := time.Now().In(local)
	return int(now.Unix())
}
