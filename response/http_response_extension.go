package goresponse

// 设置状态元的状态码
func (r *HTTPResponse) SS(status int) *HTTPResponse {
	r.State.S(status)
	r.Result = r.IsSuccess()
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
