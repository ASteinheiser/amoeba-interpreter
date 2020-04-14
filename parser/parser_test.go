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
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{`let z = 4;`, "z", 4},
		{`let y = false;`, "y", false},
		{`let foobar = something;`, "foobar", "something"},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not have 1 statement, got=%d",
				len(program.Statements))
		}

		stmt := program.Statements[0]
		if !testLetStatement(t, stmt, test.expectedIdentifier) {
			return
		}

		val := stmt.(*ast.LetStatement).Value

		if !testLiteralExpression(t, val, test.expectedValue) {
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
	tests := []struct {
		input              string
		expectedExpression interface{}
	}{
		{`return 7;`, 7},
		{`return true;`, true},
		{`return foobar;`, "foobar"},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not have 1 statement, got=%d",
				len(program.Statements))
		}

		stmt := program.Statements[0]
		if !testReturnStatement(t, stmt) {
			return
		}

		val := stmt.(*ast.ReturnStatement).ReturnValue

		if !testLiteralExpression(t, val, test.expectedExpression) {
			return
		}
	}
}

func testReturnStatement(t *testing.T, s ast.Statement) bool {
	returnStmt, ok := s.(*ast.ReturnStatement)
	if !ok {
		t.Errorf("stmt not *ast.ReturnStatement, got=%T", s)
		return false
	}

	if returnStmt.TokenLiteral() != "return" {
		t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
			returnStmt.TokenLiteral())
		return false
	}

	return true
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

func TestStringExpression(t *testing.T) {
	input := `"hello world!"`
	expected := "hello world!"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not a *ast.ExpressionStatement, got=%T",
			program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.StringLiteral)
	if !ok {
		t.Fatalf("literal is not a *ast.StringLiteral, got=%T", stmt.Expression)
	}

	if literal.Value != expected {
		t.Errorf("literal.Value is not %q, got=%q", expected, literal.Value)
	}
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
		t.Errorf("boolean is not *ast.BooleanLiteral, got=\"%T\"", b)
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
		{
			"1 + (2 + 3) + 4",
			"((1 + (2 + 3)) + 4)",
		},
		{
			"(4 + 4) * 3",
			"((4 + 4) * 3)",
		},
		{
			"2 / (4 + 4)",
			"(2 / (4 + 4))",
		},
		{
			"-(5 + 5)",
			"(-(5 + 5))",
		},
		{
			"!(true == true)",
			"(!(true == true))",
		},
		{
			"a + add(b * c) + d",
			"((a + add((b * c))) + d)",
		},
		{
			"add(a, b, 1, 2 * 3, 4 - 5, add(6, 7 + 8))",
			"add(a, b, 1, (2 * 3), (4 - 5), add(6, (7 + 8)))",
		},
		{
			"add(a + b + c * d / f + g)",
			"add((((a + b) + ((c * d) / f)) + g))",
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

func TestIfExpression(t *testing.T) {
	input := `
		if (x > y) {
			x
		}
	`

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

	exp, ok := stmt.Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not a *ast.IfExpression, got=%T",
			stmt.Expression)
	}

	if !testInfixLiteral(t, exp.Condition, "x", ">", "y") {
		return
	}

	if len(exp.Consequence.Statements) != 1 {
		t.Errorf("exp.Consequence is not 1 statement, got=%d\n",
			len(exp.Consequence.Statements))
	}

	consequence, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statements[0] is not a *ast.ExpressionStatement, got=%T",
			exp.Consequence.Statements[0])
	}

	if !testIdentifierLiteral(t, consequence.Expression, "x") {
		return
	}

	if exp.Alternative != nil {
		t.Errorf("exp.Alternative.Statements was not nil. got=%+v", exp.Alternative)
	}
}

func TestIfElseExpression(t *testing.T) {
	input := `
		if (x > y) {
			x
		} else {
			7
		}
	`

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

	exp, ok := stmt.Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not a *ast.IfExpression, got=%T",
			stmt.Expression)
	}

	if !testInfixLiteral(t, exp.Condition, "x", ">", "y") {
		return
	}

	if len(exp.Consequence.Statements) != 1 {
		t.Errorf("exp.Consequence is not 1 statement, got=%d\n",
			len(exp.Consequence.Statements))
	}

	consequence, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statements[0] is not a *ast.ExpressionStatement, got=%T",
			exp.Consequence.Statements[0])
	}

	if !testLiteralExpression(t, consequence.Expression, "x") {
		return
	}

	if len(exp.Alternative.Statements) != 1 {
		t.Errorf("exp.Alternative is not 1 statement, got=%d\n",
			len(exp.Alternative.Statements))
	}

	alternative, ok := exp.Alternative.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statements[0] is not a *ast.ExpressionStatement, got=%T",
			exp.Alternative.Statements[0])
	}

	if !testLiteralExpression(t, alternative.Expression, 7) {
		return
	}
}

func TestFunctionLiteral(t *testing.T) {
	input := `
		fn(x, z) {
			x * z;
		}
	`

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

	function, ok := stmt.Expression.(*ast.FunctionLiteral)
	if !ok {
		t.Fatalf("stmt.Expression is not a *ast.FunctionLiteral, got=%T",
			stmt.Expression)
	}

	if len(function.Parameters) != 2 {
		t.Fatalf("function.Parameters not equal to 2, got=%d\n",
			len(function.Parameters))
	}

	testLiteralExpression(t, function.Parameters[0], "x")
	testLiteralExpression(t, function.Parameters[1], "z")

	if len(function.Body.Statements) != 1 {
		t.Fatalf("function.Body.Statements not equal to 1, got=%d\n",
			len(function.Parameters))
	}

	bodyStmt, ok := function.Body.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("function body statement is not a *ast.ExpressionStatement, got=%T",
			function.Body.Statements[0])
	}

	testInfixLiteral(t, bodyStmt.Expression, "x", "*", "z")
}

func TestFunctionParameterParsing(t *testing.T) {
	tests := []struct {
		input          string
		expectedParams []string
	}{
		{input: `fn() {};`, expectedParams: []string{}},
		{input: `fn(x) {};`, expectedParams: []string{"x"}},
		{input: `fn(x, y, z) {};`, expectedParams: []string{"x", "y", "z"}},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		stmt := program.Statements[0].(*ast.ExpressionStatement)
		function := stmt.Expression.(*ast.FunctionLiteral)

		if len(function.Parameters) != len(test.expectedParams) {
			t.Errorf("function parameters length wrong, expected=%d, got=%d",
				len(test.expectedParams), len(function.Parameters))
		}

		for i, ident := range test.expectedParams {
			testLiteralExpression(t, function.Parameters[i], ident)
		}
	}
}

func TestCallExpression(t *testing.T) {
	input := `doSomethin(1, 2 - 3, 4 * 5)`

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

	call, ok := stmt.Expression.(*ast.CallExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not a *ast.CallExpression, got=%T",
			stmt.Expression)
	}

	if !testIdentifierLiteral(t, call.Function, "doSomethin") {
		return
	}

	if len(call.Arguments) != 3 {
		t.Fatalf("call.Arguments length not 3, got=%d", len(call.Arguments))
	}

	testLiteralExpression(t, call.Arguments[0], 1)
	testInfixLiteral(t, call.Arguments[1], 2, "-", 3)
	testInfixLiteral(t, call.Arguments[2], 4, "*", 5)
}

func TestCallExpressionParameterParsing(t *testing.T) {
	tests := []struct {
		input         string
		expectedIdent string
		expectedArgs  []string
	}{
		{
			input:         "doSomethin();",
			expectedIdent: "doSomethin",
			expectedArgs:  []string{},
		},
		{
			input:         "doSomethin(1);",
			expectedIdent: "doSomethin",
			expectedArgs:  []string{"1"},
		},
		{
			input:         "doSomethin(1, 2 - 3, 4 * 5);",
			expectedIdent: "doSomethin",
			expectedArgs:  []string{"1", "(2 - 3)", "(4 * 5)"},
		},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		stmt := program.Statements[0].(*ast.ExpressionStatement)
		call, ok := stmt.Expression.(*ast.CallExpression)
		if !ok {
			t.Fatalf("stmt.Expression is not ast.CallExpression, got=%T",
				stmt.Expression)
		}

		if !testLiteralExpression(t, call.Function, test.expectedIdent) {
			return
		}

		if len(call.Arguments) != len(test.expectedArgs) {
			t.Fatalf("call.Arguments length not %d, got=%d",
				len(test.expectedArgs), len(call.Arguments))
		}

		for i, arg := range test.expectedArgs {
			if call.Arguments[i].String() != arg {
				t.Errorf("argument %d wrong, expected=%q, got=%q", i,
					arg, call.Arguments[i].String())
			}
		}
	}
}

func TestArrayLiteralParsing(t *testing.T) {
	input := `[1, 2 * 4, 3 + 5, true]`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	array, ok := stmt.Expression.(*ast.ArrayLiteral)
	if !ok {
		t.Fatalf("stmt is not *ast.ArrayLiteral, got=%T", stmt.Expression)
	}

	if len(array.Elements) != 4 {
		t.Fatalf("length of array is wrong. wanted=4, got=%d", len(array.Elements))
	}

	testIntegerLiteral(t, array.Elements[0], 1)
	testInfixLiteral(t, array.Elements[1], 2, "*", 4)
	testInfixLiteral(t, array.Elements[2], 3, "+", 5)
	testBooleanLiteral(t, array.Elements[3], true)
}
