// Package ast is an Abstract Syntax Tree, which will be used to
// represent Amoeba programs as data that we can later evaluate
package ast

import (
	"bytes"

	"github.com/ASteinheiser/amoeba-interpreter/token"
)

// Node is a single element in a program
type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// LetStatement is a Statement Node that assigns an expression to an identifier
type LetStatement struct {
	Token token.Token // should be a LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral returns the token literal for the let statement
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// Identifier is an Expression Node that is bound to an expression,
// used to create a new variable or return a variables value
type Identifier struct {
	Token token.Token // should be an IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns the token literal for the identifier expression
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (i *Identifier) String() string { return i.Value }

// ReturnStatement is a Statement Node that ends a function call and
// returns an expression to the caller
type ReturnStatement struct {
	Token       token.Token // should be a RETURN token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral returns the token literal for the return statement
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement is a Statement Node consisting solely of an expression
type ExpressionStatement struct {
	Token      token.Token // first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral returns the token literal for
// the first token in the expression statement
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// IntegerLiteral is an Expression Node consisting solely of an integer
type IntegerLiteral struct {
	Token token.Token // should be an INT token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral returns the token literal for the integer
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

func (il *IntegerLiteral) String() string { return il.Token.Literal }

// PrefixExpression is an Expression Node that applies
// a prefix to another expression
type PrefixExpression struct {
	Token    token.Token // should be a prefix token ("!" or "-")
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral returns the token literal for the prefix expression
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// InfixExpression is an Expression Node that applies
// an operator to the expressions on either side of it
type InfixExpression struct {
	Token    token.Token // should be an infix token ("+", "/", "==", ">")
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}

// TokenLiteral returns the token literal for the infix expression
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }

func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}
