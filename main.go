package main

import (
	"fmt"
	"path/filepath"
)

var (
	CREDENTAIL_FILE = filepath.Join(".spotify-credentials")
)

func main() {
	client_id, client_secret, err := readCredentials()

	if err != nil {
		askForCredentials()
		client_id, client_secret, err = readCredentials()
		maybePanic(err)
	}

	access_token, err := spotifyAuth(client_id, client_secret)
	maybePanic(err)

	playlist_id := parseSpotifyUrlOrId(askForUserInput("playlist id or url"))

	playlistName, err := getPlaylistName(playlist_id, access_token)
	maybePanic(err)

	tracks, err := getTracks(playlist_id, access_token)
	maybePanic(err)

	for _, track := range tracks {
		fmt.Println("downloading", track)
		downloadTrack(track, playlistName)

		println("\n")
	}
}
