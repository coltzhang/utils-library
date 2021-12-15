package redis

import (
	"testing"
)

func TestRedis(t *testing.T) {
	//exp := time.Duration(300) * time.Second
	//err := SetCache("date", "2021-12-15", exp)
	//if err != nil {
	//	t.Errorf("err:%+v\n", err)
	//}
	//t.Logf("ok")
	val, err := GetCache("date")
	if err != nil {
		t.Errorf("err:%+v\n", err)
	}
	t.Logf("val:%v", val)
}
