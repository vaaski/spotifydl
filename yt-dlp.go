package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
)

var (
	// append #songs to search only the songs and include the "Top result"
	// https://github.com/yt-dlp/yt-dlp/issues/6007#issuecomment-1769137538
	youtubeMusicSearchPostfix = "#songs"
	youtubeMusicSearchUrl     = "https://music.youtube.com/search?q="
	ytDlpPath                 = "yt-dlp"
	downloadFolder            = "ytdl-download"
	args                      = []string{
		"-x",
		"--audio-format",
		"mp3",
		"-I",
		"1", // only download the first result
		"--extractor-args",
		"youtube:player_client=web;player_skip=configs", // makes it faster
	}
)

func downloadTrack(track string, playlistName string) {
	downloadArgs := args

	cpuCount := runtime.NumCPU()
	downloadArgs = append(downloadArgs, "-N", fmt.Sprint(cpuCount))
	downloadArgs = append(downloadArgs, "-P", path.Join(downloadFolder, playlistName))
	downloadArgs = append(downloadArgs, "-o", track) // set the output file name since we know it
	downloadArgs = append(downloadArgs, youtubeMusicSearchUrl+track+youtubeMusicSearchPostfix)

	downloadCmd := exec.Command(ytDlpPath, downloadArgs...)
	downloadCmd.Stdout = os.Stdout
	downloadCmd.Stderr = os.Stderr

	downloadCmd.Run()
}
