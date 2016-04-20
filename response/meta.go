package goresponse

type Meta struct {
	Keys      map[string]interface{} `json:"keys"`
	Sorts     []string               `json:"sorts"`
	Limit     int                    `json:"limit"`
	Skip      int                    `json:"skip"`
	Total     int64                  `json:"total"`
	StartDate string                 `json:"start"`
	EndDate   string                 `json:"end"`
}

func NewMeta(lmt, skip int) *Meta {
	return &Meta{Limit: lmt, Skip: skip}
}

func (s *Meta) K(keys map[string]interface{}) *Meta {
	s.Keys = keys
	return s
}

func (s *Meta) S(sorts []string) *Meta {
	s.Sorts = sorts
	return s
}

func (s *Meta) T(total int64) *Meta {
	s.Total = total
	return s
}

func (s *Meta) Time(st, et string) *Meta {
	s.StartDate = st
	s.EndDate = et
	return s
}
