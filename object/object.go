package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/ASteinheiser/amoeba-interpreter/ast"
)

// Type is the type of an object
type Type string

const (
	// INTEGER_OBJ is the object type for integers
	INTEGER_OBJ = "INTEGER"
	// BOOLEAN_OBJ is the object type for booleans
	BOOLEAN_OBJ = "BOOLEAN"
	// NULL_OBJ is the object type for nulls
	NULL_OBJ = "NULL"
	// RETURN_VALUE_OBJ is the object type for return values
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	// ERROR_OBJ is the object type for errors
	ERROR_OBJ = "ERROR"
	// FUNCTION_OBJ is the object type for functions
	FUNCTION_OBJ = "FUNCTION"
	// STRING_OBJ is the object type for strings
	STRING_OBJ = "STRING"
	// BUILTIN_OBJ is the object type for builtin functions
	BUILTIN_OBJ = "BUILTIN"
	// ARRAY_OBJ is the object type for arrays
	ARRAY_OBJ = "ARRAY"
)

// BuiltinFunction is the type for functions defined by the interpreter
type BuiltinFunction func(args ...Object) Object

// Object is a wrapper for values that we evaluate
type Object interface {
	Type() Type
	Inspect() string
}

// Integer is the object that holds integers
type Integer struct {
	Value int64
}

// Inspect returns a formatted string with the value of the integer
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Type returns the type string for the integer
func (i *Integer) Type() Type { return INTEGER_OBJ }

// Boolean is the object that holds booleans
type Boolean struct {
	Value bool
}

// Inspect returns a formatted string with the value of the boolean
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// Type returns the type string for the boolean
func (b *Boolean) Type() Type { return BOOLEAN_OBJ }

// Null is the object that holds nulls
type Null struct{}

// Inspect returns a string with the value of the null
func (n *Null) Inspect() string { return "null" }

// Type returns the type string for the null
func (n *Null) Type() Type { return NULL_OBJ }

// ReturnValue is the object that holds return values
type ReturnValue struct {
	Value Object
}

// Inspect returns a string with the value of the return value
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

// Type returns the type string for the return value
func (rv *ReturnValue) Type() Type { return RETURN_VALUE_OBJ }

// Error is the object that holds internal error messages
type Error struct {
	Message string
}

// Inspect returns a string with the message of the error
func (e *Error) Inspect() string { return "ERROR: " + e.Message }

// Type returns the type string for the error
func (e *Error) Type() Type { return ERROR_OBJ }

// Function is the object that holds the reference to an executable function
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type returns the type string for the function
func (f *Function) Type() Type { return FUNCTION_OBJ }

// Inspect returns a string representing the function
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

// String is the object that holds strings
type String struct {
	Value string
}

// Inspect returns the value of the string
func (s *String) Inspect() string { return s.Value }

// Type returns the type string for the string
func (s *String) Type() Type { return STRING_OBJ }

// Builtin is the object that holds builtin functions
type Builtin struct {
	Fn BuiltinFunction
}

// Inspect returns a string for the builtin function
func (b *Builtin) Inspect() string { return "builtin function" }

// Type returns the type string for the builtin function
func (b *Builtin) Type() Type { return BUILTIN_OBJ }

// Array is the object that holds arrays
type Array struct {
	Elements []Object
}

// Inspect returns a string representing the array
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// Type returns the type string for the array
func (ao *Array) Type() Type { return ARRAY_OBJ }
