package goresponse

type HTTPResponse struct {
	Tag    string      `json:"tag"`     // tag
	Result bool        `json:"success"` // 请求是否成功处理
	Status int         `json:"status"`  // 状态码
	Code   *Code       `json:"code"`    // 执行码
	Meta   *Meta       `json:"meta"`    // 扩展元
	Data   interface{} `json:"data"`    // 核心数据
	Error  string      `json:"error"`   // 错误描述
}

func NewHTTPResponse() *HTTPResponse {
	return &HTTPResponse{
		Result: false,
		Status: Status_fail,
		Code:   NewCode(),
		Meta:   nil,
		Data:   "",
		Error:  "",
	}
}

// 设置状态码
func (r *HTTPResponse) S(s int) *HTTPResponse {
	r.Status = s
	r.Result = r.IsSuccess()
	return r
}

// 设置扩展元
func (r *HTTPResponse) ME(m *Meta) *HTTPResponse {
	if m != nil {
		r.Meta = m
	}
	return r
}

// 设置执行码
func (r *HTTPResponse) C(c int) *HTTPResponse {
	r.Code.C(c)
	return r
}

// 设置执行描述
func (r *HTTPResponse) M(m int) *HTTPResponse {
	r.Code.M(m)
	return r
}

// 设置核心数据
func (r *HTTPResponse) D(d interface{}) *HTTPResponse {
	if d != nil {
		r.Data = d
	}
	return r
}

// 设置错误描述
func (r *HTTPResponse) E(err error) *HTTPResponse {
	if err != nil {
		r.Error = err.Error()
	}
	return r
}

// 设置执行成功
func (r *HTTPResponse) Success() *HTTPResponse {
	r.Status = Status_success
	r.Result = true
	r.Error = ""
	return r
}

func (r *HTTPResponse) IsSuccess() bool {
	return (r.Status == Status_success || r.Status == Status_ignore)
}

func (r *HTTPResponse) IsInvalidToken() bool {
	return (r.Status == Status_invalid_token)
}

func (r *HTTPResponse) IsTimeoutToken() bool {
	return (r.Status == Status_token_timeout)
}
