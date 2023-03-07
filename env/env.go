package env

import (
	"os"
	"strconv"
	"time"
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
//
//	string: 환경변수값, 키가 없으면 def 값
//	bool: 키가 있으면 true, 없으면 false
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
//
//	int: 환경변수값, 키가 없으면 def 값
//	bool: 키가 있으면 true, 없거나 int 로 변환할 수 없으면 false
func GetEnvInt(key string, def int) (int, bool) {
	if value, ok := os.LookupEnv(key); ok {
		iv, err := strconv.Atoi(value)
		if err != nil {
			return def, false
		}
		return iv, ok
	}
	return def, false
}

// GetEnvInt64 환경변수 가져오기
// key: 환경변수명
// def: 기본값
// return:
//
//	int64: 환경변수값, 키가 없으면 def 값
//	bool: 키가 있으면 true, 없거나 int64 로 변환할 수 없으면 false
func GetEnvInt64(key string, def int64) (int64, bool) {
	if value, ok := os.LookupEnv(key); ok {
		iv, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return def, false
		}
		return iv, ok
	}
	return def, false
}

// GetEnvFloat32 환경변수 가져오기
// key: 환경변수명
// def: 기본값
// return:
//
//	float32: 환경변수값, 키가 없으면 def 값
//	bool: 키가 있으면 true, 없거나 float32 로 변환할 수 없으면 false
func GetEnvFloat32(key string, def float32) (float32, bool) {
	if value, ok := os.LookupEnv(key); ok {
		fv, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return def, false
		}
		return float32(fv), ok
	}
	return def, false
}

// GetEnvFloat64 환경변수 가져오기
// key: 환경변수명
// def: 기본값
// return:
//
//	float64: 환경변수값, 키가 없으면 def 값
//	bool: 키가 있으면 true, 없거나 float64 로 변환할 수 없으면 false
func GetEnvFloat64(key string, def float64) (float64, bool) {
	if value, ok := os.LookupEnv(key); ok {
		fv, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return def, false
		}
		return fv, ok
	}
	return def, false
}

// GetEnvBool 환경변수 가져오기
// key: 환경변수명
// def: 기본값
// return:
//
//	bool: 환경변수값, 키가 없으면 def 값, 1,True,true-->true, 그외 false
//	bool: 키가 있으면 true, 없거나 bool 로 변환할 수 없으면 false
func GetEnvBool(key string, def bool) (bool, bool) {
	if value, ok := os.LookupEnv(key); ok {
		bv, err := strconv.ParseBool(value)
		if err != nil {
			return def, false
		}
		return bv, ok
	}
	return def, false
}

// GetEnvTimeDuration 환경변수 가져오기
// key: 환경변수명
// def: 기본값
// return:
//
//	time.Duration: 환경변수값, 키가 없으면 def 값
//	bool: 키가 있으면 true, 없거나 duration 으로 변환할 수 없으면 false
func GetEnvTimeDuration(key string, def time.Duration) (time.Duration, bool) {
	if value, ok := os.LookupEnv(key); ok {
		duration, err := time.ParseDuration(value)
		if err != nil {
			return def, false
		}
		return duration, ok
	}
	return def, false
}
