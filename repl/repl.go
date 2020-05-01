package repl

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os/user"
	"strings"

	"github.com/ASteinheiser/amoeba-interpreter/color"
	"github.com/ASteinheiser/amoeba-interpreter/evaluator"
	"github.com/ASteinheiser/amoeba-interpreter/lexer"
	"github.com/ASteinheiser/amoeba-interpreter/object"
	"github.com/ASteinheiser/amoeba-interpreter/parser"
)

// Start will start a new amoeba REPL
func Start(in io.Reader, out io.Writer) {
	env := object.NewEnvironment()

	filePath := flag.String("file", "", "file path to read from")
	flag.Parse()

	if *filePath != "" {
		data, err := ioutil.ReadFile(*filePath)
		if err != nil {
			fmt.Println("File reading error:", err)
			return
		}

		evaluateProgram(string(data), out, env)
	} else {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}

		showWelcomeMessage(user)

		scanner := bufio.NewScanner(in)

		for {
			ShowPrompt()
			scanned := scanner.Scan()
			if !scanned {
				return
			}

			line := scanner.Text()
			if line == "exit" || line == "quit" {
				return
			}

			evaluateProgram(line, out, env)
		}
	}
}

func evaluateProgram(input string, out io.Writer, env *object.Environment) {
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		return
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		io.WriteString(out, "\n")
		io.WriteString(out, evaluated.Inspect())
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
	io.WriteString(out, "\n  Oops! Looks like your syntax has an issue...\n")
	io.WriteString(out, "    parser errors:\n\n")

	color.Foreground(color.Red, false)
	for _, msg := range errors {
		io.WriteString(out, "      "+msg+"\n\n")
	}
	color.ResetColor()
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
	fmt.Print("\n")
}
