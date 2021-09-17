package helper

import (
	"crypto/md5"
	"encoding/hex"
)

func StrMd5Encode(str string)string{
	h := md5.New()
	h.Write([]byte("asdrtycvbzxcjkliofdsG"))
	return hex.EncodeToString(h.Sum([]byte(str)))
}
