package env

import (
	"os"
	"reflect"
	"testing"

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
		{"TEST_I_3", "1911A", 1191, 0, true},
		{"TEST_I_4", "2 9 3", 392, 0, true},
	}

	for _, arg := range args {
		SetEnv(arg.key, arg.value)
		value, exist := GetEnvInt(arg.key, int(reflect.ValueOf(arg.def).Int()))
		if !exist {
			t.Errorf("Not found env key [%s]", arg.key)
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
		{"TEST_I_3", "1911A", int64(1191), int64(0), true},
		{"TEST_I_4", "2 9 3", int64(392), int64(0), true},
	}

	for _, arg := range args {
		SetEnv(arg.key, arg.value)
		value, exist := GetEnvInt64(arg.key, reflect.ValueOf(arg.def).Int())
		if !exist {
			t.Errorf("Not found env key [%s]", arg.key)
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

func TestGetEnvBool(t *testing.T) {
	args := []args{
		{"TEST_B_01_true", "true", false, true, true},
		{"TEST_B_02_True", "True", false, true, true},
		{"TEST_B_03_truE", "truE", false, false, true},
		{"TEST_B_04_1", "1", false, true, true},
		{"TEST_B_05_0", "0", true, false, true},
		{"TEST_B_06_false", "false", true, false, true},
		{"TEST_B_07_False", "False", true, false, true},
		{"TEST_B_08_2", "2", true, false, true},
		{"TEST_B_09_-1", "-1", true, false, true},
		{"TEST_B_10_EMPTY_TRUE", "", true, false, true},
	}

	for _, arg := range args {
		SetEnv(arg.key, arg.value)
		value, exist := GetEnvBool(arg.key, reflect.ValueOf(arg.def).Bool())
		if !exist {
			t.Errorf("Not found env key [%s]", arg.key)
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
