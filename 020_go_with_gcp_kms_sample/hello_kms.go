package main

import (
	"020_go_with_gcp_kms_sample/samples"
	"fmt"
	"log"
	"strings"
)

const (
	AlreadyExist = "AlreadyExists"
)

func main() {
	projectID := "lines-infra"
	locationID := "global"
	keyRingId := "lines_keyring"
	keyId := "lines_alarm_key"

	uriWithKeyring := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", projectID, locationID, keyRingId)
	_ = fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", projectID, locationID, keyRingId, keyId)
	// name := "projects/my-project/locations/us-east1/keyRings/my-key-ring/cryptoKeys/my-key/cryptoKeyVersions/123"x
	uriWithKeyVersion := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/1", projectID, locationID, keyRingId, keyId)

	samples.ListOfKeyRingsTest()

	// CAVIUM_V2_COMPRESSED
	keyVersionErr := samples.GetKeyVersionAttestation(uriWithKeyVersion)
	if keyVersionErr != nil {
		log.Fatal(keyVersionErr.Error())
	}

	err := samples.CreateKeyHSM(log.Writer(), uriWithKeyring, keyId)

	if err != nil {
		if strings.Index(err.Error(), AlreadyExist) > -1 {
			log.Println("Key already exists")
		} else {
			log.Fatal(err.Error())
		}
	}
}
