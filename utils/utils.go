package utils

import (
	"encoding/json"
	"strings"
)

func JsonDecode(s string) (m map[string]interface{}, err error) {
	if len(s) == 0 {
		return map[string]interface{}{}, nil
	}
	var v interface{}
	dec := json.NewDecoder(strings.NewReader(s))
	err = dec.Decode(&v)
	if err != nil {
		return m, err
	}
	m = v.(map[string]interface{})
	return m, nil
}
