# Credentail.json 

- 실제 KMS 가 설정된 프로젝트로 지정하여 해당 코드를 동작시켜야 정상적으로 동작함. 

# keyring 의 목록을 조회해 오는 코드

```go
package main

import (
	"context"
	"fmt"
	"log"

	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"google.golang.org/api/iterator"
)

func main() {
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
```