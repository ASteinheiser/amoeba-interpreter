# amoeba-interpreter
An interpreter for the amoeba programming language, written from scratch in Go.

## Features
- C-like syntax
- variables (integers, booleans, strings, arrays, objects)
- arithmetic expressions
- first-class and higher-order functions
- closures

## Run the Amoeba REPL
1. `git clone https://github.com/ASteinheiser/amoeba-interpreter.git`
1. `cd amoeba-interpreter`
1. `./amoeba-interpreter`
```

    Hello ANDREW, Welcome to the Amoeba REPL!

    You can like... type code and stuff...


 ¯\_(ツ)_/¯ >>>> let x = 7;

  x = 7

 ¯\_(ツ)_/¯ >>>> x == 8;

  false

 ¯\_(ツ)_/¯ >>>> let y 4

             ,,,,g,
           #"`    `@
          @        \b
          jb    ##m @      ,smWWm
           7m  ]#### '`7^""      @
             %p 7##b      #j@     b
              @            ,,,,,,M`
              @    ,w    ,M|'
            ,#`   7m#`  ]b
            @b         {^
             %m     a#/
               ^""`^

  Oops! Looks like your syntax got infected...
    parser errors:

      expected '4' to be =, got INT instead

 ¯\_(ツ)_/¯ >>>> █
```

## Local Dev
1. Install [Go](https://golang.org/dl/)
1. `git clone https://github.com/ASteinheiser/amoeba-interpreter.git`
1. `cd amoeba-interpreter`
1. `go run main.go`

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

## Backburner Features
- [ ] add <= and >= operators
- [ ] enhance error messages with line number and file name
- [ ] add postfix operators (such as `++`)

## Potential Long-Term Goals
- Make Amoeba Turing Complete?
- Rewrite the Amoeba interpreter in Amoeba?
