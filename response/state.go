package goresponse

type State struct {
	Code    int `json:"code"`    // 执行码
	Message int `json:"message"` // 执行描述码
	Status  int `json:"status"`  // 状态码
}

func NewState() *State {
	return &State{
		Code:    1,
		Message: 0,
		Status:  Status_fail,
	}
}

// 设置执行码
func (c *State) C(code int) *State {
	c.Code = code
	return c
}

// 设置执行描述码
func (c *State) M(message int) *State {
	c.Message = message
	return c
}

// 设置状态码
func (c *State) S(status int) *State {
	c.Status = status
	return c
}

func (c *State) Success() *State {
	return c.C(0).M(0).S(Status_success)
}
