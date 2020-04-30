# amoeba-interpreter
An interpreter for the Amoeba programming language. Written from scratch in Go with no dependencies 🎉

<img
  src="https://s3-us-west-2.amazonaws.com/images.iamandrew.io/Screen+Shot+2020-04-29+at+11.27.28+PM.png"
  width="500px"
  alt="Amoeba Screenshot"
/>

## Features
- C-like syntax
- variables (integers, booleans, strings, arrays, objects)
- arithmetic expressions
- first-class and higher-order functions
- closures
- builtin functions:
  - amoeba(): prints out awesome ascii art
  - len(ARRAY or STRING): returns length of string or array
  - push(ARRAY, ANY): adds new item to array (does not mutate)
  - first(ARRAY): returns first item in array
  - rest(ARRAY): returns all but first item in array
  - last(ARRAY): returns last item in array
  - print(ANY, ANY, ...): prints out to the console

## Give the Amoeba REPL a try!
1. `git clone https://github.com/ASteinheiser/amoeba-interpreter.git`
1. `cd amoeba-interpreter`
1. `./amoeba-interpreter`

<img
  src="https://s3-us-west-2.amazonaws.com/images.iamandrew.io/Screen+Shot+2020-04-29+at+6.11.42+PM.png"
  width="800px"
  alt="REPL Screenshot"
/>

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
