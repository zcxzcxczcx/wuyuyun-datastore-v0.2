package common

import (
	"net/http"
	"encoding/json"
)

/**
返回json数据
response：主要数据部分，任意类型
status：http状态码
 */
func JsonResponse(response interface{}, status int, w http.ResponseWriter) {
	httpStatusText := http.StatusText(status)
	if httpStatusText == "" {
		//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		//return
		status = http.StatusInternalServerError
	}

	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status) // 发送一个状态代码的HTTP响应头。
	w.Write(data)
}
