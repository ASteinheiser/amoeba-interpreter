package lexer

// Lexer stores the program to be interpreted
// and reads over the characters one at a time
type Lexer struct {
	input    string
	position int  // current character position
	readPos  int  // reading position, ahead of current char
	ch       byte // current character value
}

// New will create a Lexer to interpret the code
func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
