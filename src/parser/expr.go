package parser

import (
	"ast/expr"
)

func (p *P) parseExpr() (expr.Expr, error) {
	text, tok, err := p.consumeText()
	if tok.Typ == TokenComma || tok.Typ == TokenSemicolon || tok.Typ == TokenParensClose {
		p.tokens.unread()
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &expr.Value{
		Text: text,
	}, nil
}
