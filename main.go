package main

import (
	"fmt"
	"github.com/petrvelicka/amo/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s (%s)\n", user.Name, user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
