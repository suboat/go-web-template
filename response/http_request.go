package response

import "net/http"

const (
	HeaderKeyQuery string = "Meta-Query"
)

func HEAD_QUERY_KEY() string {
	return HeaderKeyQuery
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
