package types

import (
	"fmt"
)

// Context is the type registry.
type Context struct {
	m map[string]Type
}

// NewContext returns a new type registry with builtin types filled.
func NewContext() *Context {
	chk := func(err error) {
		if err != nil {
			panic(err)
		}
	}
	c := &Context{
		m: make(map[string]Type),
	}
	chk(c.Add("int", &Int{}))
	chk(c.Add("bool", &Bool{}))
	chk(c.Add("string", &String{}))
	return c
}

// Add adds the given type to the registry. Returns an error if it conflicts
// with an existing type.
func (c *Context) Add(name string, t Type) error {
	if prev, ok := c.m[name]; ok {
		return fmt.Errorf("type %s already declared as %v", name, prev)
	}
	c.m[name] = t
	return nil
}

// Get retrieves the type associated with the given name. If no type exists,
// returns ErrTypeNotFound.
func (c *Context) Get(name string) (Type, error) {
	if t, ok := c.m[name]; ok {
		return t, nil
	}
	return nil, fmt.Errorf("unknown type or func: %s", name)
}
