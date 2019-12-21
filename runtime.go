package inspect

import "unsafe"

//go:linkname typelinks reflect.typelinks
func typelinks() (sections []unsafe.Pointer, offset [][]int32)

//go:linkname resolveTypeOff reflect.resolveTypeOff
func resolveTypeOff(rtype unsafe.Pointer, off int32) unsafe.Pointer

type iface struct {
	tab  unsafe.Pointer
	data unsafe.Pointer
}
