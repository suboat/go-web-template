package response

type HTTPResponse struct {
	Tag     string      `json:"tag"`     // tag
	Success bool        `json:"success"` // 请求是否成功处理
	State   *State      `json:"state"`   // 状态元
	Meta    *Meta       `json:"meta"`    // 扩展元
	Data    interface{} `json:"data"`    // 核心数据
	Error   string      `json:"error"`   // 错误描述
}

func NewHTTPResponse() *HTTPResponse {
	return &HTTPResponse{
		Success: false,
		State:   NewState(),
		Meta:    nil,
		Data:    nil,
		Error:   "",
	}
}

// 设置状态元
func (r *HTTPResponse) S(s *State) *HTTPResponse {
	if s != nil {
		r.State = s
	}
	r.Success = r.IsSuccess()
	return r
}

// 设置扩展元
func (r *HTTPResponse) ME(m *MetaQuery) *HTTPResponse {
	if m != nil {
		r.Meta = &m.Meta
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

// 设置错误描述
func (r *HTTPResponse) E(err error) *HTTPResponse {
	if err != nil {
		r.Error = err.Error()
	} else {
		r.Error = ""
	}
	return r
}
