package response

// 设置状态元的状态码
func (r *HTTPResponse) SS(status int) *HTTPResponse {
	r.State.S(status)
	r.Success = r.IsSuccess()
	return r
}

func (r *HTTPResponse) SF() *HTTPResponse {
	r.State.S(Status_fail)
	r.Success = false
	return r
}

func (r *HTTPResponse) SFF() *HTTPResponse {
	r.State.S(Status_fail_frequently)
	r.Success = false
	return r
}

// 设置执行成功
func (r *HTTPResponse) Finish() *HTTPResponse {
	r.State.Success()
	r.E(nil).Success = true
	return r
}

func (r *HTTPResponse) IsSuccess() bool {
	return (r.State.Status == Status_success || r.State.Status == Status_ignore)
}

func (r *HTTPResponse) IsInvalidToken() bool {
	return (r.State.Status == Status_invalid_token)
}

func (r *HTTPResponse) IsTimeoutToken() bool {
	return (r.State.Status == Status_invalid_token)
}
