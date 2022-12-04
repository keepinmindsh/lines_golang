package samples

import (
	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"context"
	"fmt"
)

// destroyKeyVersion marks a specified key version for deletion. The key can be
// restored if requested within 24 hours.
func DestroyKeyVersion(name string) error {
	// name := "projects/my-project/locations/us-east1/keyRings/my-key-ring/cryptoKeys/my-key/cryptoKeyVersions/123"

	// Create the client.
	ctx := context.Background()
	client, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create kms client: %v", err)
	}
	defer client.Close()

	// Build the request.
	req := &kmspb.DestroyCryptoKeyVersionRequest{
		Name: name,
	}

	// Call the API.
	result, err := client.DestroyCryptoKeyVersion(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to destroy key version: %v", err)
	}
	fmt.Printf("Destroyed key version: %s\n", result)
	return nil
}
