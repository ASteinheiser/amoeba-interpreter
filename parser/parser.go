package parser

import (
	"github.com/ASteinheiser/amoeba-interpreter/ast"
	"github.com/ASteinheiser/amoeba-interpreter/lexer"
	"github.com/ASteinheiser/amoeba-interpreter/token"
)

// Parser contains the lexer and parses tokens one at a time,
// generating an AST
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// New creates a new Parser from a given lexer
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// initialize both tokens by reading twice
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram generates an AST based on the input
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
