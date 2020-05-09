package util

import (
	"C"
	"unsafe"
	"unicode/utf8"
)

//SubStr will split string using index and length
func SubStr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

//Substr will split string using index
func Substr(s string, pos int) string {
	runes := []rune(s)
	l := pos + (utf8.RuneCountInString(s) - pos)
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

//UintptrToString uinptr转GoString
func UintptrToString(ptr uintptr) string {
	return C.GoString((*C.char)(unsafe.Pointer(ptr)))
}