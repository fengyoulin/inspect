# inspect #

The golang function `inspect.TypeOf` works just like `Class.forName` in Java. It's useful when needing the Type of a unexported type.

Example:
```go
package main

import (
	"fmt"
	"github.com/fengyoulin/inspect"
)

func main() {
	typ := inspect.TypeOf("runtime.g")
	if typ != nil {
		for i := 0; i < typ.NumField(); i++ {
			f := typ.Field(i)
			fmt.Println(f.Name, f.Type.Name())
		}
	}
}
```
