package uuid

import (
	"strings"

	"github.com/google/uuid"
)

// GenerateUUID uuid 를 생성해서 리턴, 오류발생하면 ""
// format xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
func GenerateUUID() string {
	uid, _ := uuid.NewRandom()
	return uid.String()
}

// GenerateUUIDv7 시간 기반 순차 보장되는 UUID v7을 생성해서 리턴 (하이픈 제거)
// UUID v7은 타임스탬프를 포함하여 자연스럽게 정렬 가능
// format xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx (32자, 하이픈 없음)
func GenerateUUIDv7() string {
	uid, _ := uuid.NewV7()
	return strings.ReplaceAll(uid.String(), "-", "")
}

// GenerateShortUID8 8자리 고유 식별자를 생성해서 리턴 (하이픈 제거)
// UUID의 앞 8자리만 사용
// format xxxxxxxx (8자, 하이픈 없음)
func GenerateShortUID8() string {
	uid, _ := uuid.NewRandom()
	// UUID의 앞 8자리만 추출
	return strings.ReplaceAll(uid.String(), "-", "")[:8]
}
