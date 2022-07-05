package net

import (
	"bytes"

	"github.com/lfin-open/lfin.utils.golang/strings"
)

// ConvertToMacAddressStyle mac address style 로 변환하기
//  80CA4B4B6F0B --> 80:CA:4B:4B:6F:0B or 80-CA-4B-4B-6F-0B
func ConvertToMacAddressStyle(s string, v string) string {
	var buffer bytes.Buffer
	if len(s) < 3 || strings.IsEmptyString(v) {
		return s
	}

	for i := 0; i < len(s); i++ {
		if i == 0 || i%2 == 1 {
			buffer.WriteString(s[i : i+1])
		} else {
			buffer.WriteString(v)
			buffer.WriteString(s[i : i+1])
		}
	}

	return buffer.String()
}
