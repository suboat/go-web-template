package goresponse

import "net/http"

const (
	HeaderKeyQuery string = "Meta-Query"
	HeaderKeyOrder        = "Meta-Order"
	HeaderKeyLimit        = "Meta-Limit"
)

type MetaQuery struct {
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

func getHTTPHeaderValue(r *http.Request, key string) string {
	if r == nil {
	} else if len(key) == 0 {
	} else if r.Header == nil {
	} else {
		return r.Header.Get(key)
	}
	return ""
}

func GetMetaQuery(r *http.Request) string {
	return getHTTPHeaderValue(r, HeaderKeyQuery)
}

func GetMetaOrder(r *http.Request) string {
	return getHTTPHeaderValue(r, HeaderKeyOrder)
}

func GetMetaLimit(r *http.Request) string {
	return getHTTPHeaderValue(r, HeaderKeyLimit)
}
