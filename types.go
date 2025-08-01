package inspect

import (
	"reflect"
	"sync"
	"unsafe"
)

var (
	types     map[string]reflect.Type
	typesOnce sync.Once
)

func initTypes() {
	typesOnce.Do(func() {
		types = make(map[string]reflect.Type)

		typ := reflect.TypeOf(0)
		face := (*iface)(unsafe.Pointer(&typ))

		sections, offset := typelinks()
		for i, offs := range offset {
			rodata := sections[i]
			for _, off := range offs {
				face.data = resolveTypeOff(rodata, off)
				if typ.Kind() == reflect.Ptr && len(typ.Elem().Name()) > 0 {
					types[typ.String()] = typ
					types[typ.Elem().String()] = typ.Elem()
				}
			}
		}
	})
}
