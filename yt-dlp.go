package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var (
	youtubeMusicSearchUrl     = "https://music.youtube.com/search?q="
	youtubeMusicSearchPostfix = "#songs"
	ytDlpPath                 = "yt-dlp"
	args                      = []string{
		"-x",
		"--audio-format",
		"mp3",
		"-P",
		"ytdl-download",
		"-I",
		"1", // only download the first result
		"--extractor-args",
		"youtube:player_client=web;player_skip=configs", // makes it faster
	}
)

func downloadTrack(track string) {
	downloadArgs := args

	cpuCount := runtime.NumCPU()
	downloadArgs = append(downloadArgs, "-N", fmt.Sprint(cpuCount))
	downloadArgs = append(downloadArgs, "-o", track) // set the output file name since we know it
	downloadArgs = append(downloadArgs, youtubeMusicSearchUrl+track+youtubeMusicSearchPostfix)

	downloadCmd := exec.Command(ytDlpPath, downloadArgs...)
	downloadCmd.Stdout = os.Stdout
	downloadCmd.Stderr = os.Stderr

	downloadCmd.Run()
}
