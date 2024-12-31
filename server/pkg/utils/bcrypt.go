package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// SHA256Hash 使用 bcrypt 对密码进行加密
func SHA256Hash(password string, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password + salt))
	hashedPassword := hash.Sum(nil)
	return hex.EncodeToString(hashedPassword)
}
