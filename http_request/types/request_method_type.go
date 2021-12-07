package types

type RequestMethod uint8

// 请求类型
const (
	REQUEST_METHOD_UNKNOWN RequestMethod = iota
	REQUEST_METHOD__GET                  // GET
	REQUEST_METHOD__POST                 // POST
	REQUEST_METHOD__PUT                  // PUT
	REQUEST_METHOD__DELETE               // DELETE
)

func (v RequestMethod) String() string {
	switch v {
	case REQUEST_METHOD_UNKNOWN:
		return ""
	case REQUEST_METHOD__GET:
		return "GET"
	case REQUEST_METHOD__POST:
		return "POST"
	case REQUEST_METHOD__PUT:
		return "PUT"
	case REQUEST_METHOD__DELETE:
		return "DELETE"
	}
	return "UNKNOWN"
}
