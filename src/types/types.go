package types

// Type represents a type in the language.
type Type interface {
}

// Int is an integer type.
type Int struct {
}

// Bool is a boolean type.
type Bool struct {
}

// String is a string type.
type String struct {
}

// Func is a function type.
type Func struct {
	Args   []Type
	Return Type
}
