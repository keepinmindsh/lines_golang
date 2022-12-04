package samples

import (
	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"context"
	"fmt"
	"io"
	"log"
)

// GetKeyLabel fetches the labels on a KMS key.
func GetKeyLabel(w io.Writer, name string) (bool, error) {
	// name := "projects/my-project/locations/us-east1/keyRings/my-key-ring/cryptoKeys/my-key"

	// Create the client.
	ctx := context.Background()
	client, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to create kms client: %v", err)
	}
	defer client.Close()

	// Build the request.
	req := &kmspb.GetCryptoKeyRequest{
		Name: name,
	}

	// Call the API.
	result, err := client.GetCryptoKey(ctx, req)
	if err != nil {
		return false, fmt.Errorf("failed to get key: %v", err)
	}

	// Extract and print the labels.
	for k, v := range result.Labels {
		value, err := fmt.Fprintf(w, "%s=%s\n", k, v)
		if err != nil {
			log.Fatal(err)
		}
		log.Print(value)
	}
	return true, nil
}
