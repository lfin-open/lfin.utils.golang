package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/lfin-open/lfin.utils.golang/crypto/random"
	"strings"
)

// GenerateKey AES-256 키생성
func GenerateKey(size int) []byte {
	return random.GenerateRandomBytes(size)
}

func PaddingPKCS7(src []byte, size int) []byte {
	padSize := size - len(src)%size
	padByte := bytes.Repeat([]byte{byte(padSize)}, padSize)
	return append(src, padByte...)
}

func trimPKCS5(text []byte) []byte {
	padding := text[len(text)-1]
	return text[:len(text)-int(padding)]
}

type Crypto interface {
	Encrypt(plainText string) (string, error)
	DecryptToString(encryptedBase64 string) (string, error)
	DecryptToByte(encryptedBase64 string) ([]byte, error)
}

type AesCipher struct {
	iv  []byte
	key []byte
}

// Encrypt AesCipher 의 method 로 정의
func (c AesCipher) Encrypt(plainText string) (string, error) {
	if strings.TrimSpace(plainText) == "" {
		return plainText, nil
	}

	block, err := aes.NewCipher(c.key)
	if err != nil {
		return "", err
	}

	encrypter := cipher.NewCBCEncrypter(block, c.iv)
	paddedPlainText := PaddingPKCS7([]byte(plainText), encrypter.BlockSize())

	cipherText := make([]byte, len(paddedPlainText))
	// CryptBlocks 함수에 데이터(paddedPlainText)와 암호화 될 데이터를 저장할 슬라이스(cipherText)를 넣으면 암호화가 된다.
	encrypter.CryptBlocks(cipherText, paddedPlainText)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// DecryptToString AES-256 복호화후 string 으로 리턴
func (c AesCipher) DecryptToString(encryptedBase64 string) (string, error) {
	var result string
	resultB, err := c.DecryptToByte(encryptedBase64)
	if err == nil {
		result = string(resultB)
	}
	return result, err
}

// DecryptToByte  AES-256 복호화후 byte 로 리턴
func (c AesCipher) DecryptToByte(encryptedBase64 string) (result []byte, err error) {

	if strings.TrimSpace(encryptedBase64) == "" {
		return nil, nil
	}

	// throw error if panic occurred
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint(r))
		}
	}()

	decodedCipherText, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return nil, err
	}

	block, errN := aes.NewCipher(c.key)
	if err != nil {
		return nil, errN
	}

	decrypter := cipher.NewCBCDecrypter(block, c.iv)
	plainText := make([]byte, len(decodedCipherText))

	decrypter.CryptBlocks(plainText, decodedCipherText)
	result = trimPKCS5(plainText)

	return result, nil
}

// NewAesCipher AesCypher 를 생성해서 AesCipher 구조체 리턴
func NewAesCipher(iv, key []byte) (Crypto, error) {
	if keyLen := len(key); keyLen != aes.BlockSize*2 {
		return nil, aes.KeySizeError(keyLen)
	}

	if ivLen := len(iv); ivLen != aes.BlockSize {
		return nil, aes.KeySizeError(ivLen)
	}

	return &AesCipher{iv, key}, nil
}
