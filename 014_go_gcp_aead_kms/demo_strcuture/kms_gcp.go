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
	a tink.AEAD
}

func (kms *KMSGcp) Register() {
	// Initialization
	fakeKmsClient, err := fakekms.NewClient("fake-kms://")
	if err != nil {
		fmt.Errorf("fakekms.NewClient('fake-kms://') failed: %v", err)
	}
	registry.RegisterKMSClient(fakeKmsClient)
}

// todo 해당 코드가 메모리 상에 떠있어야함!
func (kms *KMSGcp) Init() {
	// KMS Setting - 해당 부분의 경우 별도 DB에 저장해두고 사용 가능함. DB에 저장해도 좋음. - Cache 를 적용해야할까?
	fixedKeyURI := "fake-kms://CM2b3_MDElQKSAowdHlwZS5nb29nbGVhcGlzLmNvbS9nb29nbGUuY3J5cHRvLnRpbmsuQWVzR2NtS2V5EhIaEIK75t5L-adlUwVhWvRuWUwYARABGM2b3_MDIAE"

	template1 := aead.KMSEnvelopeAEADKeyTemplate(fixedKeyURI, aead.AES128GCMKeyTemplate())

	// Key Handler Initialization
	handle1, err := keyset.NewHandle(template1)
	if err != nil {
		fmt.Errorf("keyset.NewHandle(template1) failed: %v", err)
	}
	aead1, err := aead.New(handle1)
	if err != nil {
		fmt.Errorf("aead.New(handle) failed: %v", err)
	}

	kms.a = aead1
}

func (kms *KMSGcp) Encrypt(masking *KMSValueT) *KMSValueR {
	encrypted, err := kms.a.Encrypt(masking.plaintext, masking.aead)
	if err != nil {
		return nil
	}
	return &KMSValueR{
		result: encrypted,
	}
}

func (kms *KMSGcp) Decrypt(masking *KMSValueT) *KMSValueR {
	retVal, err := kms.a.Decrypt(masking.plaintext, masking.aead)
	if err != nil {
		return nil
	}

	return &KMSValueR{
		result: retVal,
	}
}
