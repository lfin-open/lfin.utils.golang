package reflection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	Id       string `json:"id" binding:"required" `
	Name     string `json:"name" `
	Password string `json:"password" `
	Age      int64  `json:"age" `
}

var (
	id   = "myid"
	name = "LFin Developer 엘핀개발자"
	age  = int64(20)

	user  = User{Id: id, Name: name, Age: age}
	user2 User
)

func TestGetObjFieldToInt64(t *testing.T) {
	// 있는필드 가져오기
	r, exist := GetObjFieldToInt64(&user, "Age")
	t.Logf("value:%v, exist:%v", r, exist)
	assert.Equal(t, age, r)

	// 없는필드 가져오기
	r2, exist2 := GetObjFieldToInt64(&user, "NoField")
	t.Logf("value:%v, exist:%v", r2, exist2)
	assert.Equal(t, false, exist2)
}

func TestGetObjFieldToString(t *testing.T) {
	r, exist := GetObjFieldToString(&user, "Name")
	t.Logf("value:%v, exist:%v", r, exist)
	assert.Equal(t, name, r)
}

func TestSetObjFieldToInt64(t *testing.T) {
	// 있는필드 세팅하기
	exist := SetObjFieldToInt64(&user2, "Age", age)
	t.Logf("setted object:%v", user2)
	t.Logf("setted value:%v", user2.Age)
	assert.Equal(t, age, user2.Age)
	assert.Equal(t, true, exist)

	// 필드가 없으면
	exist = SetObjFieldToInt64(&user2, "NoField", age)
	assert.Equal(t, false, exist)
}

func TestSetObjFieldToString(t *testing.T) {
	// 있는필드 세팅하기
	exist := SetObjFieldToString(&user2, "Name", name)
	t.Logf("setted object:%v", user2)
	t.Logf("setted value:%v", user2.Name)
	assert.Equal(t, name, user2.Name)
	assert.Equal(t, true, exist)

	// 필드가 없으면
	exist = SetObjFieldToString(&user2, "NoField", "this filed is no exist")
	assert.Equal(t, false, exist)
}
