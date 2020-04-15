// Package ast is an Abstract Syntax Tree, which will be used to
// represent Amoeba programs as data that we can later evaluate
package ast

import (
	"bytes"
	"strings"

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

// StringLiteral is an Expression Node consisting solely of a string
type StringLiteral struct {
	Token token.Token // should be a STRING token
	Value string
}

func (sl *StringLiteral) expressionNode() {}

// TokenLiteral returns the token literal for the string
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }

func (sl *StringLiteral) String() string { return sl.Token.Literal }

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

// BooleanLiteral is an Expression Node consisting solely of a boolean
type BooleanLiteral struct {
	Token token.Token // should be a TRUE or FALSE token
	Value bool
}

func (b *BooleanLiteral) expressionNode() {}

// TokenLiteral returns the token literal for the integer
func (b *BooleanLiteral) TokenLiteral() string { return b.Token.Literal }

func (b *BooleanLiteral) String() string { return b.Token.Literal }

// IfExpression is an Expression Node representing a conditional statement
type IfExpression struct {
	Token       token.Token // should be an IF token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

// TokenLiteral returns the token literal for the if
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }

func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

// BlockStatement is a Statement Node that contains statements
type BlockStatement struct {
	Token      token.Token // should be a { token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

// TokenLiteral returns the token literal for the beginning of the block: {
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }

func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// FunctionLiteral is an Expression Node representing a function
type FunctionLiteral struct {
	Token      token.Token // should be an 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}

// TokenLiteral returns the token literal for the if
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }

func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(fl.Body.String())
	out.WriteString("\n}")

	return out.String()
}

// CallExpression is an Expression Node representing a function call
type CallExpression struct {
	Token     token.Token // should be a ( token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}

// TokenLiteral returns the token literal for the (
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }

func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

// ArrayLiteral is an Expression Node representing an array
type ArrayLiteral struct {
	Token    token.Token // should be a [ token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode() {}

// TokenLiteral returns the token literal for the array: [
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }

func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// IndexExpression is an Expression Node representing the access of an array index
type IndexExpression struct {
	Token token.Token // should be a [ token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode() {}

// TokenLiteral returns the token literal for the index expression: [
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }

func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}
