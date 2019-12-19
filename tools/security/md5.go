package security

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 生成md5加密
func Md5(src string) string {
	return getResult(src)
}

// Md5WithSalt 加密时简单加盐
func Md5WithSalt(src string, salt string) string {
	str := src + "#" + salt
	return getResult(str)
}

func getResult(src string) string {
	h := md5.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}
