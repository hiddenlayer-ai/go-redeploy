package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

// make these accessible with a config file
const (
	WAITING    = 5
	REPOPATH   = "../"
	EXECUTABLE = "example"
)

var (
	DIFF  = "git diff origin/master"
	RESET = "git reset --hard HEAD"
	CLONE = "git pull --force"
	BUILD = "go build"
	EXEC  = fmt.Sprintf("./%s", EXECUTABLE)
)

func RUN(command string) *bytes.Buffer {
	cmd := exec.Command(command)
	cmd.Dir = REPOPATH
	out := new(bytes.Buffer)
	cmd.Stdout = out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return out
}

func main() {
	for {
		out := RUN(DIFF)

		// we found something
		if len(strings.Split(out.String(), " ")) != 0 {
			RUN(RESET)
			RUN(CLONE)
			RUN(BUILD)
			RUN(EXEC)
		}

		time.Sleep(WAITING)
	}
}
