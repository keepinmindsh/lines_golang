package samples

import (
	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/durationpb"
	"time"
)

// CreateKeyHSM creates a new symmetric encrypt/decrypt key on Cloud KMS.
func CreateKeyHSM(parent, id string) error {
	// parent := "projects/my-project/locations/us-east1/keyRings/my-key-ring"
	// id := "my-hsm-encryption-key"

	// Create the client.
	ctx := context.Background()
	client, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create kms client: %v", err)
	}
	defer client.Close()

	// Build the request.
	req := &kmspb.CreateCryptoKeyRequest{
		Parent:      parent,
		CryptoKeyId: id,
		CryptoKey: &kmspb.CryptoKey{
			Purpose: kmspb.CryptoKey_ENCRYPT_DECRYPT,
			VersionTemplate: &kmspb.CryptoKeyVersionTemplate{
				ProtectionLevel: kmspb.ProtectionLevel_SOFTWARE,
				Algorithm:       kmspb.CryptoKeyVersion_GOOGLE_SYMMETRIC_ENCRYPTION,
			},

			// Optional: customize how long key versions should be kept before destroying.
			DestroyScheduledDuration: durationpb.New(24 * time.Hour),
		},
	}

	// Call the API.
	result, err := client.CreateCryptoKey(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to create key: %v", err)
	}
	fmt.Printf("Created key: %s\n", result.Name)
	return nil
}
