package uuid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateUUID(t *testing.T) {
	uid := GenerateUUID()
	assert.Equal(t, 36, len(uid), fmt.Sprintf("uuid lenght not equual 36, length:%d", len(uid)))
	t.Logf("generated uuid=[%s]", uid)
}
