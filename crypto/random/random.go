package random

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/lfin-open/lfin.utils.golang/encoding"
)

// GenerateRandomBytes Secure Ramdom 생성, 바이트로 리턴
func GenerateRandomBytes(size int) []byte {
	b := make([]byte, size)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil
	}
	return b
}

// GenerateRandomHexString Secure 랜덤바이트를 생성후 hex string 으로 리턴
func GenerateRandomHexString(size int) string {
	return hex.EncodeToString(GenerateRandomBytes(size))
}

// GenerateRandomBase64 Secure 랜덤바이트를 생성후 Base64 로 리턴
func GenerateRandomBase64(size int) string {
	return encoding.Base64EncodingFromByte(GenerateRandomBytes(size))
}
