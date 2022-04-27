/*
 * Copyright (c) 2022 LFin and others.
 *
 * All rights reserved.
 *
 *
 * Contributors:
 *    Ted KIM
 *
 */

package utils

import (
	"fmt"
	"os"
	"strings"
)

// GetHostname 시스템의 호스트이름을 리턴
func GetHostname() string {
	// get hostname
	hostname, err := os.Hostname()
	if err != nil {
		hostname = ""
	}
	return hostname
}

// IsEmptyString string 이 empty 인지 체크
func IsEmptyString(s string) bool {
	if strings.TrimSpace(s) == "" {
		return true
	} else {
		return false
	}
}

// I64ToS int64 를 String 으로 변환
func I64ToS(i64 int64) string {
	return fmt.Sprintf("%d", i64)
}
