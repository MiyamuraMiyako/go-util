package util

import (
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

//FileExist will check file or dir exist.
func FileExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//RandomHex 随机十六进制字符串
func RandomHex(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandInt(65, 90))
	}
	return string(bytes)
}

//RandInt 随机整数
func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
