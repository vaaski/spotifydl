package main

import (
	"bufio"
	b64 "encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
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
	fmt.Print("client_id: ")
	reader := bufio.NewReader(os.Stdin)
	bClient_id, err := reader.ReadString('\n')
	maybePanic(err)
	client_id := strings.TrimSpace(string(bClient_id))

	fmt.Print("client_secret: ")
	bClient_secret, err := reader.ReadString('\n')
	maybePanic(err)
	client_secret := strings.TrimSpace(string(bClient_secret))

	joined := strings.Join([]string{client_id, client_secret}, ":")
	encoded := b64.StdEncoding.EncodeToString([]byte(joined))
	writeErr := os.WriteFile(CREDENTAIL_FILE, []byte(encoded), 0644)
	maybePanic(writeErr)

	log.Println("credentials saved at", CREDENTAIL_FILE)
}