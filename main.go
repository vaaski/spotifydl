package main

import (
	"fmt"
	"path/filepath"
)

var (
	CREDENTAIL_FILE = filepath.Join(".spotify-credentials")
)

// append #songs to search only the songs and include the "Top result"
// https://github.com/yt-dlp/yt-dlp/issues/6007#issuecomment-1769137538

// yt-dlp -I 1 -x --audio-format mp3 --extractor-args 'youtube:player_client=web;player_skip=configs' "https://music.youtube.com/search?q=query#songs"

// short playlist   7fBWGZ99ymBeGXeIKWebyh
// long playlist    0T8npk4GpmL564lMzaynPd
// medium playlist  62KQaqwTfsOViSU49uUozv

// todo parse spotify playlist url

func main() {
	client_id, client_secret, err := readCredentials()

	if err != nil {
		askForCredentials()
		client_id, client_secret, err = readCredentials()
		maybePanic(err)
	}

	access_token, err := spotifyAuth(client_id, client_secret)
	maybePanic(err)

	playlist_id := askForUserInput("playlist_id")

	tracks, err := getTracks(playlist_id, access_token)
	maybePanic(err)

	for _, track := range tracks {
		fmt.Println("downloading", track)
		downloadTrack(track)

		println("\n")
	}
}
