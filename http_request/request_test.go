package http_request

import (
	"github.com/utils-library/http_request/types"
	"testing"
)

func TestRequest(t *testing.T) {
	req := MakeRequest("http://139.155.72.59/touch-fish/one", types.REQUEST_METHOD__GET)
	req.AddParam("name", "coltzhang")
	result, err := req.Do()
	if err != nil {
		t.Fatalf("err: %+v\n", err)
	}
	t.Logf("result:%+v\n", result)
	t.Logf("string_result:%+v\n", string(result))
}
