package word

import (
	"crypto/md5"
	"encoding/hex"
)

func XKeyGen(payload ...string) string {
	var key string

	for _, p := range payload {
		if len(key) > 0 {
			key += "."
		}
		key += p
	}
	md5 := md5.New()

	md5.Write([]byte(key))
	return hex.EncodeToString(md5.Sum(nil))
}
