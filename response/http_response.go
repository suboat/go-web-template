package goresponse

type HTTPResponse struct {
	Tag     string      `json:"tag"`
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Code    *Code       `json:"code"`
	Meta    *Meta       `json:"meta"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

func NewHTTPResponse() *HTTPResponse {
	return &HTTPResponse{
		Success: false,
		Status:  Status_fail,
		Code:    NewCode(),
		Meta:    nil,
		Data:    "",
		Error:   "",
	}
}

func (this *HTTPResponse) S(s int) *HTTPResponse {
	this.Status = s
	this.Success = this.IsSuccess()
	return this
}

func (this *HTTPResponse) ME(m *Meta) *HTTPResponse {
	if m != nil {
		this.Meta = m
	}
	return this
}

func (this *HTTPResponse) C(c int) *HTTPResponse {
	this.Code.C(c)
	return this
}

func (this *HTTPResponse) M(m int) *HTTPResponse {
	this.Code.M(m)
	return this
}

func (this *HTTPResponse) D(d interface{}) *HTTPResponse {
	if d != nil {
		this.Data = d
	}
	return this
}

func (this *HTTPResponse) E(err error) *HTTPResponse {
	if err != nil {
		this.Error = err.Error()
	}
	return this
}

func (this *HTTPResponse) Success() *HTTPResponse {
	this.Status = Status_success
	this.Success = true
	this.Error = ""
	return this
}

func (this *HTTPResponse) IsSuccess() bool {
	return (this.Status == Status_success || this.Status == Status_ignore)
}

func (this *HTTPResponse) IsInvalidToken() bool {
	return (this.Status == Status_invalid_token)
}

func (this *HTTPResponse) IsTimeoutToken() bool {
	return (this.Status == Status_token_timeout)
}
