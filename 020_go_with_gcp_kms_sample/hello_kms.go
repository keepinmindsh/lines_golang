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
	keyValidationErr := samples.GetKeyVersionAttestation(uriWithKeyVersion)
	if keyValidationErr != nil {
		log.Println(" ")
		log.Fatal(keyValidationErr.Error())
	}

	// 삭제는 Scheduled에 의해서 동작함.
	// failed to update key version: rpc error: code = FailedPrecondition desc = The request cannot be fulfilled. Resource projects/lines-infra/locations/global/keyRings/lines_keyring/cryptoKeys/lines_alarm_key/cryptoKeyVersions/1 has value DESTROY_SCHEDULED in field crypto_key_version.state.
	keyDestroyErr := samples.DestroyKeyVersion(uriWithKeyVersion)
	if keyDestroyErr != nil {
		log.Println(" ")
		log.Fatal(keyDestroyErr)
	}

	err := samples.CreateKeyHSM(log.Writer(), uriWithKeyring, keyId)

	if err != nil {
		if strings.Index(err.Error(), AlreadyExist) > -1 {
			log.Println("Key already exists")
		} else {
			log.Println(" ")
			log.Fatal(err.Error())
		}
	}

	// 이미 삭제 스케쥴링이 되고 다면 활성화 불가한 것으로 판단됨.
	keyDisableErr := samples.DisableKeyVersion(uriWithKeyVersion)
	if keyDisableErr != nil {
		log.Fatal(keyDisableErr)
	}
}
