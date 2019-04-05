package redis

import "testing"

func TestRedis(t *testing.T) {
	r, err := SET("test", "test", 60)
	if r == nil || err != nil {
		t.Error(r, err)
	}
	r, err = GET("test")
	if r == nil || err != nil {
		t.Error(r, err)
	}
	r, err = DEL("test")
	if r.(int64) == 0 || err != nil {
		t.Error(r, err)
	}
	r, err = GET("test")
	if r != nil || err != nil {
		t.Error(r, err)
	}
}
