package ast

import "github.com/ASteinheiser/amoeba-interpreter/token"

// Node is a single element in a program
type Node interface {
	TokenLiteral() string
}

// Statement is a Node that tells the language to do something
type Statement interface {
	Node
	statementNode()
}

// Expression is a Node that tells the language to evaluate something
type Expression interface {
	Node
	expressionNode()
}

// Program is the root Node of all amoeba programs and contains all
// the statements for a given peice of code
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the token literal for the next statement
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement is a Statement Node that assigns an expression to an identifier
type LetStatement struct {
	Token token.Token // should be token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral returns the token literal for the let statement
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier is an Expression Node that is bound to an expression
// Used to create a new variable or return a variables value
type Identifier struct {
	Token token.Token // should be token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns the token literal for the identifier expression
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
