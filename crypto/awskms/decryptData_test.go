package awskms

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/stretchr/testify/assert"
	"testing"
)

//var (
//	globalConfig   Config
//	configFileName = "configForTest.json"
//)
//
//func init() {
//	fmt.Println("start init function")
//
//	err := populateConfiguration()
//	if err != nil {
//		log.Fatal(err)
//	}
//}

var (
	// 복호화 후 비교할 데이터
	text    = "이것은 평문입니다. This is plaintext. 一二三四五六七八九十 🏆 😀 🤘 !@#$%^&*()_+"
	encData = "AQICAHgBcE3gf6aL4bfLqYodS8EwpDdOIqhoI+UmzQ6FiE42GAE45A7xCIFsZlTVhjCLu602AAAAyzCByAYJKoZIhvcNAQcGoIG6MIG3AgEAMIGxBgkqhkiG9w0BBwEwHgYJYIZIAWUDBAEuMBEEDF1QCwEoujLAP5rxrAIBEICBgxN0FU3Wuu3qO7dk1C+o1KOczQ4jtOkE40ybrRbSpca/e3mIJ5rDzoUNcSDWOKw3/q9pelK3rIUpyi/2GP0F3PxUwS7R0VCuohMeos+j8ifd6+lctV3gi8BCjeTXR9LSdFwnu8HgdN41iIUF//GuIHp32aZCjs55Fmo6M41zytHx83Ig"
)

type KMSDecryptImpl struct{}

func (KMSDecryptImpl) Decrypt(ctx context.Context,
	params *kms.DecryptInput,
	optFns ...func(*kms.Options)) (*kms.DecryptOutput, error) {

	// Blah, blah, blah
	//plainText := []byte{66, 108, 97, 104, 44, 32, 98, 108, 97, 104, 44, 32, 98, 108, 97, 104}
	plainText := []byte("이것은 Dummuy 에서 복호화한 값. I am dummy")

	output := &kms.DecryptOutput{
		Plaintext: plainText,
	}

	return output, nil
}

func TestDecryptDataByKms(t *testing.T) {

	// 실 KMS 를 타지 않고 위에 정의한 dummy Decrypt 를 사용함.
	api := &KMSDecryptImpl{}
	blob := []byte(encData)
	input := &kms.DecryptInput{
		CiphertextBlob: blob,
	}

	result, err := DecryptDataByKms(context.Background(), api, input)
	if err != nil {
		t.Error("error kms decrypt.", err)
	} else {
		// t.Logf("decoded data. KeyId:%s, Encryption Argorithm:%s", *result.KeyId, result.EncryptionAlgorithm)
		t.Logf("decrypted PlainText:[%s]", result.Plaintext)
	}
}

func TestDecryptData(t *testing.T) {

	t.Logf("encrypted data:[%s]", encData)
	result, err := DecryptData(encData)

	if err != nil {
		t.Error("error kms decrypt.", err)
	} else {
		t.Log("decrypted plaintext")
		t.Logf("[%s]", result)
		assert.Equal(t, text, result)
	}

	encData = ""
	result, err = DecryptData(encData)
	t.Logf("encrypted org:[%s] decrypted:[%s] error:[%v]", encData, result, err)

	assert.Equal(t, encData, result)
}
