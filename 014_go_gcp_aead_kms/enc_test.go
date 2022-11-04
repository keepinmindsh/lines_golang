package gcp_key_projectsample

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/core/registry"
	"github.com/google/tink/go/integration/gcpkms"
	"github.com/google/tink/go/keyset"
)

const (
	// Change this. AWS KMS, Google Cloud KMS and HashiCorp Vault are supported out of the box.
	keyURI = "gcp-kms://projects/tink-examples/locations/global/keyRings/foo/cryptoKeys/bar"
)

func Test_KMSEncryption(t *testing.T) {
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
	memKeyset := &keyset.MemReaderWriter{}

	// Write encrypts the keyset handle with the master key and writes to the
	// io.Writer implementation (memKeyset). We recommend that you encrypt the
	// keyset handle before persisting it.
	if err := kh1.Write(memKeyset, masterKey); err != nil {
		log.Fatal(err)
	}

	// Read reads the encrypted keyset handle back from the io.Reader
	// implementation and decrypts it using the master key.
	kh2, err := keyset.Read(memKeyset, masterKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(kh2)
}
