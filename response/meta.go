package goresponse

type Meta struct {
	Keys      map[string]interface{} `json:"keys"`  // 关键字
	Sorts     []string               `json:"sorts"` // 排序方式
	Limit     int                    `json:"limit"` // 数量限制
	Skip      int                    `json:"skip"`  // 忽略数量
	Total     int64                  `json:"total"` // 总数
	StartDate string                 `json:"start"` // 起始时间
	EndDate   string                 `json:"end"`   // 终止时间
}

func NewMeta(lmt, skip int) *Meta {
	return &Meta{Limit: lmt, Skip: skip}
}

// 设置关键字
func (s *Meta) K(keys map[string]interface{}) *Meta {
	s.Keys = keys
	return s
}

// 设置排序方式
func (s *Meta) S(sorts []string) *Meta {
	s.Sorts = sorts
	return s
}

// 设置总数
func (s *Meta) T(total int64) *Meta {
	s.Total = total
	return s
}

// 设置起始终止时间
func (s *Meta) Time(st, et string) *Meta {
	s.StartDate = st
	s.EndDate = et
	return s
}
