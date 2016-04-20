package goresponse

// Tag key
const (
	// normal
	TagAccession = "accession"
	TagUid       = "uid"
	// mean value
	TagValNo   = `%no%`   // 不等于 {"name":"%no%jack"}
	TagValLike = `%like%` // like {"name":"%like%ac"}
	TagValLt   = `%lt%`   // 小于 {"createTime":"%lt%2000-01-01 00:00:00 +0800"}
	TagValLte  = `%lte%`  // 小于等于
	TagValGt   = `%gt%`   // 大于
	TagValGte  = `%gte%`  // 大于等于
	// mean key
	TagQueryKeyOr  = `%or%`  // {"%or%":[{"name":"jack","name":"tim"}]}
	TagQueryKeyAnd = `%and%` // {"%and%":[{"name":"jack","age":"13"}]}
	TagQueryKeyIn  = `%in%`  // {"%in%":[{"sex":"man","age":"male"}]}
)
