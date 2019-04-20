package utils

import "fmt"

// SliceJoin 合并字符串
func SliceJoin(s []interface{}, sep string) string {
	str := ""
	for _, item := range s {
		str = fmt.Sprintf("%s%s%v", str, sep, item)
	}
	str = str[1:]
	return str
}
