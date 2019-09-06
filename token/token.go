package token

// Token is a single element
type Token struct {
	Type    tokenType
	Literal string
}

type tokenType string

const (
	// ILLEGAL : bad token
	ILLEGAL = "ILLEGAL"
	// EOF : end of file
	EOF = "EOF"
	// IDENT : identifiers or variables
	IDENT = "IDENT"
	// INT : integers
	INT = "INT"
	// ASSIGN : sets an identifier equal to a literal
	ASSIGN = "="
	// PLUS : adds two integers
	PLUS = "+"
	// COMMA : separator for elements in a list
	COMMA = ","
	// SEMICOLON : ends a line of code
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
