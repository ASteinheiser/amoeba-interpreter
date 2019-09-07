package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/ASteinheiser/amoeba-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, Welcome to the Amoeba REPL!\n\n",
		strings.ToUpper(user.Username))
	fmt.Printf("You can like... type code and stuff...\n\n")
	repl.Start(os.Stdin, os.Stdout)
}
