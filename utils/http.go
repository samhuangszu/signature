package utils

import (
	"compress/gzip"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// CheckSign 检查签名
func CheckSign(request *http.Request, key string) error {
	signature := request.Header.Get("signature")
	sign, err := HTTPSign(request, key)
	if err != nil {
		return err
	}
	if signature == sign {
		return nil
	}
	return errors.New("签名不合法")
}

// HTTPSign 基于原生的http.Request生成签名
// query+form+postbody+timeStamp
func HTTPSign(request *http.Request, key string) (string, error) {
	timeStamp := request.Header.Get("TimeStamp")
	//检验时间
	delta := Now() - StrToInt(timeStamp)
	if delta < 0 || delta > 5 {
		return "", errors.New("时间错误")
	}
	method := request.Method
	if method == "GET" {
		signStr := request.URL.Query().Encode() + timeStamp
		return Sign(signStr, key), nil
	}
	contentType := request.Header.Get("Content-Type")
	signStr := ""
	switch {
	case strings.HasPrefix(contentType, "multipart/form-data"):
		{
			signStr = signStr + request.Form.Encode()
		}
	case strings.HasPrefix(contentType, "application/x-www-form-urlencoded"):
		{
			signStr = signStr + request.Form.Encode()
		}
	case strings.HasPrefix(contentType, "application/json"),
		strings.HasPrefix(contentType, "application/javascript"):
		{
			body := requestBody(request)
			signStr = signStr + request.Form.Encode() + JSONSorted(string(body))
		}
	default:
		{
			signStr = signStr + request.Form.Encode()
		}
	}
	signStr = signStr + timeStamp
	return Sign(signStr, key), nil
}

func requestBody(req *http.Request) []byte {
	if req.Body == nil {
		return []byte{}
	}
	var requestbody []byte
	var maxMemory = 1 << 26 //64MB
	safe := &io.LimitedReader{R: req.Body, N: int64(maxMemory)}
	if req.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(safe)
		if err != nil {
			return []byte{}
		}
		requestbody, _ = ioutil.ReadAll(reader)
	} else {
		requestbody, _ = ioutil.ReadAll(safe)
	}
	return requestbody
}
