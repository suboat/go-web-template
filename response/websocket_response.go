package goresponse

import "net/http"

type WSResponse struct {
	ResponseWriter http.ResponseWriter `json:"-"`
	RequestId      string              `json:"id"`    // 请求ID,直接回复
	URL            string              `json:"url"`   // 请求URL
	Data           interface{}         `json:"data"`  // 回复核心数据
	Error          string              `json:"error"` // 错误码
}
