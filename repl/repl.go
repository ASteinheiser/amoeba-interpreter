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
	io.WriteString(out, "                     ,,emmw,\n")
	io.WriteString(out, "                 a#M7|``  `^\"#p\n")
	io.WriteString(out, "               #M`            #p\n")
	io.WriteString(out, "              #b              @#\n")
	io.WriteString(out, "              #b              ]#\n")
	io.WriteString(out, "              \"#        ####m  @Q           ,e###M##m\n")
	io.WriteString(out, "               \"#      #######  7W#Mmm###M57^`      '@b\n")
	io.WriteString(out, "                \"@N    @######b               a#Mm    #p\n")
	io.WriteString(out, "                  7@w   @#####`              #b  #~   @#\n")
	io.WriteString(out, "                    %#   \"%##\"               \"%WM\\    @#\n")
	io.WriteString(out, "                     @b                             ,##\n")
	io.WriteString(out, "                     @#                   a###W%55577`\n")
	io.WriteString(out, "                    ,#=                 ##|\n")
	io.WriteString(out, "                   ;#\"    a#Mm         #b\n")
	io.WriteString(out, "                  ##p    #b  #~       #M\n")
	io.WriteString(out, "                  #b     \"%WM\\       #b\n")
	io.WriteString(out, "                  @#              ,##\"\n")
	io.WriteString(out, "                   %#p         ,##C,\n")
	io.WriteString(out, "                     7%@#####WT^`\n")
	io.WriteString(out, "\n")
	color.ResetColor()
}
