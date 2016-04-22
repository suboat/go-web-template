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
func (m *Meta) L(lmt int) *Meta {
	m.Limit = lmt
	return m
}

// 设置忽略数
func (m *Meta) S(skip int) *Meta {
	m.Skip = skip
	return m
}

// 设置总数
func (m *Meta) T(total int64) *Meta {
	m.Total = total
	return m
}
