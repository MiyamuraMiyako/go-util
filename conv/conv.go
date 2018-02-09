package conv

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func GB2312BytsToString(byts []byte, idx, l int) string {
	dat, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader(bytes.TrimRight(byts[idx:idx+l], "\x00")), simplifiedchinese.GBK.NewDecoder()))
	return string(dat)
}

func BytsToString(byts []byte, idx, l int) string {
	return string(bytes.TrimRight(byts[idx:idx+l], "\x00"))
}

func UTF8StringToGBKByts(str string) []byte {
	d, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader(([]byte)(str)), simplifiedchinese.GBK.NewEncoder()))
	return d
}
