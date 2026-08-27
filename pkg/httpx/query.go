package httpx

import (
	"net/http"
	"strconv"
)

// QueryString 从查询参数读取字符串，缺省返回默认值。
func QueryString(r *http.Request, key, def string) string {
	if v := r.URL.Query().Get(key); v != "" {
		return v
	}
	return def
}

// QueryInt 从查询参数读取整数，解析失败或非正数返回默认值。
func QueryInt(r *http.Request, key string, def int) int {
	v := r.URL.Query().Get(key)
	if v == "" {
		return def
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return n
}

// QueryFloat 从查询参数读取浮点数，解析失败返回默认值。
func QueryFloat(r *http.Request, key string, def float64) float64 {
	v := r.URL.Query().Get(key)
	if v == "" {
		return def
	}
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return def
	}
	return f
}

// QueryBool 从查询参数读取布尔值，解析失败返回默认值。
func QueryBool(r *http.Request, key string, def bool) bool {
	v := r.URL.Query().Get(key)
	if v == "" {
		return def
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		return def
	}
	return b
}
