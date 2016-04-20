package goresponse

type HTTPResponse struct {
	Tag    string      `json:"tag"`
	Result bool        `json:"success"`
	Status int         `json:"status"`
	Code   *Code       `json:"code"`
	Meta   *Meta       `json:"meta"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
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

func (r *HTTPResponse) S(s int) *HTTPResponse {
	r.Status = s
	r.Result = r.IsSuccess()
	return r
}

func (r *HTTPResponse) ME(m *Meta) *HTTPResponse {
	if m != nil {
		r.Meta = m
	}
	return r
}

func (r *HTTPResponse) C(c int) *HTTPResponse {
	r.Code.C(c)
	return r
}

func (r *HTTPResponse) M(m int) *HTTPResponse {
	r.Code.M(m)
	return r
}

func (r *HTTPResponse) D(d interface{}) *HTTPResponse {
	if d != nil {
		r.Data = d
	}
	return r
}

func (r *HTTPResponse) E(err error) *HTTPResponse {
	if err != nil {
		r.Error = err.Error()
	}
	return r
}

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
