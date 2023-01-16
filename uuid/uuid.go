package uuid

import "github.com/google/uuid"

// GenerateUUID uuid 를 생성해서 리턴, 오류발생하면 ""
// format xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
func GenerateUUID() string {
	uid, _ := uuid.NewRandom()
	return uid.String()
}
