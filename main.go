package main

import (
	"os"

	"github.com/ASteinheiser/amoeba-interpreter/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
