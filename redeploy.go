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
	WAITING    = 5                // time in seconds to wait before update lookup
	REPOPATH   = "../go-redeploy" // relative path to repo to manage
	EXECUTABLE = "go-redeploy"    // name of executable to run
)

var (
	DIFF  = "git diff origin/master"
	RESET = "git reset --hard HEAD"
	CLONE = "git pull --force"
	BUILD = "go build"
	EXEC  = fmt.Sprintf("./%s", EXECUTABLE)
)

func RUN(command string) *bytes.Buffer {
	c := strings.Split(command, " ")
	cmd := exec.Command(c[0], c[1:]...)
	cmd.Dir = REPOPATH
	out := new(bytes.Buffer)
	cmd.Stdout = out

	err := cmd.Run()

	if err != nil {
		log.Print(err)
		return nil
	}

	return out
}

func main() {
	fmt.Println("ayyy")
	for {
		out := RUN(DIFF)

		// we found something
		if len(strings.Split(out.String(), " ")) != 0 {
			RUN(RESET)
			RUN(CLONE)
			RUN(BUILD)
			RUN(EXEC)
		}

		time.Sleep(WAITING * time.Second)
	}
}
