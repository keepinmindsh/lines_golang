package main

import (
	kms "cloud.google.com/go/kms/apiv1"
	"context"
)

func main() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := kms.NewEkmClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()
}
