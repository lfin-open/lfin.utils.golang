package env

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type args struct {
	key      string
	value    string
	def      interface{}
	expected interface{}
	exist    bool
}

func TestSetEnv(t *testing.T) {
	args := []args{
		{"TEST_S_1", "123456 789ABC", "DEFAULT", "123456 789ABC", true},
		{"TEST_S_2", "한글ABC", "한글", "한글ABC", true},
	}

	for _, arg := range args {
		SetEnv(arg.key, arg.value)
		value, exist := os.LookupEnv(arg.key)
		if !exist {
			t.Errorf("Not found env key [%s]", arg.key)
		} else {
			assert.Equal(t, arg.expected, value)
		}
	}
}

func TestGetEnv(t *testing.T) {
	args := []args{
		{"TEST_S_1", "123456 789ABC", "DEFAULT", "123456 789ABC", true},
		{"TEST_S_2", "한글ABC", "한글", "한글ABC", true},
	}

	for _, arg := range args {
		SetEnv(arg.key, arg.value)
		value, exist := GetEnv(arg.key, reflect.ValueOf(arg.def).String())
		if !exist {
			t.Errorf("Not found env key [%s]", arg.key)
		}
		if !assert.Equal(t, arg.expected, value) {
			t.Errorf("key:%s, value:%s, expected:%v, actual:%v", arg.key, arg.value, arg.expected, value)
		}
	}

	// default value if not found
	v, exist := GetEnv("NOT_FOUND_S", "DEFAULT")
	assert.Equal(t, false, exist)
	assert.Equal(t, "DEFAULT", v)
}

func TestGetEnvInt(t *testing.T) {
	args := []args{
		{"TEST_I_1", "123456", 654321, 123456, true},
		{"TEST_I_2", "-91", -19, -91, true},
		{"TEST_I_3", "1911A", 1191, 1191, false},
		{"TEST_I_4", "2 9 3", 392, 392, false},
	}

	for _, arg := range args {
		SetEnv(arg.key, arg.value)
		value, exist := GetEnvInt(arg.key, int(reflect.ValueOf(arg.def).Int()))
		if !assert.Equal(t, arg.exist, exist) {
			t.Errorf("key:%s, value:%s, expected:%v, actual:%v", arg.key, arg.value, arg.exist, exist)
		}
		if !assert.Equal(t, arg.expected, value) {
			t.Errorf("key:%s, value:%s, expected:%v, actual:%v", arg.key, arg.value, arg.expected, value)
		}
	}

	// default value if not found
	v, exist := GetEnvInt("NOT_FOUND_I", 99)
	assert.Equal(t, false, exist)
	assert.Equal(t, 99, v)
}

func TestGetEnvInt64(t *testing.T) {
	args := []args{
		{"TEST_I_1", "123456", int64(654321), int64(123456), true},
		{"TEST_I_2", "-91", int64(-19), int64(-91), true},
		{"TEST_I_3", "1911A", int64(1191), int64(1191), false},
		{"TEST_I_4", "2 9 3", int64(392), int64(392), false},
	}

	for _, arg := range args {
		SetEnv(arg.key, arg.value)
		value, exist := GetEnvInt64(arg.key, reflect.ValueOf(arg.def).Int())
		if !assert.Equal(t, arg.exist, exist) {
			t.Errorf("key:%s, value:%s, expected:%v, actual:%v", arg.key, arg.value, arg.exist, exist)
		}
		if !assert.Equal(t, arg.expected, value) {
			t.Errorf("key:%s, value:%s, expected:%v, actual:%v", arg.key, arg.value, arg.expected, value)
		}
	}

	// default value if not found
	v, exist := GetEnvInt64("NOT_FOUND_I", int64(99))
	assert.Equal(t, false, exist)
	assert.Equal(t, int64(99), v)
}

func TestGetEnvFloat32(t *testing.T) {
	args := []args{
		{"TEST_F32_1", "123456", float32(654321), float32(123456), true},
		{"TEST_F32_2", "-91", float32(-19), float32(-91), true},
		{"TEST_F32_3", "1911A", float32(1191), float32(1191), false},
		{"TEST_F32_4", "2 9 3", float32(392), float32(392), false},
	}

	for _, arg := range args {
		SetEnv(arg.key, arg.value)
		value, exist := GetEnvFloat32(arg.key, float32(reflect.ValueOf(arg.def).Float()))
		if !assert.Equal(t, arg.exist, exist) {
			t.Errorf("key:%s, value:%s, expected:%v, actual:%v", arg.key, arg.value, arg.exist, exist)
		}
		if !assert.Equal(t, arg.expected, value) {
			t.Errorf("key:%s, value:%s, expected:%v, actual:%v", arg.key, arg.value, arg.expected, value)
		}
	}

	// default value if not found
	v, exist := GetEnvFloat32("NOT_FOUND_I", float32(99))
	assert.Equal(t, false, exist)
	assert.Equal(t, float32(99), v)
}

func TestGetEnvFloat64(t *testing.T) {
	args := []args{
		{"TEST_F64_1", "123456", float64(654321), float64(123456), true},
		{"TEST_F64_2", "-91", float64(-19), float64(-91), true},
		{"TEST_F64_3", "1911A", float64(1191), float64(1191), false},
		{"TEST_F64_4", "2 9 3", float64(392), float64(392), false},
	}

	for _, arg := range args {
		SetEnv(arg.key, arg.value)
		value, exist := GetEnvFloat64(arg.key, reflect.ValueOf(arg.def).Float())
		if !assert.Equal(t, arg.exist, exist) {
			t.Errorf("key:%s, value:%s, expected:%v, actual:%v", arg.key, arg.value, arg.exist, exist)
		}
		if !assert.Equal(t, arg.expected, value) {
			t.Errorf("key:%s, value:%s, expected:%v, actual:%v", arg.key, arg.value, arg.expected, value)
		}
	}

	// default value if not found
	v, exist := GetEnvFloat64("NOT_FOUND_I", float64(99))
	assert.Equal(t, false, exist)
	assert.Equal(t, float64(99), v)
}

func TestGetEnvBool(t *testing.T) {
	args := []args{
		{"TEST_B_01_true", "true", false, true, true},
		{"TEST_B_02_True", "True", false, true, true},
		{"TEST_B_03_truE", "truE", false, false, false},
		{"TEST_B_04_1", "1", false, true, true},
		{"TEST_B_05_0", "0", true, false, true},
		{"TEST_B_06_false", "false", true, false, true},
		{"TEST_B_07_False", "False", true, false, true},
		{"TEST_B_08_2", "2", false, false, false},
		{"TEST_B_09_-1", "-1", false, false, false},
		{"TEST_B_10_EMPTY_TRUE", "", false, false, false},
		{"TEST_B_11_PARSE_ERROR", "ABC", false, false, false},
	}

	for _, arg := range args {
		SetEnv(arg.key, arg.value)
		value, exist := GetEnvBool(arg.key, reflect.ValueOf(arg.def).Bool())
		if !assert.Equal(t, arg.exist, exist) {
			t.Errorf("key:%s, value:%s, expected:%v, actual:%v", arg.key, arg.value, arg.exist, exist)
		}
		if !assert.Equal(t, arg.expected, value) {
			t.Errorf("key:%s, value:%s, expected:%v, actual:%v", arg.key, arg.value, arg.expected, value)
		}
	}

	vf, exist := GetEnvBool("NOT_FOUND_BF", false)
	assert.Equal(t, false, exist)
	assert.Equal(t, false, vf)

	vt, exist := GetEnvBool("NOT_FOUND_BT", true)
	assert.Equal(t, false, exist)
	assert.Equal(t, true, vt)
}

func TestGetEnvTimeDuration(t *testing.T) {
	args := []args{
		{"TEST_T_01_10sec", "10s", "10s", "10s", true},
		{"TEST_T_02_60m", "60m", "60m", "60m", true},
		{"TEST_T_03_5ms", "5ms", "5ms", "5ms", true},
		{"TEST_T_04_0sec", "0s", "0s", "0s", true},
		{"TEST_T_05_ABC", "ABC", "9s", "9s", false},
	}

	for _, arg := range args {
		SetEnv(arg.key, arg.value)
		def, _ := time.ParseDuration(fmt.Sprintf("%v", reflect.ValueOf(arg.def).String()))
		exp, _ := time.ParseDuration(fmt.Sprintf("%v", reflect.ValueOf(arg.expected).String()))
		value, exist := GetEnvTimeDuration(arg.key, def)
		if !assert.Equal(t, arg.exist, exist) {
			t.Errorf("key:%s, value:%s, expected:%v, actual:%v", arg.key, arg.value, arg.exist, exist)
		}
		if !assert.Equal(t, exp.Seconds(), value.Seconds()) {
			t.Errorf("key:%s, value:%s, expected:%v, actual:%v", arg.key, arg.value, exp.Seconds(), value.Seconds())
		}

		v, exist := GetEnvTimeDuration("NOT_FOUND_BF", time.Duration(99))
		assert.Equal(t, false, exist)
		assert.Equal(t, time.Duration(99).Seconds(), v.Seconds())

	}
}
