package response

import (
	"net/http"
	"strings"
)

const (
	HeaderKeyQuery string = "Meta-Query"
)

func HEAD_QUERY_KEY() string {
	return HeaderKeyQuery
}

func getHTTPHeaderValue(r *http.Request, key string) string {
	if r == nil {
	} else if key = strings.TrimSpace(key); len(key) == 0 {
	} else if r.Header == nil {
	} else if value := r.Header.Get(key); value != "" {
		return value
	} else if value = r.Header.Get(strings.ToLower(key)); value != "" {
		return value
	} else if value = r.Header.Get(strings.ToUpper(key)); value != "" {
		return value
	}
	return ""
}

func GetMetaQuery(r *http.Request) string {
	return getHTTPHeaderValue(r, HeaderKeyQuery)
}
