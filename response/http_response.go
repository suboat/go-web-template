package goresponse

type HTTPResponse struct {
	Tag    string      `json:"tag"`     // tag
	Result bool        `json:"success"` // 请求是否成功处理
	State  *State      `json:"state"`   // 状态元
	Meta   *Meta       `json:"meta"`    // 扩展元
	Data   interface{} `json:"data"`    // 核心数据
	Token  string      `json:"token"`   // 身份令牌
	Error  string      `json:"error"`   // 错误描述
}

func NewHTTPResponse() *HTTPResponse {
	return &HTTPResponse{
		Result: false,
		State:  NewState(),
		Meta:   nil,
		Data:   "",
		Error:  "",
	}
}

// 设置状态元
func (r *HTTPResponse) S(s *State) *HTTPResponse {
	r.State = s
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

// 设置核心数据
func (r *HTTPResponse) D(d interface{}) *HTTPResponse {
	if d != nil {
		r.Data = d
	}
	return r
}

// 设置新身份令牌
func (r *HTTPResponse) T(t string) *HTTPResponse {
	r.Token = t
	return r
}

// 设置错误描述
func (r *HTTPResponse) E(err error) *HTTPResponse {
	if err != nil {
		r.Error = err.Error()
	} else {
		r.Error = ""
	}
	return r
}

// 设置执行成功
func (r *HTTPResponse) Success() *HTTPResponse {
	r.State.Success()
	r.E(nil).Result = true
	return r
}

func (r *HTTPResponse) IsSuccess() bool {
	return (r.State.Status == Status_success || r.State.Status == Status_ignore)
}

func (r *HTTPResponse) IsInvalidToken() bool {
	return (r.State.Status == Status_invalid_token)
}

func (r *HTTPResponse) IsTimeoutToken() bool {
	return (r.State.Status == Status_token_timeout)
}
