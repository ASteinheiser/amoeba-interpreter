package object

import "fmt"

// Type is the type of an object
type Type string

const (
	// INTEGER_OBJ is the object type for integers
	INTEGER_OBJ = "INTEGER"
	// BOOLEAN_OBJ is the object type for booleans
	BOOLEAN_OBJ = "BOOLEAN"
	// NULL_OBJ is the object type for nulls
	NULL_OBJ = "NULL"
)

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
