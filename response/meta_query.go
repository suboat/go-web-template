package goresponse

import "net/http"

type MetaQuery struct {
	Meta
	Value string
}

func NewMetaQuery(r *http.Request) *MetaQuery {
	return &MetaQuery{
		Value: GetMetaQuery(r),
	}
}

func (m *MetaQuery) Valid() bool {
	return len(m.Value) != 0
}
