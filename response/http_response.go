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
