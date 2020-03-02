package inspect

import "testing"

func TestTypeOf(t *testing.T) {
	typ := TypeOf("runtime.g")
	if typ == nil {
		t.Error("runtime.g not found")
	}
	typ = TypeOf("*runtime.g")
	if typ == nil {
		t.Error("*runtime.g not found")
	}
}

func TestTypes(t *testing.T) {
	ts := Types()
	if len(ts) == 0 || len(ts[0]) == 0 {
		t.Error("get types failed")
	}
}
