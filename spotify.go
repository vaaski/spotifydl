package main

import (
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/buger/jsonparser"
)

func spotifyAuth(client_id string, client_secret string) (string, error) {
	url := "https://accounts.spotify.com/api/token"
	payload := strings.NewReader("grant_type=client_credentials&client_id=" + client_id + "&client_secret=" + client_secret)

	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	access_token, err := jsonparser.GetString(body, "access_token")
	if err != nil {
		return "", err
	}

	return access_token, nil
}

func getPlaylistName(playlist_id string, access_token string) (string, error) {
	url := "https://api.spotify.com/v1/playlists/" + playlist_id + "?fields=name"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+access_token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	name, err := jsonparser.GetString(body, "name")
	if err != nil {
		return "", err
	}

	return name, nil
}

func getTracks(playlist_id string, access_token string) ([]string, error) {
	url := "https://api.spotify.com/v1/playlists/" + playlist_id + "/tracks?fields=next,items(track(name,artists(name)))&limit=50"

	getPage := func(url string) ([]byte, error) {
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("Authorization", "Bearer "+access_token)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	}

	var accumulated [][]byte
	body, err := getPage(url)
	if err != nil {
		return nil, err
	}
	accumulated = append(accumulated, body)
	next, nextErr := jsonparser.GetString(body, "next")

	for nextErr == nil && next != "" {
		body, err := getPage(next)
		if err != nil {
			return nil, err
		}
		accumulated = append(accumulated, body)
		next, nextErr = jsonparser.GetString(body, "next")
	}

	return parseTracks(accumulated), nil
}

// takes the body of the response from the spotify api
// and returns a slice of strings containing the formatted tracks
func parseTracks(bodies [][]byte) []string {
	var tracks []string

	for _, page := range bodies {
		jsonparser.ArrayEach(page, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			track, err := jsonparser.GetString(value, "track", "name")
			maybePanic(err)

			var artists []string
			jsonparser.ArrayEach(value, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				artist, err := jsonparser.GetString(value, "name")
				maybePanic(err)

				artists = append(artists, artist)
			}, "track", "artists")

			tracks = append(tracks, strings.Join(artists, ", ")+" - "+track)
		}, "items")
	}

	return tracks
}

func parseSpotifyUrlOrId(urlOrId string) string {
	if len(urlOrId) == 22 {
		return urlOrId
	}

	regex := regexp.MustCompile(`spotify\.com\/playlist\/(\w{22})`)
	matches := regex.FindStringSubmatch(urlOrId)
	if len(matches) == 2 {
		return matches[1]
	}

	panic("invalid spotify url or id")
}
