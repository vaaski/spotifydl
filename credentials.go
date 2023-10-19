package main

import (
	b64 "encoding/base64"
	"log"
	"os"
	"path"
	"strings"
)

var (
	CREDENTAIL_FILE = path.Join(exeRoot(), ".spotify-credentials")
)

func readCredentials() (string, string, error) {
	data, err := os.ReadFile(CREDENTAIL_FILE)
	if err != nil {
		return "", "", err
	}

	bDecoded, decodeErr := b64.StdEncoding.DecodeString(string(data))
	if decodeErr != nil {
		return "", "", decodeErr
	}

	decoded := string(bDecoded)
	credentials := strings.Split(string(decoded), ":")
	return credentials[0], credentials[1], nil
}

func askForCredentials() {
	client_id := askForUserInput("api client_id")
	client_secret := askForUserInput("api client_secret")

	joined := strings.Join([]string{client_id, client_secret}, ":")
	encoded := b64.StdEncoding.EncodeToString([]byte(joined))
	writeErr := os.WriteFile(CREDENTAIL_FILE, []byte(encoded), 0644)
	maybePanic(writeErr)

	log.Println("credentials saved at", CREDENTAIL_FILE)
}
