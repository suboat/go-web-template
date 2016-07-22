package response

import "net/http"

type WSRequest struct {
	Request   *http.Request `json:"-"`
	RequestId string        `json:"id"`      // 请求ID,自由分配
	Methods   []string      `json:"methods"` // 请求方法集
	URL       string        `json:"url"`     // 请求URL
	Data      interface{}   `json:"data"`    // 请求核心数据
	Ignore    bool          `json:"ignore"`  // 是否忽略
}
