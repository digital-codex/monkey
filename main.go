package main

import (
	"github.com/digital-codex/monkey/repl"
	"os"
	"os/user"
)

func main() {
	current, err := user.Current()
	if err != nil {
		panic(err)
	}
	repl.Start(os.Stdin, os.Stdout, current)
}
