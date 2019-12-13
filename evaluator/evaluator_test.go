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

	return Eval(program)
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
