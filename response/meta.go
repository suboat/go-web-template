package goresponse

type Meta struct {
	Keys  map[string]interface{} `json:"keys"`
	Sorts []string               `json:"sorts"`
	Limit int                    `json:"limit"`
	Skip  int                    `json:"skip"`
	Total int64                  `json:"total"`
}
