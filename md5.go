package gophp

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 Calculate the md5 hash of a string
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
