package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// Sign 签名
func Sign(str, key string) string {
	hmac := func(key, data string) string {
		hmac := hmac.New(md5.New, []byte(key))
		hmac.Write([]byte(data))
		return hex.EncodeToString(hmac.Sum([]byte("")))
	}
	return hmac(key, JSONSorted(str))
}

// JSONSorted json排序字符串
func JSONSorted(str string) string {
	if len(str) <= 0 {
		return ""
	}
	b := []byte(str)
	//不是json字符串直接返回
	if !json.Valid(b) {
		return str
	}
	// 数组
	if strings.HasPrefix(str, "[") {
		o := []interface{}{}
		if e := json.Unmarshal(b, &o); e != nil {
			return str
		}
		return JSONWithSlice(o)
	}
	//对象
	o := map[string]interface{}{}
	if e := json.Unmarshal(b, &o); e != nil {
		return str
	}
	return JSONWithMap(o)
}

// JSONWithMap 把map排序并转成字符串
func JSONWithMap(m map[string]interface{}) string {
	v := reflect.ValueOf(m)
	if len(v.MapKeys()) <= 0 {
		return ""
	}
	return jsonWithMap(v)
}

func sortedKeys(keys []reflect.Value) []reflect.Value {
	sort.Slice(keys, func(i, j int) bool {
		return strings.Compare(keys[i].Interface().(string), keys[j].Interface().(string)) < 0
	})
	return keys
}

func jsonWithMap(v reflect.Value) string {
	items := []interface{}{}
	sortKeys := sortedKeys(v.MapKeys())
	for _, k := range sortKeys {
		v := v.MapIndex(k)
		t := v.Type()
		for t.Kind() == reflect.Ptr || t.Kind() == reflect.Interface {
			v = v.Elem()
			t = v.Type()
		}
		kind := t.Kind()
		switch kind {
		case reflect.Map:
			{
				str := fmt.Sprintf("\"%v\":%s", k.Interface(), jsonWithMap(v))
				items = append(items, str)
			}
		case reflect.Slice, reflect.Array:
			{
				str := fmt.Sprintf("\"%v\":%s", k.Interface(), JSONWithSlice(v.Interface().([]interface{})))
				items = append(items, str)
			}
		case reflect.String:
			{
				str := fmt.Sprintf("\"%v\":\"%s\"", k.Interface(), v.Interface().(string))
				items = append(items, str)
			}
		default:
			{
				str := fmt.Sprintf("\"%v\":%v", k.Interface(), v.Interface())
				items = append(items, str)
			}
		}
	}
	return fmt.Sprintf("{%s}", SliceJoin(items, ","))
}

// JSONWithSlice 转成字符串
func JSONWithSlice(s []interface{}) string {
	items := []interface{}{}
	for _, item := range s {
		v := reflect.ValueOf(item)
		t := reflect.TypeOf(item)
		for t.Kind() == reflect.Ptr || t.Kind() == reflect.Interface {
			v = v.Elem()
			t = v.Type()
		}
		kind := t.Kind()
		switch kind {
		case reflect.Map:
			{
				items = append(items, jsonWithMap(v))
			}
		case reflect.Slice, reflect.Array:
			{
				items = append(items, JSONWithSlice(v.Interface().([]interface{})))
			}
		case reflect.String:
			{
				str := fmt.Sprintf("\"%s\"", v.Interface().(string))
				items = append(items, str)
			}
		default:
			{
				items = append(items, v.Interface())
			}
		}
	}
	return fmt.Sprintf("[%s]", SliceJoin(items, ","))
}
