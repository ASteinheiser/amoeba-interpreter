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

	if !testLiteralExpression(t, stmt.Expression, "varName") {
		return
	}
}

func testIdentifierLiteral(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp is not a *ast.Identifier, got=\"%T\"",
			exp)
		return false
	}

	if ident.Value != value {
		t.Errorf("ident.Value is not \"%s\", got=\"%s\"", value, ident.Value)
		return false
	}

	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral() is not \"%s\", got=\"%s\"",
			value, ident.TokenLiteral())
		return false
	}

	return true
}

func TestIntegerExpression(t *testing.T) {
	input := "400;"

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

	if !testLiteralExpression(t, stmt.Expression, 400) {
		return
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integer, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il is not *ast.IntegerLiteral, got=\"%T\"", il)
		return false
	}

	if integer.Value != value {
		t.Errorf("integer.Value was not \"%d\", got=\"%d\"", value, integer.Value)
		return false
	}

	if integer.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integer.TokenLiteral() was not \"%d\", got=\"%s\"",
			value, integer.TokenLiteral())
		return false
	}

	return true
}

func TestBooleanExpression(t *testing.T) {
	input := `
		false;
		true;
	`
	expected := [2]bool{false, true}

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 2 {
		t.Fatalf("expected program to have 2 statements, got=%d",
			len(program.Statements))
	}

	for i, stmt := range program.Statements {
		boolean, ok := stmt.(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("stmt is not a *ast.ExpressionStatement, got=%T",
				program.Statements[0])
		}

		if !testLiteralExpression(t, boolean.Expression, expected[i]) {
			return
		}
	}
}

func testBooleanLiteral(t *testing.T, b ast.Expression, value bool) bool {
	boolean, ok := b.(*ast.BooleanLiteral)
	if !ok {
		t.Errorf("b is not *ast.BooleanLiteral, got=\"%T\"", b)
		return false
	}

	if boolean.Value != value {
		t.Errorf("boolean.Value was not \"%t\", got=\"%t\"", value, boolean.Value)
		return false
	}

	if boolean.TokenLiteral() != fmt.Sprintf("%t", value) {
		t.Errorf("boolean.TokenLiteral() was not \"%t\", got=\"%s\"",
			value, boolean.TokenLiteral())
		return false
	}

	return true
}

func TestPrefixExpression(t *testing.T) {
	tests := []struct {
		input    string
		operator string
		value    interface{}
	}{
		{"!4;", "!", 4},
		{"-17;", "-", 17},
		{"!indentifier;", "!", "indentifier"},
		{"!true", "!", true},
		{"!false", "!", false},
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

		if !testLiteralExpression(t, exp.Right, test.value) {
			return
		}
	}
}

func TestInfixExpression(t *testing.T) {
	tests := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{"4 + 7", 4, "+", 7},
		{"4 - 7", 4, "-", 7},
		{"4 * 7", 4, "*", 7},
		{"4 / 7", 4, "/", 7},
		{"4 > 7", 4, ">", 7},
		{"4 < 7", 4, "<", 7},
		{"4 == 7", 4, "==", 7},
		{"4 != 7", 4, "!=", 7},
		{"bob != joe", "bob", "!=", "joe"},
		{"something + anotherVar", "something", "+", "anotherVar"},
		{"true != false", true, "!=", false},
		{"true == true", true, "==", true},
		{"false == false", false, "==", false},
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

		if !testInfixLiteral(t, stmt.Expression,
			test.leftValue, test.operator, test.rightValue) {
			return
		}
	}
}

func testInfixLiteral(
	t *testing.T,
	exp ast.Expression,
	left interface{},
	operator string,
	right interface{},
) bool {
	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("exp is not a *ast.InfixExpression, got=\"%T\"(%s)", exp, exp)
		return false
	}

	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}

	if opExp.Operator != operator {
		t.Fatalf("opExp.Operator is not \"%s\", got=\"%s\"",
			operator, opExp.Operator)
		return false
	}

	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}

	return true
}

func testLiteralExpression(
	t *testing.T,
	exp ast.Expression,
	expected interface{},
) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case string:
		return testIdentifierLiteral(t, exp, v)
	case bool:
		return testBooleanLiteral(t, exp, v)
	}
	t.Errorf("type of exp not handled. got=\"%T\"", exp)
	return false
}

func TestPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"false",
			"false",
		},
		{
			"true",
			"true",
		},
		{
			"4 > 6 == true",
			"((4 > 6) == true)",
		},
		{
			"4 < 6 == false",
			"((4 < 6) == false)",
		},
		{
			"-x * y",
			"((-x) * y)",
		},
		{
			"!-x",
			"(!(-x))",
		},
		{
			"x + y + z",
			"((x + y) + z)",
		},
		{
			"x + y - z",
			"((x + y) - z)",
		},
		{
			"x * y * z",
			"((x * y) * z)",
		},
		{
			"x * y / z",
			"((x * y) / z)",
		},
		{
			"x + y * z + x / y - z",
			"(((x + (y * z)) + (x / y)) - z)",
		},
		{
			"4 + 5; -7 * 18",
			"(4 + 5)((-7) * 18)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		actual := program.String()
		if actual != test.expected {
			t.Errorf("expected:\n\"%s\"\nto evaluate to:\n\"%s\"\nbut received:\n\"%s\"",
				test.input, test.expected, actual)
		}
	}
}
