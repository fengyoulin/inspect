package inspect

import (
	"reflect"
	"unsafe"
)

type eFace struct {
	typ, data unsafe.Pointer
}

// PackEFace from type and data
func PackEFace(typ reflect.Type, data unsafe.Pointer) (i interface{}) {
	e := (*eFace)(unsafe.Pointer(&i))
	t := (*eFace)(unsafe.Pointer(&typ))
	e.typ = t.data
	e.data = data
	return
}
