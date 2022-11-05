package demo_strcuture

import (
	"fmt"

	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/core/registry"
	"github.com/google/tink/go/keyset"
	"github.com/google/tink/go/testing/fakekms"
	"github.com/google/tink/go/tink"
)

type KMSGcp struct {
}

func (kms *KMSGcp) Register() {
	// Initialization
	fakeKmsClient, err := fakekms.NewClient("fake-kms://")
	if err != nil {
		fmt.Errorf("fakekms.NewClient('fake-kms://') failed: %v", err)
	}
	registry.RegisterKMSClient(fakeKmsClient)
}

// returnAEAD
// keyURI := "gcp-kms://projects/" + {project} + "/locations/global/keyRings/" + {keyRings} + "/cryptoKeys/" + {encryptKey}
func returnAEAD(encryptKey string) tink.AEAD {
	// KMS Setting - 해당 부분의 경우 별도 DB에 저장해두고 사용 가능함. DB에 저장해도 좋음. - Cache 를 적용해야할까?
	fixedKeyURI := "fake-kms://CM2b3_MDElQKSAowdHlwZS5nb29nbGVhcGlzLmNvbS9nb29nbGUuY3J5cHRvLnRpbmsuQWVzR2NtS2V5EhIaEIK75t5L-adlUwVhWvRuWUwYARABGM2b3_MDIAE"

	envTemplate := aead.KMSEnvelopeAEADKeyTemplate(fixedKeyURI, aead.AES256GCMKeyTemplate())

	// Key Handler Initialization
	keyHandle, err := keyset.NewHandle(envTemplate)
	if err != nil {
		fmt.Errorf("keyset.NewHandle(envTemplate) failed: %v", err)
	}

	aeadKey, err := aead.New(keyHandle)
	if err != nil {
		fmt.Errorf("aeadKey.New(handle) failed: %v", err)
	}

	return aeadKey
}

func (kms *KMSGcp) Encrypt(masking *KMSValueT) *KMSValueR {
	encrypted, err := returnAEAD(masking.cmpId).Encrypt(masking.plaintext, masking.aead)
	if err != nil {
		return nil
	}
	return &KMSValueR{
		result: encrypted,
	}
}

func (kms *KMSGcp) Decrypt(masking *KMSValueT) *KMSValueR {
	retVal, err := returnAEAD(masking.cmpId).Decrypt(masking.plaintext, masking.aead)
	if err != nil {
		return nil
	}

	return &KMSValueR{
		result: retVal,
	}
}
