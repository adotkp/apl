package types

// Type represents a type in the language.
type Type interface {
	Equals(Type) bool
}

// Int is an integer type.
type Int struct {
}

// Equals returns true if t is Int.
func (i *Int) Equals(t Type) bool {
	_, ok := t.(*Int)
	return ok
}

func (i *Int) String() string {
	return "type<int>"
}

// Bool is a boolean type.
type Bool struct {
}

// Equals returns true if t is Bool.
func (b *Bool) Equals(t Type) bool {
	_, ok := t.(*Bool)
	return ok
}

func (b *Bool) String() string {
	return "type<bool>"
}

// String is a string type.
type String struct {
}

// Equals returns true if t is String.
func (s *String) Equals(t Type) bool {
	_, ok := t.(*String)
	return ok
}

func (s *String) String() string {
	return "type<string>"
}

// Func is a function type.
type Func struct {
	Args   []Type
	Return Type
}

// Equals returns true if t is Func with the same args and return type.
func (f *Func) Equals(t Type) bool {
	f2, ok := t.(*Func)
	if !ok {
		return false
	}
	if len(f.Args) != len(f2.Args) {
		return false
	}
	for i, arg := range f.Args {
		if !arg.Equals(f2.Args[i]) {
			return false
		}
	}
	return f.Return.Equals(f2.Return)
}

func (f *Func) String() string {
	return "type<func>"
}
