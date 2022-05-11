package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Online Bcrypt Hash Generator & Checker
// https://bcrypt-generator.com/

var (
	password = "ThisIsPlassword &*^%$^"
	// cost           = 10
	hashedPassword = "$2a$10$Ye1Whun2EiJR1tIC65VGBO4FWnTchdgEY13u5RN.mECGzz3SSbIsK"
)

func TestCompareHashBcryptAndPassword(t *testing.T) {
	assert.Equal(t, true, CompareHashBcryptAndPassword(password, hashedPassword))
}

func TestHashBcryptPassword(t *testing.T) {
	hashedPassword2, _ := HashBcryptPassword(password, 10)
	assert.Equal(t, true, CompareHashBcryptAndPassword(password, hashedPassword2))
}
