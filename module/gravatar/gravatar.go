package gravatar

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Generate(email string) string {
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)
	hash := md5.Sum([]byte(email))
	return hex.EncodeToString(hash[:])
}
