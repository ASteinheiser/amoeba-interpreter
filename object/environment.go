package object

// NewEnvironment creates a newly scoped environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// NewEnclosedEnvironment creates a newly scoped environment with a reference to the outer env
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// Environment keeps track of identifiers and their values
type Environment struct {
	store map[string]Object
	outer *Environment
}

// Get will return the value for an identifier if it exists
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set will save a new identifier and value to the environment
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
