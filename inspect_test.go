package inspect_test

import (
	"github.com/fengyoulin/inspect"
	"testing"
)

func TestTypeOf(t *testing.T) {
	typ := inspect.TypeOf("runtime.g")
	if typ == nil {
		t.Error("runtime.g not found")
	}
	typ = inspect.TypeOf("*runtime.g")
	if typ == nil {
		t.Error("*runtime.g not found")
	}
}

func TestTypes(t *testing.T) {
	ts := inspect.Types()
	if len(ts) == 0 || len(ts[0]) == 0 {
		t.Error("get types failed")
	}
}
