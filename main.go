package main

import (
	"monkey/repl"
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
