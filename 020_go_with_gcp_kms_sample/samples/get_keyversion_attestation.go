package samples

import (
	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"context"
	"fmt"
)

// GetKeyVersionAttestation gets the attestation on a key version, if one
// exists.
func GetKeyVersionAttestation(name string) error {
	// name := "projects/my-project/locations/us-east1/keyRings/my-key-ring/cryptoKeys/my-key/cryptoKeyVersions/123"x

	// Create the client.
	ctx := context.Background()
	client, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create kms client: %v", err)
	}

	// Build the request.
	req := &kmspb.GetCryptoKeyVersionRequest{
		Name: name,
	}

	// Call the API.
	result, err := client.GetCryptoKeyVersion(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to get key: %v", err)
	}

	// Only HSM keys have an attestation. For other key types, the attestion will
	// be nil.
	attestation := result.Attestation
	if attestation == nil {
		return fmt.Errorf("no attestation for %s", name)
	}

	// Print the attestation, hex-encoded.
	fmt.Printf("%s: %x", attestation.Format, attestation.Content)
	return nil
}
