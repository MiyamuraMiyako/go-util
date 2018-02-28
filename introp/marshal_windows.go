package introp

import (
	"fmt"
	"syscall"
	"unsafe"
)

//ToPtr will get var or struct' address and return uintptr.
func ToPtr(i interface{}) uintptr {
	switch i.(type) {
	case int:
		return uintptr(i.(int))
	case string:
		return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(i.(string))))
	case []byte:
		return uintptr(unsafe.Pointer(&i.([]byte)[0]))
	default:
		return uintptr(unsafe.Pointer(&i))
	}
}

//GetProcAddr return Proc's address.
func GetProcAddr(handle syscall.Handle, name string) uintptr {
	addr, err := syscall.GetProcAddress(handle, name)
	if err != nil {
		panic(fmt.Sprintf("%s %s", name, err.Error()))
	}
	return addr
}
