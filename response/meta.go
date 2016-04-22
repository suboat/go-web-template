package goresponse

type Meta struct {
	Limit int   `json:"limit"` // 限制数
	Skip  int   `json:"skip"`  // 忽略数
	Total int64 `json:"total"` // 总数
}

func NewMeta(lmt, skip int) *Meta {
	return &Meta{Limit: lmt, Skip: skip}
}

// 设置限制数
func (s *Meta) L(lmt int) *Meta {
	s.Limit = lmt
	return s
}

// 设置忽略数
func (s *Meta) S(skip int) *Meta {
	s.Skip = skip
	return s
}

// 设置总数
func (s *Meta) T(total int64) *Meta {
	s.Total = total
	return s
}
