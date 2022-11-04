package gcp_key_projectsample

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"

	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/core/registry"
	"github.com/google/tink/go/keyset"
	"github.com/google/tink/go/testing/fakekms"
)

func Test_AeadEnvelopTest(t *testing.T) {
	// Initialization
	fakeKmsClient, err := fakekms.NewClient("fake-kms://")
	if err != nil {
		t.Fatalf("fakekms.NewClient('fake-kms://') failed: %v", err)
	}
	registry.RegisterKMSClient(fakeKmsClient)

	// KMS Setting - 해당 부분의 경우 별도 DB에 저장해두고 사용 가능함. DB에 저장해도 좋음. - Cache 를 적용해야할까?
	fixedKeyURI := "fake-kms://CM2b3_MDElQKSAowdHlwZS5nb29nbGVhcGlzLmNvbS9nb29nbGUuY3J5cHRvLnRpbmsuQWVzR2NtS2V5EhIaEIK75t5L-adlUwVhWvRuWUwYARABGM2b3_MDIAE"

	template1 := aead.KMSEnvelopeAEADKeyTemplate(fixedKeyURI, aead.AES128GCMKeyTemplate())

	// Key Handler Initialization
	handle1, err := keyset.NewHandle(template1)
	if err != nil {
		t.Fatalf("keyset.NewHandle(template1) failed: %v", err)
	}
	aead1, err := aead.New(handle1)
	if err != nil {
		t.Fatalf("aead.New(handle) failed: %v", err)
	}

	for i := 0; i < 500; i++ {
		// Data and Key
		plaintext := []byte("some data to encrypt" + strconv.Itoa(i))
		aad := []byte("extra data to authenticate")

		// Encrypt
		ciphertext, err := aead1.Encrypt(plaintext, aad)
		if err != nil {
			t.Fatalf("encryption failed, error: %v", err)
		}

		// Decrypt
		decrypted, err := aead1.Decrypt(ciphertext, aad)
		if err != nil {
			t.Fatalf("decryption failed, error: %v", err)
		}
		if !bytes.Equal(plaintext, decrypted) {
			t.Fatalf("decrypted data doesn't match plaintext, got: %q, want: %q", decrypted, plaintext)
		}

		fmt.Printf("Cipher Data: %s \n", ciphertext)
		fmt.Printf("Decrypt Data: %s \n", decrypted)
	}

}
