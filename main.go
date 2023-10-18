package main

import (
	"fmt"
	"path/filepath"
)

var (
	youtubeMusicSearchUrl = "https://music.youtube.com/search?q="
	CREDENTAIL_FILE       = filepath.Join(".spotify-credentials")
)

func main() {
	client_id, client_secret, err := readCredentials()

	if err != nil {
		askForCredentials()
		client_id, client_secret, err = readCredentials()
		maybePanic(err)
	}

	fmt.Println(spotifyAuth(client_id, client_secret))
}
