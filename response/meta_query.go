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
