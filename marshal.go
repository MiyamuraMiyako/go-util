package util

import (
	"fmt"
	"syscall"
	"unsafe"
)

func IntPtr(n int) uintptr {
	ToPtr(true)
	return uintptr(n)
}

func UTF16Ptr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}

func ToPtr(i interface{}) uintptr {
	return uintptr(unsafe.Pointer(&i))
}

func GetProcAddr(handle syscall.Handle, name string) uintptr {
	addr, err := syscall.GetProcAddress(handle, name)
	if err != nil {
		panic(fmt.Sprintf("%s %s", name, err.Error()))
	}
	return addr
}
