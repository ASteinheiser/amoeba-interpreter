package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ASteinheiser/amoeba-interpreter/color"
	"github.com/ASteinheiser/amoeba-interpreter/lexer"
	"github.com/ASteinheiser/amoeba-interpreter/parser"
)

// AMOEBA is an ascii amoeba string
const AMOEBA = ``

// ShowPrompt represents the symbols in the amoeba REPL
// directly before where the user types
func ShowPrompt() {
	color.ChangeColor(color.Magenta, false, color.Black, false)
	fmt.Print(` ¯\_(ツ)_/¯`)
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
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, "\n  ")
		io.WriteString(out, program.String())
		io.WriteString(out, "\n\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "\n")
	io.WriteString(out, AMOEBA)
	io.WriteString(out, "\n")
	io.WriteString(out, "  Oops! Looks like your syntax got infected...\n")
	io.WriteString(out, "    parser errors:\n\n")

	color.Foreground(color.Red, false)
	for _, msg := range errors {
		io.WriteString(out, "      "+msg+"\n\n")
	}
	color.ResetColor()
}
