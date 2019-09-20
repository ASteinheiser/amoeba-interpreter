package parser

import (
	"fmt"
	"testing"

	"github.com/ASteinheiser/amoeba-interpreter/ast"
	"github.com/ASteinheiser/amoeba-interpreter/lexer"
)

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	pluralS := "s"
	if len(errors) == 1 {
		pluralS = ""
	}
	t.Errorf("parser has %d error%s", len(errors), pluralS)

	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestLetStatements(t *testing.T) {
	input := `
		let x = 4;
		let z = 16;
		let blahblah = 90091;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"z"},
		{"blahblah"},
	}

	for i, test := range tests {
		curStatement := program.Statements[i]
		if !testLetStatement(t, curStatement, test.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}

	return true
}

func TestReturnStatements(t *testing.T) {
	input := `
		return 4;
		return 17;
		return 8801;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements, got=%d",
			len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement, got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
				returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "varName;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("expected program to have 1 statement, got=%d",
			len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not a *ast.ExpressionStatement, got=%T",
			program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("stmt.Expression is not a *ast.Identifier, got=%T",
			stmt.Expression)
	}
	if ident.Value != "varName" {
		t.Errorf("ident.Value is not \"%s\", got=\"%s\"",
			"varName", ident.Value)
	}
	if ident.TokenLiteral() != "varName" {
		t.Errorf("ident.TokenLiteral() is not \"%s\", got=\"%s\"",
			"varName", ident.TokenLiteral())
	}
}

func TestIntegerExpression(t *testing.T) {
	input := "4;"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("expected program to have 1 statement, got=%d",
			len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not a *ast.ExpressionStatement, got=%T",
			program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("stmt.Expression is not a *ast.IntegerLiteral, got=%T",
			stmt.Expression)
	}
	if literal.Value != 4 {
		t.Errorf("literal.Value is not \"%d\", got=\"%d\"",
			4, literal.Value)
	}
	if literal.TokenLiteral() != "4" {
		t.Errorf("literal.TokenLiteral() is not \"%s\", got=\"%s\"",
			"4", literal.TokenLiteral())
	}
}

func TestPrefixExpression(t *testing.T) {
	tests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!4;", "!", 4},
		{"-17;", "-", 17},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("expected program to have 1 statement, got=%d",
				len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not a *ast.ExpressionStatement, got=%T",
				program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt.Expression is not a *ast.PrefixExpression, got=%T",
				stmt.Expression)
		}
		if exp.Operator != test.operator {
			t.Fatalf("exp.Operator is not '%s', got=\"%s\"",
				test.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, test.integerValue) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integer, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il is not *ast.IntegerLiteral, got=%T", il)
		return false
	}

	if integer.Value != value {
		t.Errorf("integer.Value was not '%d', got=%d", value, integer.Value)
		return false
	}

	if integer.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integer.TokenLiteral() was not '%d', got=%s",
			value, integer.TokenLiteral())
	}

	return true
}

func TestInfixExpression(t *testing.T) {
	tests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"4 + 7", 4, "+", 7},
		{"4 - 7", 4, "-", 7},
		{"4 * 7", 4, "*", 7},
		{"4 / 7", 4, "/", 7},
		{"4 > 7", 4, ">", 7},
		{"4 < 7", 4, "<", 7},
		{"4 == 7", 4, "==", 7},
		{"4 != 7", 4, "!=", 7},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("expected program to have 1 statement, got=%d",
				len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not a *ast.ExpressionStatement, got=%T",
				program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("stmt.Expression is not a *ast.InfixExpression, got=%T",
				stmt.Expression)
		}

		if !testIntegerLiteral(t, exp.Left, test.leftValue) {
			return
		}

		if exp.Operator != test.operator {
			t.Fatalf("exp.Operator is not '%s', got=\"%s\"",
				test.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, test.rightValue) {
			return
		}
	}
}
