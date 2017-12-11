package parser

import (
	"ast/expr"
)

// TODO(adi): Real expr parsing needs work.
func (p *P) parseExpr() (expr.Expr, error) {
	tok, err := p.tokens.get()
	if err != nil {
		return nil, err
	}
	if tok.Typ == TokenComma || tok.Typ == TokenSemicolon || tok.Typ == TokenParensClose {
		p.tokens.unread()
		return nil, nil
	}
	value, err := p.toValue(tok)
	if err != nil {
		return nil, err
	}
	return &expr.Value{
		Source: TokenSource{tok},
		V:      value,
	}, nil
}
