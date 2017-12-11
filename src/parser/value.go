package parser

import (
	"strconv"

	"values"
)

func (p *P) toValue(tok Token) (values.Value, error) {
	if tok.Typ == TokenString {
		return &values.String{V: string(tok.Lit)}, nil
	}
	if tok.Typ != TokenText {
		return nil, p.errf(tok, "expected constant value")
	}
	if '0' <= tok.Lit[0] && tok.Lit[0] <= '9' {
		v, err := strconv.Atoi(string(tok.Lit))
		if err != nil {
			return nil, err
		}
		return &values.Int{V: v}, nil
	}
	if string(tok.Lit) == "true" {
		return &values.Bool{V: true}, nil
	}
	if string(tok.Lit) == "false" {
		return &values.Bool{V: false}, nil
	}
	return nil, p.errf(tok, "could not convert to any value")
}
