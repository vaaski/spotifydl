package main

import (
	"io"
	"net/http"
	"strings"

	"github.com/buger/jsonparser"
)

func spotifyAuth(client_id string, client_secret string) string {
	url := "https://accounts.spotify.com/api/token"
	payload := strings.NewReader("grant_type=client_credentials&client_id=" + client_id + "&client_secret=" + client_secret)

	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	maybePanic(err)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	access_token, err := jsonparser.GetString(body, "access_token")
	maybePanic(err)

	return access_token
}
