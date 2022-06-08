package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(s string) string {
	b := []byte(s)
	m := md5.New()
	m.Write(b)
	return hex.EncodeToString(m.Sum(nil))
}
