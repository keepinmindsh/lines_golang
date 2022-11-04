package gcp_key_projectsample

import (
	"fmt"
	"log"
	"testing"

	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/keyset"
)

func Test_AeadSample(t *testing.T) {

	kh, err := keyset.NewHandle(aead.AES256GCMKeyTemplate())
	if err != nil {
		log.Fatal(err)
	}

	a, err := aead.New(kh)
	if err != nil {
		log.Fatal(err)
	}

	ct, err := a.Encrypt([]byte("this data needs to be encrypted"), []byte("associated data"))
	if err != nil {
		log.Fatal(err)
	}

	pt, err := a.Decrypt(ct, []byte("associated data"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Cipher text: %s\nPlain text: %s\n", ct, pt)
}
