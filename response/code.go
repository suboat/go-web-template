package goresponse

type Code struct {
	Code    int `json:"code"`    // 执行码
	Message int `json:"message"` // 执行描述码
	Status  int `json:"status"`  // 状态码
}

func NewCode() *Code {
	return &Code{
		Code:    0,
		Message: 0,
		Status:  Status_fail,
	}
}

func (c *Code) C(code int) *Code {
	c.Code = code
	return c
}

func (c *Code) M(message int) *Code {
	c.Message = message
	return c
}

func (c *Code) S(status int) *Code {
	c.Status = status
	return c
}
