package ymtcrypto

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(keys ...string) string {
	h := md5.New()
	for _, v := range keys {
		h.Write([]byte(v))
	}
	md5Bts := h.Sum(nil)
	md5Str := hex.EncodeToString(md5Bts)
	return md5Str
}
