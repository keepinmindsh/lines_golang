package samples

import (
	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
)

// todo 키 버전을 가져오는 것이 아닌가?
// GetKeyVersion creates a new key version for the given key.
func GetKeyVersion(parent string) error {
	// parent := "projects/my-project/locations/us-east1/keyRings/my-key-ring/cryptoKeys/my-key"

	// Create the client.
	ctx := context.Background()
	client, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to get kms client: %v", err)
	}
	defer client.Close()

	// Build the request.
	req := &kmspb.ListCryptoKeyVersionsRequest{
		Parent: parent,
	}

	// Call the API.
	result := client.ListCryptoKeyVersions(ctx, req)

	if err != nil {
		return fmt.Errorf("failed to get key version: %v", err)
	}

	for {
		resp, err := result.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err.Error())
		}
		// todo 여기에서 이해가 안되는 부분은 Rotate 되는 키가 이전 키를 대체한다는 여부를 확정할 수 있는가? ...?
		log.Printf("%s, %s, %s, %s, %s, %s, %s \r\n", resp.Name, resp.Algorithm, resp.CreateTime, resp.GenerateTime, resp.State, resp.ProtectionLevel, resp.ExternalProtectionLevelOptions)
		_ = resp
	}
	return nil

}
