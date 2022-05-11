// Package awskms Copyright Amazon.com, Inc. or its affiliates. All Rights
// SPDX-License-Identifier: Apache-2.0
// https://github.com/awsdocs/aws-doc-sdk-examples/tree/main/gov2/kms/DecryptData
package awskms

import (
	"context"
	b64 "encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/lfin-open/lfin.utils.golang/strings"
)

// KMSDecryptAPI defines the interface for the Decrypt function.
// We use this interface to test the function using a mocked service.
type KMSDecryptAPI interface {
	Decrypt(ctx context.Context,
		params *kms.DecryptInput,
		optFns ...func(*kms.Options)) (*kms.DecryptOutput, error)
}

// DecryptDataByKms decrypts some text that was encrypted with an AWS Key Management Service (AWS KMS) key (KMS key).
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a DecryptOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to Decrypt.
func DecryptDataByKms(c context.Context, api KMSDecryptAPI, input *kms.DecryptInput) (*kms.DecryptOutput, error) {
	return api.Decrypt(c, input)
}

// DecryptData decrypts some text that was encrypted with an AWS Key Management Service (AWS KMS) key (KMS key).
// Inputs:
//     e is encrypted data by AWS KMS
// Output:
//     If success, a DecryptOutput object containing the result of the service call and "".
//     Otherwise, "" and an error from the call to Decrypt.
func DecryptData(e string) (string, error) {

	if strings.IsEmptyString(e) {
		return "", nil
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", fmt.Errorf("aws kms configuration error, %s", err)
	}
	client := kms.NewFromConfig(cfg)

	blob, err := b64.StdEncoding.DecodeString(e)
	if err != nil {
		return "", fmt.Errorf("aws kms error converting string to blob, %s", err)
	}

	input := &kms.DecryptInput{
		CiphertextBlob: blob,
	}

	result, err := DecryptDataByKms(context.TODO(), client, input)
	if err != nil {
		return "", fmt.Errorf("aws kms decrypting error, %s", err)
	}

	return string(result.Plaintext), nil
}
