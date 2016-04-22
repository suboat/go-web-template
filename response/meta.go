package goresponse

type Meta struct {
	Limit int   `json:"limit"` // 数量限制
	Skip  int   `json:"skip"`  // 忽略数量
	Total int64 `json:"total"` // 总数
}

func NewMeta(lmt, skip int) *Meta {
	return &Meta{Limit: lmt, Skip: skip}
}

// 设置总数
func (s *Meta) T(total int64) *Meta {
	s.Total = total
	return s
}
