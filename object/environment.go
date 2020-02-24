package object

// NewEnvironment creates a newly scoped environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

// Environment keeps track of identifiers and their values
type Environment struct {
	store map[string]Object
}

// Get will return the value for an identifier if it exists
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

// Set will save a new identifier and value to the environment
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
