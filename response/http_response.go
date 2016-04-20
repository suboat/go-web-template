package goresponse

type HTTPResponse struct {
	Tag    string      `json:"tag"`
	Result bool        `json:"result"`
	Meta   *Meta       `json:"meta"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
	// TODO: 还有很多要补充
}
