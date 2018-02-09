package introp

import (
	"unsafe"
)

//ToPtr will get var or struct' address and return uintptr.
func ToPtr(i interface{}) uintptr {
	switch i.(type) {
	case int:
		return uintptr(i.(int))
	case string:
		fallthrough
	default:
		return uintptr(unsafe.Pointer(&i))
	}
}
