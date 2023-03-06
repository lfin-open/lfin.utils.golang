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
// return:
//  string: 환경변수값, 키가 없으면 def 값
//  bool: 키가 있으면 true, 없으면 false
func GetEnv(key, def string) (string, bool) {
	if value, ok := os.LookupEnv(key); ok {
		return value, ok
	}
	return def, false
}

// GetEnvInt 환경변수 가져오기
// key: 환경변수명
// def: 기본값
// return:
//	int: 환경변수값, 키가 없으면 def 값
//	bool: 키가 있으면 true, 없으면 false
func GetEnvInt(key string, def int) (int, bool) {
	if value, ok := os.LookupEnv(key); ok {
		iv, _ := strconv.Atoi(value)
		return iv, ok
	}
	return def, false
}

// GetEnvInt64 환경변수 가져오기
// key: 환경변수명
// def: 기본값
// return:
//  int64: 환경변수값, 키가 없으면 def 값
//  bool: 키가 있으면 true, 없으면 false
func GetEnvInt64(key string, def int64) (int64, bool) {
	if value, ok := os.LookupEnv(key); ok {
		iv, _ := strconv.ParseInt(value, 10, 64)
		return iv, ok
	}
	return def, false
}

// GetEnvBool 환경변수 가져오기
// key: 환경변수명
// def: 기본값
// return:
//  bool: 환경변수값, 키가 없으면 def 값, 1,True,true-->true, 그외 false
//  bool: 키가 있으면 true, 없으면 false
func GetEnvBool(key string, def bool) (bool, bool) {
	if value, ok := os.LookupEnv(key); ok {
		bv, _ := strconv.ParseBool(value)
		return bv, ok
	}
	return def, false
}
