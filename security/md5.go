package security

import (
	"crypto/md5"
	_ "encoding/base64"
	"encoding/hex"
)

func Md5(str string) string {
	salt := "bms_salt"
	hash := md5.New()
	hash.Write([]byte(str))
	pw := hash.Sum(nil)
	hash.Write(pw)
	hash.Write([]byte(salt))
	return hex.EncodeToString(hash.Sum(nil)) //md5 -> 加salt -> md5
}
