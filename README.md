# amoeba-interpreter
An interpreter for the amoeba programming language, written from scratch in Go.

## Features
- C-like syntax
- variables (integers, booleans, strings, arrays, objects)
- arithmetic expressions
- first-class and higher-order functions
- closures

## Run the Amoeba REPL
1. Install [Go](https://golang.org/dl/)
1. `git clone https://github.com/ASteinheiser/amoeba-interpreter.git`
1. `cd amoeba-interpreter`
1. `go run main.go`
```

    Hello ANDREW, Welcome to the Amoeba REPL!

    You can like... type code and stuff...


 ¯\_(ツ)_/¯ >>>> let x = 7;
 ¯\_(ツ)_/¯ >>>> x == 8;

false

 ¯\_(ツ)_/¯ >>>> █
```

## Run the test suite
You can run the tests for a sub-module individually as long as it has a `*_test.go` file:
```
go test ./ast/
go test ./lexer/
go test ./parser/
```
**OR** you can run all the tests at once:
```
./run-tests.sh
```
```

    Running the Amoeba test suite!

              ¯\_(ツ)_/¯

AST Test Results:
ok  	github.com/ASteinheiser/amoeba-interpreter/ast      0.005s

Lexer Test Results:
ok  	github.com/ASteinheiser/amoeba-interpreter/lexer    0.005s

Parser Test Results:
ok  	github.com/ASteinheiser/amoeba-interpreter/parser   0.005s

```
