package introp

import (
	"syscall"
	"unsafe"
)

//ToPtr will get var or struct' address and return uintptr.
func ToPtr(i interface{}) uintptr {
	switch i.(type) {
	case int:
		return uintptr(i.(int))
	case string:
		s, _ := syscall.BytePtrFromString(i.(string))
		return uintptr(unsafe.Pointer(s))
	case []byte:
		return uintptr(unsafe.Pointer(&i.([]byte)[0]))
	default:
		return uintptr(unsafe.Pointer(&i))
	}
}
