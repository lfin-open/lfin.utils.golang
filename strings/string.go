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
	"regexp"
	"strings"
)

// IsEmptyString checks if a string is empty or contains only whitespace
func IsEmptyString(s string) bool {
	if strings.TrimSpace(s) == "" {
		return true
	} else {
		return false
	}
}

// I64ToS converts int64 to string
func I64ToS(i64 int64) string {
	return fmt.Sprintf("%d", i64)
}

// ToSnakeCase converts a string to snake_case (spaces are removed)
// ex) TestToSnakeCase -> test_to_snake_case
func ToSnakeCase(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	if len(str) < 1 {
		return ""
	}
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
