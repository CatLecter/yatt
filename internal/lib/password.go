package lib

import (
	"crypto/sha512"
	"encoding/base64"
)

func HashPassword(password string) string {
	// TODO: "соль" брать из конфига.
	salt := "salt"
	hasher := sha512.New()
	hasher.Write([]byte(password))
	hashPassword := base64.URLEncoding.EncodeToString(hasher.Sum([]byte(salt)))
	return hashPassword
}
