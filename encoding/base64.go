package encoding

import (
	b64 "encoding/base64"
	"github.com/lfin-open/lfin.utils.golang/strings"
)

// Base64EncodingFromString string 을 base64로 인코딩
func Base64EncodingFromString(s string) string {
	if strings.IsEmptyString(s) {
		return ""
	}
	return b64.StdEncoding.EncodeToString([]byte(s))
}

// Base64EncodingFromByte byte 을 base64로 인코딩
func Base64EncodingFromByte(b []byte) string {
	if len(b) < 1 {
		return ""
	}
	return b64.StdEncoding.EncodeToString(b)
}

// Base64DecodingToByte base64 디코딩, byte 로 리턴
func Base64DecodingToByte(s string) []byte {
	if strings.IsEmptyString(s) {
		return []byte("")
	}
	sDec, _ := b64.StdEncoding.DecodeString(s)
	return sDec
}

// Base64DecodingToString base64 디코딩, string 으로 리턴
func Base64DecodingToString(s string) string {
	return string(Base64DecodingToByte(s))
}
