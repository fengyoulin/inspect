package inspect_test

import (
	"github.com/fengyoulin/inspect"
	"reflect"
	"testing"
	"unsafe"
)

func TestPackEFace(t *testing.T) {
	n := int64(123)
	p := unsafe.Pointer(&n)
	e := inspect.PackEFace(reflect.TypeOf(int64(0)), p)
	if v, ok := e.(int64); !ok || v != n {
		t.Error(ok, v)
	}
}
