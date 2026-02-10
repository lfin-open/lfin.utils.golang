package encoding

import (
	b64 "encoding/base64"

	"github.com/lfin-open/lfin.utils.golang/strings"
)

// Base64EncodingFromString encodes a string to base64
func Base64EncodingFromString(s string) string {
	if strings.IsEmptyString(s) {
		return ""
	}
	return b64.StdEncoding.EncodeToString([]byte(s))
}

// Base64EncodingFromByte encodes bytes to base64
func Base64EncodingFromByte(b []byte) string {
	if len(b) < 1 {
		return ""
	}
	return b64.StdEncoding.EncodeToString(b)
}

// Base64DecodingToByte decodes base64 and returns bytes
func Base64DecodingToByte(s string) []byte {
	if strings.IsEmptyString(s) {
		return []byte("")
	}
	sDec, _ := b64.StdEncoding.DecodeString(s)
	return sDec
}

// Base64DecodingToString decodes base64 and returns a string
func Base64DecodingToString(s string) string {
	return string(Base64DecodingToByte(s))
}
