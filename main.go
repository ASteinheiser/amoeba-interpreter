package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/ASteinheiser/amoeba-interpreter/color"
	"github.com/ASteinheiser/amoeba-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	showWelcomeMessage(user)

	repl.Start(os.Stdin, os.Stdout)
}

func showWelcomeMessage(user *user.User) {
	color.ChangeColor(color.None, false, color.Black, false)
	fmt.Print("                                                 ")
	color.ResetColor()
	fmt.Print("\n")
	color.ChangeColor(color.Cyan, false, color.Black, false)
	fmt.Print("    Hello ")
	color.ChangeColor(color.Magenta, false, color.Black, false)
	fmt.Printf("%s", strings.ToUpper(user.Username))
	color.ChangeColor(color.Cyan, false, color.Black, false)
	fmt.Print(", Welcome to the Amoeba REPL!    ")
	color.ResetColor()
	fmt.Print("\n")
	color.ChangeColor(color.None, false, color.Black, false)
	fmt.Print("                                                 ")
	color.ResetColor()
	fmt.Print("\n")
	color.ChangeColor(color.Cyan, false, color.Black, false)
	fmt.Printf("    You can like... type code and stuff...       ")
	color.ResetColor()
	fmt.Print("\n")
	color.ChangeColor(color.None, false, color.Black, false)
	fmt.Print("                                                 ")
	color.ResetColor()
	fmt.Print("\n")
}
