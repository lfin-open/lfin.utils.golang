package reflection

import "reflect"

// GetObjFieldToString object(interface{}) 의 필드값을 string 으로 리턴
//
// Parameters:
//  - obj: 대상 object interface{}
//  - key: 찾을 키 이름
// Return:
//  - string: 필드가 있는경우 string value
//  - bool: 필드존재여부 true: 필드있음, false:필드없음
func GetObjFieldToString(obj interface{}, key string) (string, bool) {

	r := reflect.ValueOf(obj)
	v := reflect.Indirect(r).FieldByName(key)

	if v.IsValid() {
		return v.String(), true
	} else {
		return "", false
	}
}

// GetObjFieldToInt64 object(interface{}) 의 필드값을 int64 로 리턴
//
// Parameters:
//  - obj: 대상 object interface{}
//  - key: 찾을 키 이름
// Return:
//  - int64: 필드가 있는경우 int64 value
//  - bool: 필드존재여부 true: 필드있음, false:필드없음
func GetObjFieldToInt64(obj interface{}, key string) (int64, bool) {

	r := reflect.ValueOf(obj)
	v := reflect.Indirect(r).FieldByName(key)

	if v.IsValid() {
		return v.Int(), true
	} else {
		return 0, false
	}
}

// SetObjFieldToString object(interface{}) 의 string 값을 저장
//
// Parameters:
//  - obj: 대상 object interface{}
//  - key: 찾을 키 이름
//  - value: 저장할 값 string
// Return:
//  - bool: 필드존재여부 true: 필드있음, false:필드없음
func SetObjFieldToString(obj interface{}, key, value string) bool {

	r := reflect.ValueOf(obj)
	v := reflect.Indirect(r).FieldByName(key)

	if v.IsValid() {
		v.SetString(value)
		return true
	} else {
		return false
	}

}

// SetObjFieldToInt64 object(interface{}) 의 int64 값을 저장
//
// Parameters:
//  - obj: 대상 object interface{}
//  - key: 찾을 키 이름
//  - value: 저장할 값 int64
// Return:
//  - bool: 필드존재여부 true: 필드있음, false:필드없음
func SetObjFieldToInt64(obj interface{}, key string, value int64) bool {

	r := reflect.ValueOf(obj)
	v := reflect.Indirect(r).FieldByName(key)

	if v.IsValid() {
		v.SetInt(value)
		return true
	} else {
		return false
	}

}
