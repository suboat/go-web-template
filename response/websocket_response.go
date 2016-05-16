package goresponse

import (
	"encoding/json"
	"net/http"
)

type WSResponse struct {
	ResponseWriter http.ResponseWriter `json:"-"`
	RequestId      string              `json:"id"`    // 请求ID,直接回复
	URL            string              `json:"url"`   // 请求URL
	Data           interface{}         `json:"data"`  // 回复核心数据，推荐HTTPResponse
	Error          string              `json:"error"` // 错误码
}

func (r *WSResponse) Json() (s string) {
	if b, err := json.Marshal(r); err == nil {
		s = string(b)
	}
	return
}
