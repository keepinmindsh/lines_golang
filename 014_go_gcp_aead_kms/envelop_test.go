package gcp_key_projectsample

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"testing"

	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/core/registry"
	"github.com/google/tink/go/integration/gcpkms"
	"github.com/google/tink/go/keyset"
)

func Test_EnvelopEncryption(t *testing.T) {
	// Initialization
	gcpclient, err := gcpkms.NewClientWithOptions(context.Background(), "gcp-kms://")
	if err != nil {
		log.Fatal(err)
	}
	registry.RegisterKMSClient(gcpclient)

	// Decryption Key
	dek := aead.AES128CTRHMACSHA256KeyTemplate()
	kh, err := keyset.NewHandle(aead.KMSEnvelopeAEADKeyTemplate(keyURI, dek))
	if err != nil {
		log.Fatal(err)
	}

	// AEAD Creation
	a, err := aead.New(kh)
	if err != nil {
		log.Fatal(err)
	}

	// 암호화 대상 데이터
	msg := []byte("this message needs to be encrypted")

	// 암호화를 위한 키 정보
	aad := []byte("this data needs to be authenticated, but not encrypted")

	ct, err := a.Encrypt(msg, aad)
	if err != nil {
		log.Fatal(err)
	}

	pt, err := a.Decrypt(ct, aad)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Ciphertext: %s\n", base64.StdEncoding.EncodeToString(ct))
	fmt.Printf("Original  plaintext: %s\n", msg)
	fmt.Printf("Decrypted Plaintext: %s\n", pt)

}
