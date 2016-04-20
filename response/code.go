package goresponse

type Code struct {
	Code    int `json:"code"`
	Message int `json:"message"`
}

func NewCode() *Code {
	return new(Code)
}

func (c *Code) C(code int) *Code {
	c.Code = code
	return c
}

func (c *Code) M(message int) *Code {
	c.Message = message
	return c
}
