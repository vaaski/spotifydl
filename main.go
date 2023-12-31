package main

import (
	"fmt"
)

// todo: compare youtube title to spotify title to make sure they match

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

	for index, track := range tracks {
		fmt.Println("downloading", track, "(", index+1, "/", len(tracks), ")")
		downloadTrack(track, playlistName)

		println("\n")
	}
}
