package util

//SequenceEqual will check 2 arrays sequence equal.
func SequenceEqual(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

//CopyBytes is simple copy util
func CopyBytes(src []byte, srcIdx int, to []byte, dstIdx int, l int) {
	for ; l != 0; l-- {
		to[dstIdx] = src[srcIdx]
		srcIdx++
		dstIdx++
	}
}
