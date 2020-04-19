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
		{`"hello" == "hello"`, true},
		{`"hello" == "world"`, false},
		{`"hello" != "hello"`, false},
		{`"hello" != "world"`, true},
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
		{
			`"Hello" - "World"`,
			"unknown operator: STRING - STRING",
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
		{"let add = fn(x, y) { x + y }; add(5 + 5, add(7, 3))", 20},
		{"fn(x) { x }(4)", 4},
	}

	for _, test := range tests {
		testIntegerObject(t, testEval(test.input), test.expected)
	}
}

func TestClosures(t *testing.T) {
	input := `
		let newAdder = fn(x) {
			fn(y) { x + y }
		}
		let addFour = newAdder(4)
		addFour(6)
	`

	testIntegerObject(t, testEval(input), 10)
}

func testStringObject(t *testing.T, obj object.Object, expected string) bool {
	result, ok := obj.(*object.String)
	if !ok {
		t.Errorf("obj not of type *object.String, got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("expected obj.Value to be %q, got=%q", expected, result.Value)
		return false
	}

	return true
}

func TestEvalStringExpression(t *testing.T) {
	input := `"Hello world!"`
	expected := "Hello world!"

	evaluated := testEval(input)

	testStringObject(t, evaluated, expected)
}

func TestEvalStringConcatenation(t *testing.T) {
	input := `"Hello" + " " + "World!"`
	expected := "Hello World!"

	evaluated := testEval(input)

	testStringObject(t, evaluated, expected)
}

func TestEvalBuiltinFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`len("")`, 0},
		{`len("four")`, 4},
		{`len(" hey world ! ")`, 13},
		{`len([])`, 0},
		{`len([1, 2, 3, 4])`, 4},
		{`len(["sup", 2, true, 4 + 7])`, 4},
		{`len(1)`, "argument to `len` not supported: INTEGER"},
		{`len(true)`, "argument to `len` not supported: BOOLEAN"},
		{`len("one", "two")`, "wrong number of arguments passed to `len`: got 2, want 1"},
		{`first([])`, nil},
		{`first([1, 2, 3, 4])`, 1},
		{`first("yo!")`, "argument to `first` must be ARRAY, got STRING"},
		{`first([1, 2], [3, 4])`, "wrong number of arguments passed to `first`: got 2, want 1"},
		{`last([])`, nil},
		{`last([1, 2, 3, 4])`, 4},
		{`last("yo!")`, "argument to `last` must be ARRAY, got STRING"},
		{`last([1, 2], [3, 4])`, "wrong number of arguments passed to `last`: got 2, want 1"},
		{`rest([])`, nil},
		{`rest([1])`, []int{}},
		{`rest([1, 2, 3, 4])`, []int{2, 3, 4}},
		{`rest(rest(rest([1, 2, 3, 4])))`, []int{4}},
		{`rest(rest(rest(rest([1, 2, 3, 4]))))`, []int{}},
		{`rest(rest(rest(rest(rest([1, 2, 3, 4])))))`, nil},
		{`rest("yo!")`, "argument to `rest` must be ARRAY, got STRING"},
		{`rest([1, 2], [3, 4])`, "wrong number of arguments passed to `rest`: got 2, want 1"},
		{`push([], 1)`, []int{1}},
		{`push([1], 2)`, []int{1, 2}},
		{`push([1, 2, 3, 4], 6 + 3)`, []int{1, 2, 3, 4, 9}},
		{`push("yo!", 5)`, "first argument to `push` must be ARRAY, got STRING"},
		{`push([1, 2])`, "wrong number of arguments passed to `push`: got 1, want 2"},
		{`push([1, 2], 1, true)`, "wrong number of arguments passed to `push`: got 3, want 2"},
	}

	for _, test := range tests {
		evaluated := testEval(test.input)

		switch expected := test.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case []int:
			arr, ok := evaluated.(*object.Array)
			if !ok {
				t.Errorf("arr is not *object.Array, got=%T (%+v)", evaluated, evaluated)
			}
			for idx, elem := range arr.Elements {
				testIntegerObject(t, elem, int64(expected[idx]))
			}
		case string:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("errObj is not *object.Error, got=%T (%+v)", evaluated, evaluated)
			}
			if errObj.Message != expected {
				t.Errorf("errObj.Message is wrong. expected=%q, got=%q", expected, errObj.Message)
			}
		default:
			testNullObject(t, evaluated)
		}
	}
}

func TestEvalArrayLiterals(t *testing.T) {
	input := "[1, 2 * 4, 5 + 8]"

	evaluated := testEval(input)
	array, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("array not of type *object.Array, got=%T (%+v)", evaluated, evaluated)
	}

	if len(array.Elements) != 3 {
		t.Fatalf("array has wrong number of elements. got=%d, want=3", len(array.Elements))
	}

	testIntegerObject(t, array.Elements[0], 1)
	testIntegerObject(t, array.Elements[1], 8)
	testIntegerObject(t, array.Elements[2], 13)
}

func TestEvalArrayIndexExpessions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"[4, 3, 2][0]", 4},
		{"[4, 3, 2][1]", 3},
		{"[4, 3, 2][2]", 2},
		{"[4, 3, 2][3]", nil},
		{"[4, 3, 2][-1]", nil},
		{"let i = 0; [4, 3, 2][i]", 4},
		{"[4, 3, 2][1 + 1]", 2},
		{"let myArray = [4, 3, 2]; myArray[1]", 3},
		{"let myArray = [4, 3, 2]; myArray[0] + myArray[1] + myArray[2]", 9},
		{"let myArray = [4, 3, 2]; let i = myArray[0] - 3; myArray[i]", 3},
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
