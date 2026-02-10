package uuid

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateUUID(t *testing.T) {
	uid := GenerateUUID()
	assert.Equal(t, 36, len(uid), fmt.Sprintf("uuid lenght not equual 36, length:%d", len(uid)))
	t.Logf("generated uuid=[%s]", uid)
}

func TestGenerateUUIDv7(t *testing.T) {
	// 기본 생성 테스트
	uid := GenerateUUIDv7()
	assert.Equal(t, 32, len(uid), fmt.Sprintf("uuid v7 length not equal 32, length:%d", len(uid)))
	assert.NotContains(t, uid, "-", "uuid v7 should not contain hyphens")
	t.Logf("generated uuid v7=[%s]", uid)

	// 순차 보장 테스트 - 시간순으로 생성된 UUID v7은 정렬되어야 함
	uuids := make([]string, 10)
	for i := 0; i < 10; i++ {
		uuids[i] = GenerateUUIDv7()
		time.Sleep(1 * time.Millisecond) // 시간 차이를 보장하기 위해 약간의 대기
	}

	// 생성된 순서대로 정렬되어 있는지 확인
	for i := 0; i < len(uuids)-1; i++ {
		assert.True(t, strings.Compare(uuids[i], uuids[i+1]) < 0,
			fmt.Sprintf("UUID v7 should be sequential: %s should be less than %s", uuids[i], uuids[i+1]))
	}

	t.Logf("sequential uuid v7 test passed, first=[%s], last=[%s]", uuids[0], uuids[len(uuids)-1])

	// 고유성 테스트
	uniqueMap := make(map[string]bool)
	for i := 0; i < 100; i++ {
		uid := GenerateUUIDv7()
		assert.False(t, uniqueMap[uid], "duplicate uuid v7 generated")
		uniqueMap[uid] = true
	}
	t.Logf("uniqueness test passed with 100 uuid v7s")
}

func TestGenerateUID8(t *testing.T) {
	// 기본 생성 테스트
	uid := GenerateShortUID8()
	assert.Equal(t, 8, len(uid), fmt.Sprintf("uid8 length not equal 8, length:%d", len(uid)))
	assert.NotContains(t, uid, "-", "uid8 should not contain hyphens")
	t.Logf("generated uid8=[%s]", uid)

	// 고유성 테스트
	uniqueMap := make(map[string]bool)
	for i := 0; i < 1000; i++ {
		uid := GenerateShortUID8()
		assert.False(t, uniqueMap[uid], "duplicate uid8 generated")
		uniqueMap[uid] = true
		assert.Equal(t, 8, len(uid), "all uid8 should have length 8")
	}
	t.Logf("uniqueness test passed with 1000 uid8s")

	// Hex 문자만 포함하는지 확인
	uid = GenerateShortUID8()
	for _, c := range uid {
		assert.True(t, (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f'),
			fmt.Sprintf("uid8 should only contain hex characters, found: %c", c))
	}
	t.Logf("hex character validation passed")
}
