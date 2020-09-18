package crypt

import (
	"crypto/md5"
	"encoding/hex"
)

/**
 * 生成md5的摘要
 */
func Md5(value string) string {
	h := md5.New()
	h.Write([]byte(value))
	return hex.EncodeToString(h.Sum(nil))
}
