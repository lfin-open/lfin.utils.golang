package hash

import (
	"crypto/sha512"
	"encoding/hex"
	"strings"
)

// HashAES512ToUpper AES512 로 해싱 (대문자로 변환)
func HashAES512ToUpper(s string) string {
	h := sha512.New()
	h.Write([]byte(s))
	bh := h.Sum(nil)

	return strings.ToUpper(hex.EncodeToString(bh))
}
