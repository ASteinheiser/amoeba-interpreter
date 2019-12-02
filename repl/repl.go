package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ASteinheiser/amoeba-interpreter/color"
	"github.com/ASteinheiser/amoeba-interpreter/lexer"
	"github.com/ASteinheiser/amoeba-interpreter/parser"
)

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

// ShowPrompt prints out the symbols in the amoeba REPL
// directly before where the user types
func ShowPrompt() {
	color.ChangeColor(color.Magenta, false, color.Black, false)
	fmt.Print(` ¯\_(ツ)_/¯`)
	color.ChangeColor(color.White, true, color.Black, false)
	fmt.Print(" >>>> ")
	color.ResetColor()
}

func printParserErrors(out io.Writer, errors []string) {
	ShowAmoeba(out)
	io.WriteString(out, "  Oops! Looks like your syntax got infected...\n")
	io.WriteString(out, "    parser errors:\n\n")

	color.Foreground(color.Red, false)
	for _, msg := range errors {
		io.WriteString(out, "      "+msg+"\n\n")
	}
	color.ResetColor()
}

// ShowAmoeba prints the ASCII amoeba
func ShowAmoeba(out io.Writer) {
	color.Foreground(color.Green, false)
	io.WriteString(out, "\n")
	io.WriteString(out, "             ,,,,g,\n")
	io.WriteString(out, "           #\"`    `@\n")
	io.WriteString(out, "          @        \\b\n")
	io.WriteString(out, "          jb    ##m @      ,smWWm\n")
	io.WriteString(out, "           7m  ]#### '`7^\"\"      @\n")
	io.WriteString(out, "             %p 7##b      #j@     b\n")
	io.WriteString(out, "              @            ,,,,,,M`\n")
	io.WriteString(out, "              @    ,w    ,M|'\n")
	io.WriteString(out, "            ,#`   7m#`  ]b\n")
	io.WriteString(out, "            @b         {^\n")
	io.WriteString(out, "             %m     a#/\n")
	io.WriteString(out, "               ^\"\"`^\n")
	io.WriteString(out, "\n")
	color.ResetColor()
}
