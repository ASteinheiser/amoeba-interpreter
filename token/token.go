package token

// Token is a single element
type Token struct {
	Type    Type
	Literal string
}

// Type is the type of token
type Type string

const (
	// ILLEGAL : bad token
	ILLEGAL = "ILLEGAL"
	// EOF : end of file
	EOF = "EOF"
	// IDENT : identifiers or variable names
	IDENT = "IDENT"
	// INT : integer literal
	INT = "INT"
	// ASSIGN : sets an identifier equal to a literal
	ASSIGN = "="
	// PLUS : adds two integers
	PLUS = "+"
	// BANG : inverts an expression
	BANG = "!"
	// MINUS : subtracts two numbers
	MINUS = "-"
	// SLASH : divides two numbers
	SLASH = "/"
	// ASTERISK : multiplies two numbers
	ASTERISK = "*"
	// LT : checks if a number is less than another
	LT = "<"
	// GT : checks if a number is greater than another
	GT = ">"
	// COMMA : separator for elements in a list
	COMMA = ","
	// SEMICOLON : ends a statement
	SEMICOLON = ";"
	// LPAREN : start listing function call params
	LPAREN = "("
	// RPAREN : stop listing function call params
	RPAREN = ")"
	// LBRACE : open new block of code
	LBRACE = "{"
	// RBRACE : close block of code
	RBRACE = "}"
	// FUNCTION : create a function that accepts
	// params and returns a value
	FUNCTION = "FUNCTION"
	// LET : initialize a new identifier
	LET = "LET"
)

var keywords = map[string]Type{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent returns the token type for a
// keyword or user-defined indentifier
func LookupIdent(ident string) Type {
	if key, ok := keywords[ident]; ok {
		return key
	}
	return IDENT
}
