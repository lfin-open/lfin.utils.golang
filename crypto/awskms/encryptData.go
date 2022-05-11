// Package awskms Copyright Amazon.com, Inc. or its affiliates. All Rights
// SPDX-License-Identifier: Apache-2.0
// https://github.com/awsdocs/aws-doc-sdk-examples/tree/main/gov2/kms/EncryptData
package awskms

import (
	"context"
	b64 "encoding/base64"
	"errors"
	"fmt"
	lstrings "github.com/lfin-open/lfin.utils.golang/strings"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

// KMSEncryptAPI defines the interface for the Encrypt function.
// We use this interface to test the function using a mocked service.
type KMSEncryptAPI interface {
	Encrypt(ctx context.Context,
		params *kms.EncryptInput,
		optFns ...func(*kms.Options)) (*kms.EncryptOutput, error)
}

// EncryptDataByKms encrypts some text using an AWS Key Management Service (AWS KMS) key (KMS key).
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, an EncryptOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to Encrypt.
func EncryptDataByKms(c context.Context, api KMSEncryptAPI, input *kms.EncryptInput) (*kms.EncryptOutput, error) {
	return api.Encrypt(c, input)
}

// EncryptData encrypts some text using an AWS Key Management Service (AWS KMS) key (KMS key).
// Inputs:
//     key is aws kms key id
//     s is plaintext
// Output:
//     If success, an EncryptOutput object containing the result of the service call and "".
//     Otherwise, nil and an error from the call to Encrypt.
func EncryptData(keyID, s string) (string, error) {

	if lstrings.IsEmptyString(keyID) {
		return "", errors.New("aws kms keyID is nil or empty")
	}

	if lstrings.IsEmptyString(strings.TrimSpace(s)) {
		return "", nil
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", fmt.Errorf("aws kms configuration error, %s", err)
	}

	client := kms.NewFromConfig(cfg)

	input := &kms.EncryptInput{
		KeyId:     &keyID,
		Plaintext: []byte(s),
	}

	result, err := EncryptDataByKms(context.TODO(), client, input)
	if err != nil {
		return "", fmt.Errorf("aws kms encrypt error, %s", err)
	}

	b64String := b64.StdEncoding.EncodeToString(result.CiphertextBlob)

	return b64String, nil
}
