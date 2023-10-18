package main

import (
	"bufio"
	"fmt"
	"os"
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
