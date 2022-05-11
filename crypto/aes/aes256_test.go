package aes

import (
	"encoding/hex"
	"fmt"
	"github.com/lfin-open/lfin.utils.golang/encoding"
	"github.com/stretchr/testify/assert"
	"testing"
)

// AES Encryption and Decryption Online Tool(Calculator)
// https://www.devglan.com/online-tools/aes-encryption-decryption

// Encryption Key Generator
// https://www.allkeysgenerator.com/Random/Security-Encryption-Key-Generator.aspx

var (
	//iv     = "hVkYp3s6v9y$B&E)"
	ivHex = "68566b59703373367639792442264529"
	//key    = "D(G+KbPeShVmYp3s6v9y$B&E)H@McQfT"
	keyHex = "4428472b4b6250655368566d5970337336763979244226452948404d63516654"
	plain  = "This is Plain Text 이것은 평문이예요. !@#$%^&*() 1234567890-="
	encB64 = "qyHAnaQL3U1BlrZiJIXBDRJYibMoHlYd28KDR6U9vEbruvcNlDQpSfMzMOw+wXsJnu/fInvWUqAE8exKcpp+noH9+VZIDtEg57fJOjVnSsg="
)

func TestGenerateKey(t *testing.T) {
	size16 := 16
	keyB16 := GenerateKey(size16)
	keyH16 := hex.EncodeToString(keyB16)

	fmt.Printf("test aes GenerateKey(%d)\n", size16)
	fmt.Printf(" - key       size=[%d]\n", size16)
	fmt.Printf(" - generated size=[%d], keyHex=[%s]\n", len(keyB16), keyH16)
	fmt.Printf(" - keyBase64     =[%s]\n", encoding.Base64EncodingFromByte(keyB16))

	assert.Equal(t, size16, len(keyB16))

	size32 := 16 * 2
	keyB32 := GenerateKey(size32)
	keyH32 := hex.EncodeToString(keyB32)

	fmt.Printf("test aes GenerateKey(%d)\n", size32)
	fmt.Printf(" - key       size=[%d]\n", size32)
	fmt.Printf(" - generated size=[%d], keyHex=[%s]\n", len(keyB32), keyH32)
	fmt.Printf(" - keyBase64     =[%s]\n", encoding.Base64EncodingFromByte(keyB32))

	assert.Equal(t, size32, len(keyB32))

}

func TestAesCipher_Encrypt(t *testing.T) {
	ivB, _ := hex.DecodeString(ivHex)
	keyB, _ := hex.DecodeString(keyHex)
	c, _ := NewAesCipher(ivB, keyB)

	encB64Result, _ := c.Encrypt(plain)

	fmt.Printf("test aes encrypt\n")
	fmt.Printf("  - ivHex=[%s], keyHex=[%s]\n", ivHex, keyHex)
	fmt.Printf("  - plain       text=[%s]\n", plain)
	fmt.Printf("  - expected  base64=[%s]\n", encB64)
	fmt.Printf("  - encrypted base64=[%s]\n", encB64Result)

	assert.Equal(t, encB64, encB64Result)

}

func TestAesCipher_Decrypt(t *testing.T) {
	ivB, _ := hex.DecodeString(ivHex)
	keyB, _ := hex.DecodeString(keyHex)
	c, _ := NewAesCipher(ivB, keyB)

	plainResult, _ := c.DecryptToString(encB64)

	fmt.Printf("test aes decrypt\n")
	fmt.Printf("  - ivHex=[%s], keyHex=[%s]\n", ivHex, keyHex)
	fmt.Printf("  - encrypted base64=[%s]\n", encB64)
	fmt.Printf("  - expected   plain=[%s]\n", plain)
	fmt.Printf("  - decrypted  plain=[%s]\n", plainResult)

	assert.Equal(t, plain, plainResult)
}

func TestAesCipher_DecryptToString(t *testing.T) {
	ivB, _ := hex.DecodeString(ivHex)
	keyB, _ := hex.DecodeString(keyHex)
	c, _ := NewAesCipher(ivB, keyB)

	plainResult, _ := c.DecryptToString(encB64)

	fmt.Printf("test aes decrypt\n")
	fmt.Printf("  - ivHex=[%s], keyHex=[%s]\n", ivHex, keyHex)
	fmt.Printf("  - encrypted base64=[%s]\n", encB64)
	fmt.Printf("  - expected   plain=[%s]\n", plain)
	fmt.Printf("  - decrypted  plain=[%s]\n", plainResult)

	assert.Equal(t, plain, plainResult)

}
