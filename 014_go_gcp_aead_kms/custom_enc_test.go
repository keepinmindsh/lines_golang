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

func Test_KMSProcess(t *testing.T) {
	// Generate a new key.
	kh1, err := keyset.NewHandle(aead.AES128GCMKeyTemplate())
	if err != nil {
		log.Fatal(err)
	}

	// Fetch the master key from a KMS.
	gcpClient, err := gcpkms.NewClientWithOptions(context.Background(), "gcp-kms://")
	if err != nil {
		log.Fatal(err)
	}
	registry.RegisterKMSClient(gcpClient)

	masterKey, err := gcpClient.GetAEAD(keyURI)
	if err != nil {
		log.Fatal(err)
	}

	// An io.Reader and io.Writer implementation which simply writes to memory.
	memKeySet := &keyset.MemReaderWriter{}

	// Write encrypts the keyset handle with the master key and writes to the
	// io.Writer implementation (memKeySet). We recommend that you encrypt the
	// keyset handle before persisting it.
	if err := kh1.Write(memKeySet, masterKey); err != nil {
		log.Fatal(err)
	}

	// Read reads the encrypted keyset handle back from the io.Reader
	// implementation and decrypts it using the master key.
	kh2, err := keyset.Read(memKeySet, masterKey)
	if err != nil {
		log.Fatal(err)
	}

	t2, err := aead.New(kh2)

	if err != nil {
		log.Fatal(err)
	}

	msg := []byte("this message needs to be encrypted")
	aad := []byte("this data needs to be authenticated, but not encrypted")
	ct, err := t2.Encrypt(msg, aad)
	if err != nil {
		log.Fatal(err)
	}

	pt, err := t2.Decrypt(ct, aad)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Ciphertext: %s\n", base64.StdEncoding.EncodeToString(ct))
	fmt.Printf("Original  plaintext: %s\n", msg)
	fmt.Printf("Decrypted Plaintext: %s\n", pt)
}
