package main

import (
	"020_go_with_gcp_kms_sample/samples"
	"fmt"
	"log"
)

func main() {
	projectID := "lines-infra"
	locationID := "global"
	keyRingId := "lines_keyring"
	keyId := "lines_alarm_key"

	uriWithKeyring := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", projectID, locationID, keyRingId)
	_ = fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", projectID, locationID, keyRingId, keyId)

	samples.ListOfKeyRingsTest()

	err := samples.CreateKeyHSM(log.Writer(), uriWithKeyring, keyId)
	if err != nil {
		log.Fatal(err.Error())
	}
}
