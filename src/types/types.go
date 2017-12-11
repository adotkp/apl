package types

type Type interface {
}

type Int struct {
}

type Bool struct {
}

type String struct {
}

type Func struct {
	Args   []Type
	Return Type
}
