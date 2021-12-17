package utils

import "testing"

func TestUtils(t *testing.T) {
	str := `{"a": 1, "b": "b"}`
	m, err := JsonDecode(str)
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	for k, v := range m {
		t.Logf("k:%v v:%v", k, v)
	}
	t.Logf("m:%v", m)
}
