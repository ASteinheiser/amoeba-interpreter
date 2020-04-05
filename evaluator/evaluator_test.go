package evaluator

import (
	"testing"

	"github.com/ASteinheiser/amoeba-interpreter/lexer"
	"github.com/ASteinheiser/amoeba-interpreter/object"
	"github.com/ASteinheiser/amoeba-interpreter/parser"
)

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return Eval(program, env)
}

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"4", 4},
		{"720", 720},
		{"-4", -4},
		{"-720", -720},
		{"4 + 4 + 4 - 2", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-40 + 100 + -50", 10},
		{"4 * 3 - 2", 10},
		{"4 - 3 * 2", -2},
		{"10 + 2 * -2", 6},
		{"40 / 2 * 4 + -10", 70},
		{"40 * (4 - 3) / 4", 10},
		{"(4 + 6 * 2 + 12 / 4) * 2 + -10", 28},
	}

	for _, test := range tests {
		evaluated := testEval(test.input)
		testIntegerObject(t, evaluated, test.expected)
	}
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("obj not of type *object.Integer, got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("expected obj.Value to be %d, got=%d", expected, result.Value)
		return false
	}

	return true
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false == true", false},
		{"(1 < 2) == true", true},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
		{"(1 > 2) == true", false},
	}

	for _, test := range tests {
		evaluated := testEval(test.input)
		testBooleanObject(t, evaluated, test.expected)
	}
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("obj not of type *object.Boolean, got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("expected obj.Value to be %t, got=%t", expected, result.Value)
		return false
	}

	return true
}

func TestEvalBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!4", false},
		{"!!true", true},
		{"!!false", false},
		{"!!4", true},
	}

	for _, test := range tests {
		evaluated := testEval(test.input)
		testBooleanObject(t, evaluated, test.expected)
	}
}

func TestEvalIfElseExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if(true) { 10 }", 10},
		{"if(false) { 10 }", nil},
		{"if(1) { 10 }", 10},
		{"if(1 < 2) { 10 }", 10},
		{"if(1 > 2) { 10 }", nil},
		{"if(1 > 2) { 10 } else { 20 }", 20},
		{"if(1 < 2) { 10 } else { 20 }", 10},
	}

	for _, test := range tests {
		evaluated := testEval(test.input)
		integer, ok := test.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("obj is not NULL, got=%T (%+v)", obj, obj)
		return false
	}
	return true
}

func TestEvalReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"return 4;", 4},
		{"return 4; 6;", 4},
		{"return 2 * 4; 6;", 8},
		{"5; return 2 * 2 + 3; 8;", 7},
		{
			`
			if (4 > 1) {
				if (3 > 1) {
					return 100;
				}
				return 0;
			}
			`,
			100,
		},
	}

	for _, test := range tests {
		evaluated := testEval(test.input)
		testIntegerObject(t, evaluated, test.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"4 + false;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"4 + false; 4;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-true;",
			"unknown operator: -BOOLEAN",
		},
		{
			"true + false;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; true + false; 4;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (4 > 1) { true + false; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`
			if (4 > 1) {
				if (3 > 1) {
					return true + false;
				}
				return 7;
			}
			`,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`
			if (4 > 1) {
				if (3 > 1) {
					true - false;
				}
				return 8;
			}
			`,
			"unknown operator: BOOLEAN - BOOLEAN",
		},
		{
			"someIdentifier;",
			"identifier not found: someIdentifier",
		},
	}

	for _, test := range tests {
		evaluated := testEval(test.input)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("errObj not of type *object.Error, got=%T (%+v)", evaluated, evaluated)
			continue
		}

		if errObj.Message != test.expectedMessage {
			t.Errorf("expected errObj.Message to be %q, got=%q", test.expectedMessage, errObj.Message)
		}
	}
}

func TestEvalLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let x = 4; x;", 4},
		{"let x = 4 * 5; x;", 20},
		{"let x = 8; let z = x; z;", 8},
		{"let x = 6; let z = x; let y = x + z + 3; y;", 15},
	}

	for _, test := range tests {
		testIntegerObject(t, testEval(test.input), test.expected)
	}
}

func TestEvalFunctionExpression(t *testing.T) {
	input := "fn(y) { y + 4 };"
	expectedParam := "y"
	expectedBody := "(y + 4)"

	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("fn not of type *object.Function, got=%T (%+v)", evaluated, evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("function parameters are incorrect, got=%+v", fn.Parameters)
	}

	if fn.Parameters[0].String() != expectedParam {
		t.Fatalf("parameter is not %q, got=%q", expectedParam, fn.Parameters[0])
	}

	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q, got=%q", expectedBody, fn.Body.String())
	}
}

func TestEvalFunctionCall(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let identify = fn(x) { x; }; identify(4);", 4},
		{"let identify = fn(x) { return x }; identify(4)", 4},
		{"let double = fn(x) { x * 2 }; double(4)", 8},
		{"let add = fn(x, y) { x + y }; add(4, 7)", 11},
		{"let add = fn(x, y) { x + y }; add(5 + 5, add(7 + 3))", 20},
		{"fn(x) { x }(4)", 4},
	}

	for _, test := range tests {
		testIntegerObject(t, testEval(test.input), test.expected)
	}
}
