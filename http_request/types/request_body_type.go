package types

// openapi:enum
//go:generate tools gen enum RequestBodyType
type RequestBodyType uint8

// 请求Body类型
const (
	REQUEST_BODY_TYPE_UNKNOWN RequestBodyType = iota
	REQUEST_BODY_TYPE__JSON                   // json
	REQUEST_BODY_TYPE__FORM                   // 表单
)

func (v RequestBodyType) String() string {
	switch v {
	case REQUEST_BODY_TYPE_UNKNOWN:
		return ""
	case REQUEST_BODY_TYPE__JSON:
		return "JSON"
	case REQUEST_BODY_TYPE__FORM:
		return "FORM"
	}
	return "UNKNOWN"
}
