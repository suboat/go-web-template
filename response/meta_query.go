package goresponse

import "net/http"

type MetaQuery struct {
	Meta
	Query string
	Order string
	Limit string
}

func NewMetaQuery(r *http.Request) *MetaQuery {
	return &MetaQuery{
		Query: GetMetaQuery(r),
		Order: GetMetaOrder(r),
		Limit: GetMetaLimit(r),
	}
}

func (m *MetaQuery) ValidQuery() bool {
	return len(m.Query) != 0
}

func (m *MetaQuery) ValidOrder() bool {
	return len(m.Order) != 0
}

func (m *MetaQuery) ValidLimit() bool {
	return len(m.Limit) != 0
}

func (m *MetaQuery) Valid() bool {
	return m.ValidQuery() || m.ValidOrder() || m.ValidLimit()
}
