# amoeba-interpreter
An interpreter for the Amoeba programming language. Written from scratch in Go with no dependencies 🎉

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
go test ./evaluator/
```
**OR** you can run all the tests at once:
```
./run-tests.sh
```
<img
  src="https://s3-us-west-2.amazonaws.com/images.iamandrew.io/Screen+Shot+2019-12-04+at+12.50.45+AM.png"
  width="600px"
  alt="Tests Screenshot"
/>

## Backburner Features
- [ ] allow repl to take a file as arg
- [ ] don't allow tokens to have spaces in between them
- [ ] add <= and >= operators
- [ ] enhance error messages with line number and file name
- [ ] add postfix operators (such as `++`)
- [ ] prettier printing of function, array, and hash values
