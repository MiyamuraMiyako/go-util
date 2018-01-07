package util

import (
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
