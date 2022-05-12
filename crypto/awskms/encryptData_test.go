package awskms

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kms"
)

type KMSEncryptImpl struct{}

type Config struct {
	KeyID   string `json:"KeyID"`
	Text    string `json:"Text"`
	EncData string `json:"EncData"`
}

var (
	globalConfig   Config
	configFileName = "configForTest.json"
)

func init() {
	fmt.Println("start init function")

	err := populateConfiguration()
	if err != nil {
		log.Fatal(err)
	}
}

func (KMSEncryptImpl) Encrypt(_ context.Context,
	_ *kms.EncryptInput,
	_ ...func(*kms.Options)) (*kms.EncryptOutput, error,
) {
	blob := []byte(globalConfig.Text)

	output := &kms.EncryptOutput{
		CiphertextBlob: blob,
	}

	return output, nil
}

func populateConfiguration() error {
	content, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return err
	}

	text := string(content)

	err = json.Unmarshal([]byte(text), &globalConfig)
	if err != nil {
		return err
	}

	if globalConfig.KeyID == "" || globalConfig.Text == "" {
		msg := "You must supply a value for KeyID  " + configFileName
		return errors.New(msg)
	}

	return nil
}

func TestEncryptTextByKms(t *testing.T) {
	t.Logf("KeyId:[%s]", globalConfig.KeyID)
	t.Logf("Text:[%s]", globalConfig.Text)

	api := &KMSEncryptImpl{}
	input := &kms.EncryptInput{
		KeyId:     &globalConfig.KeyID,
		Plaintext: []byte(globalConfig.Text),
	}
	result, err := EncryptDataByKms(context.Background(), api, input)
	if err != nil {
		t.Logf("error kms encrypt. %v", err)
		t.Logf("please check, posibbly increct KeyId:[%s]", globalConfig.KeyID)
	} else {
		t.Log("encrypted Blob (base64 byte array):")
		t.Log(result.CiphertextBlob)
	}
}

func TestEncryptText(t *testing.T) {

	result, err := EncryptData(globalConfig.KeyID, globalConfig.Text)

	t.Logf("KeyId:[%s]", globalConfig.KeyID)
	t.Logf("Text:[%s]", globalConfig.Text)

	if err != nil {
		t.Logf("error kms encrypt. %v", err)
		t.Logf("please check, posibbly increct KeyId:[%s]", globalConfig.KeyID)
	} else {
		t.Log("encrypted data (base64)")
		t.Log(result)
	}

}
