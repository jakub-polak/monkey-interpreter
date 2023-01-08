package main

import (
	"fmt"
	"os"
	"os/user"

	"monkey-interpreter/internal/repl"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", u.Username)
	repl.Start(os.Stdin, os.Stdout)
}
