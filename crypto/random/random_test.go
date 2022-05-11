package random

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/lfin-open/lfin.utils.golang/encoding"
	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomBytes(t *testing.T) {
	// 16byte
	rb := GenerateRandomBytes(aes.BlockSize)
	fmt.Printf("generated random block\n - block size:%d\n - hex:[%s]\n - base64:[%s]\n", aes.BlockSize, hex.EncodeToString(rb), encoding.Base64EncodingFromByte(rb))
	assert.Equal(t, 16, len(rb))
}

func TestGenerateRandomHexString(t *testing.T) {
	rb := GenerateRandomHexString(aes.BlockSize)
	fmt.Printf("generated random block\n - block size:%d\n - hex:[%s]\n", aes.BlockSize, rb)
	assert.Equal(t, aes.BlockSize*2, len(rb))
}

func TestGenerateRandomBase64(t *testing.T) {
	rb := GenerateRandomBase64(aes.BlockSize)
	fmt.Printf("generated random block\n - block size:%d\n - base64:[%s]\n", aes.BlockSize, rb)
	assert.Equal(t, 24, len(rb))
}
