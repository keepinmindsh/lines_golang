package samples

import (
	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
)

func ListOfKeyRingsTest() {
	// GCP project with which to communicate.
	projectID := "lines-infra"

	// Location in which to list key rings.
	locationID := "global"

	// Create the client.
	ctx := context.Background()
	client, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup client: %v", err)
	}
	defer client.Close()

	// Create the request to list KeyRings.
	listKeyRingsReq := &kmspb.ListKeyRingsRequest{
		Parent: fmt.Sprintf("projects/%s/locations/%s", projectID, locationID),
	}

	// List the KeyRings.
	it := client.ListKeyRings(ctx, listKeyRingsReq)

	// Iterate and print the results.
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to list key rings: %v", err)
		}

		fmt.Printf("key ring: %s\n", resp.Name)
	}
}
