package encoding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// base64 encode/decode online
//  https://ostermiller.org/calc/encode.html

var (
	plain  = "This Is PlainText 1234567890_)!@#$%^&*()+}{"
	base64 = "VGhpcyBJcyBQbGFpblRleHQgMTIzNDU2Nzg5MF8pIUAjJCVeJiooKSt9ew=="
)

func TestBase64EncodingFromByte(t *testing.T) {
	assert.Equal(t, base64, Base64EncodingFromByte([]byte(plain)))
}

func TestBase64EncodingFromString(t *testing.T) {
	assert.Equal(t, base64, Base64EncodingFromString(plain))
}

func TestBase64DecodingToString(t *testing.T) {
	assert.Equal(t, plain, Base64DecodingToString(base64))
}

func TestBase64DecodingToByte(t *testing.T) {
	assert.Equal(t, []byte(plain), Base64DecodingToByte(base64))
}
