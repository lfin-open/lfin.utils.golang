package env

import (
	"os"
	"strconv"
)

// SetEnv 환경변수 설정
// key: 환경변수명
// value: 환경변수값
func SetEnv(key, value string) error {
	return os.Setenv(key, value)
}

// GetEnv 환경변수 가져오기
// key: 환경변수명
// def: 기본값
// return: 환경변수값, 키가 없으면 def 값
func GetEnv(key, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return def
}

// GetEnvInt 환경변수 가져오기
// key: 환경변수명
// def: 기본값
// return: 환경변수값, 키가 없으면 def 값
func GetEnvInt(key string, def int) int {
	if value, ok := os.LookupEnv(key); ok {
		iv, _ := strconv.Atoi(value)
		return iv
	}
	return def
}

// GetEnvInt64 환경변수 가져오기
// key: 환경변수명
// def: 기본값
// return: 환경변수값, 키가 없으면 def 값
func GetEnvInt64(key string, def int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		iv, _ := strconv.ParseInt(value, 10, 64)
		return iv
	}
	return def
}

// GetEnvBool 환경변수 가져오기
// key: 환경변수명
// def: 기본값
// return: 환경변수값, 키가 없으면 def 값, 1,True,true-->true, 그외 false
func GetEnvBool(key string, def bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		bv, _ := strconv.ParseBool(value)
		return bv
	}
	return def
}
