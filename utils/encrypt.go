package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// 封装md5加密
func EncryMd5(s string) string {
	ctx := md5.New()
	ctx.Write([]byte(s))
	return hex.EncodeToString(ctx.Sum(nil))
}
