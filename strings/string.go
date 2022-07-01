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

package strings

import (
	"fmt"
	"strings"
)

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
