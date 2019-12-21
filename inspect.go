package inspect

import "reflect"

// TypeOf find the type by package path and name
func TypeOf(pathName string) reflect.Type {
	if typ, ok := types[pathName]; ok {
		return typ
	}
	return nil
}
