package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func maybePanic(e error) {
	if e != nil {
		panic(e)
	}
}

func askForUserInput(prompt string) string {
	fmt.Print(prompt + ": ")
	reader := bufio.NewReader(os.Stdin)
	bClient_id, err := reader.ReadString('\n')
	maybePanic(err)

	return strings.TrimSpace(string(bClient_id))
}

// gets the folder containing the executable
func exeRoot() string {
	executablePath, _ := os.Executable()
	executableFolder := path.Join(executablePath, "..")

	if strings.HasPrefix(executableFolder, "/var/folders") {
		// the path for the executable is in some temp folder when using `go run .`
		// so we use the current working directory instead
		cwd, _ := os.Getwd()
		return cwd
	} else {
		return executableFolder
	}
}
