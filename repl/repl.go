package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ASteinheiser/amoeba-interpreter/lexer"
	"github.com/ASteinheiser/amoeba-interpreter/token"
)

// PROMPT represents the symbols in the amoeba REPL
// directly before where the user types
const PROMPT = `¯\_(ツ)_/¯ >>>> `

// Start will start a new amoeba REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
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
