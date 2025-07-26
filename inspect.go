package inspect

import (
	"reflect"
	"unsafe"
)

// TypeOf find the type by package path and name
func TypeOf(pathName string) reflect.Type {
	initTypes()
	if typ, ok := types[pathName]; ok {
		return typ
	}
	return nil
}

// Types return all types in typelinks
func Types() (ts [][]reflect.Type) {
	initTypes()
	typ := reflect.TypeOf(0)
	face := (*iface)(unsafe.Pointer(&typ))
	sections, offset := typelinks()
	ts = make([][]reflect.Type, 0, len(offset))
	for i, offs := range offset {
		section := sections[i]
		list := make([]reflect.Type, 0, len(offs))
		for _, off := range offs {
			face.data = unsafe.Pointer(uintptr(section) + uintptr(off))
			list = append(list, typ)
		}
		ts = append(ts, list)
	}
	return
}
