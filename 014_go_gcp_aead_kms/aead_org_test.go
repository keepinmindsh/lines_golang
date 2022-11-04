package gcp_key_projectsample

import (
	"encoding/base64"
	"fmt"
	"log"
	"testing"

	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/core/registry"
	"github.com/google/tink/go/keyset"
	"github.com/google/tink/go/testing/fakekms"
)

const (
	credentialsPath = "credentials.json"
)

func Test_OrginalAEADTest(t *testing.T) {

	// Initialization
	fakeKmsClient, err := fakekms.NewClient("fake-kms://")
	if err != nil {
		t.Fatalf("fakekms.NewClient('fake-kms://') failed: %v", err)
	}
	registry.RegisterKMSClient(fakeKmsClient)

	// KMS Setting - 해당 부분의 경우 별도 DB에 저장해두고 사용 가능함. DB에 저장해도 좋음. - Cache 를 적용해야할까?
	fixedKeyURI := "fake-kms://CM2b3_MDElQKSAowdHlwZS5nb29nbGVhcGlzLmNvbS9nb29nbGUuY3J5cHRvLnRpbmsuQWVzR2NtS2V5EhIaEIK75t5L-adlUwVhWvRuWUwYARABGM2b3_MDIAE"
	kh, err := keyset.NewHandle(aead.KMSEnvelopeAEADKeyTemplate(fixedKeyURI, aead.AES128GCMKeyTemplate()))
	if err != nil {
		log.Fatal(err)
	}

	a, err := aead.New(kh)
	if err != nil {
		log.Fatal(err)
	}

	// An io.Reader and io.Writer implementation which simply writes to memory.
	memKeyset := &keyset.MemReaderWriter{}

	// Write encrypts the keyset handle with the master key and writes to the
	// io.Writer implementation (memKeyset). We recommend that you encrypt the
	// keyset handle before persisting it.
	if err := kh.Write(memKeyset, a); err != nil {
		log.Fatal(err)
	}

	// Read reads the encrypted keyset handle back from the io.Reader
	// implementation and decrypts it using the master key.
	kh2, err := keyset.Read(memKeyset, a)
	if err != nil {
		log.Fatal(err)
	}

	a2, err := aead.New(kh2)

	msg := []byte("this message needs to be encrypted")
	aad := []byte("this data needs to be authenticated, but not encrypted")
	ct, err := a2.Encrypt(msg, aad)
	if err != nil {
		log.Fatal(err)
	}

	pt, err := a2.Decrypt(ct, aad)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Ciphertext: %s\n", base64.StdEncoding.EncodeToString(ct))
	fmt.Printf("Original  plaintext: %s\n", msg)
	fmt.Printf("Decrypted Plaintext: %s\n", pt)
}
