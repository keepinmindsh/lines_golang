package samples

import (
	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"context"
	"fmt"
	fieldmask "google.golang.org/genproto/protobuf/field_mask"
)

// EnableKeyVersion disables the specified key version on Cloud KMS.
func EnableKeyVersion(name string) error {
	// name := "projects/my-project/locations/us-east1/keyRings/my-key-ring/cryptoKeys/my-key/cryptoKeyVersions/123"

	// Create the client.
	ctx := context.Background()
	client, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create kms client: %v", err)
	}
	defer client.Close()

	// Build the request.
	req := &kmspb.UpdateCryptoKeyVersionRequest{
		CryptoKeyVersion: &kmspb.CryptoKeyVersion{
			Name:  name,
			State: kmspb.CryptoKeyVersion_ENABLED,
		},
		UpdateMask: &fieldmask.FieldMask{
			Paths: []string{"state"},
		},
	}

	// Call the API.
	result, err := client.UpdateCryptoKeyVersion(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to update key version: %v", err)
	}
	fmt.Printf("Enabled key version: %s\n", result)
	return nil
}
