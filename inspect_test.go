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
