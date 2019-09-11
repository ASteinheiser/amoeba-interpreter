package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ASteinheiser/amoeba-interpreter/color"
	"github.com/ASteinheiser/amoeba-interpreter/lexer"
	"github.com/ASteinheiser/amoeba-interpreter/token"
)

// ShowPrompt represents the symbols in the amoeba REPL
// directly before where the user types
func ShowPrompt() {
	color.ChangeColor(color.Magenta, false, color.Black, false)
	fmt.Print(`¯\_(ツ)_/¯`)
	color.ChangeColor(color.White, true, color.Black, false)
	fmt.Print(" >>>> ")
	color.ResetColor()
}

// Start will start a new amoeba REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		ShowPrompt()
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
