package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
)

type Config struct {
	SLEEP      uint   `toml:"sleep"`
	REPOPATH   string `toml:"repo-path"`
	ENTRYPOINT string `toml:"entry-point"`
}

var (
	DIFF  = "git diff origin/master"
	RESET = "git reset --hard HEAD"
	CLONE = "git pull --force"
	BUILD = "go build"
	EXEC  = fmt.Sprintf("go run %s.go", conf.ENTRYPOINT)
)

var conf Config

func RUN(command string) *bytes.Buffer {
	c := strings.Split(command, " ")
	cmd := exec.Command(c[0], c[1:]...)
	cmd.Dir = conf.REPOPATH
	out := new(bytes.Buffer)
	cmd.Stdout = out

	if err := cmd.Run(); err != nil {
		log.Print(err)
		return nil
	}

	return out
}

func main() {
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		log.Fatal(err)
	}

	for {
		out := RUN(DIFF)

		if len(strings.Split(out.String(), " ")) != 0 {
			// new update
			RUN(RESET)
			RUN(CLONE)
			RUN(BUILD)
			RUN(EXEC)
		}

		time.Sleep(time.Duration(conf.SLEEP) * time.Second)
	}
}
