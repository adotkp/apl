package parser

import (
	"ast/expr"
	"ast/statement"
)

func (p *P) parseStatements() ([]statement.Statement, error) {
	var stmts []statement.Statement
	for {
		stmt, err := p.parseStatement()
		if err != nil {
			return nil, err
		}
		if stmt == nil {
			return stmts, nil
		}
		stmts = append(stmts, stmt)
	}
}

func (p *P) parseStatement() (statement.Statement, error) {
	tok, err := p.tokens.get()
	if tok.Typ == TokenBraceClose {
		p.tokens.unread()
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if tok.Typ == TokenReturn {
		p.tokens.unread()
		return p.parseReturnStmt()
	}
	if tok.Typ != TokenText {
		return nil, p.errf(tok, "expected identifier")
	}
	next, err := p.tokens.get()
	if err != nil {
		return nil, err
	}
	p.tokens.unread()
	if next.Typ == TokenParensOpen {
		return p.parseFnCall(string(tok.Lit), TokenSource{tok})
	}
	if next.Typ == TokenAssign {
		panic("assignment not yet implemented")
	}
	panic("parse error - did you try to declare a value?")
}

func (p *P) parseReturnStmt() (*statement.Return, error) {
	_, tok, err := p.consume(TokenReturn)
	if err != nil {
		return nil, err
	}
	expr, err := p.parseExpr()
	if err != nil {
		return nil, err
	}
	_, _, err = p.consume(TokenSemicolon)
	if err != nil {
		return nil, err
	}
	return &statement.Return{
		Source: TokenSource{tok},
		Expr:   expr,
	}, nil
}

func (p *P) parseFnCall(name string, src TokenSource) (*statement.FnCall, error) {
	_, _, err := p.consume(TokenParensOpen)
	if err != nil {
		return nil, err
	}
	var params []expr.Expr
	for {
		expr, err := p.parseExpr()
		if err != nil {
			return nil, err
		}
		if expr == nil {
			break
		}
		params = append(params, expr)
	}
	_, _, err = p.consume(TokenParensClose)
	if err != nil {
		return nil, err
	}
	_, _, err = p.consume(TokenSemicolon)
	if err != nil {
		return nil, err
	}
	return &statement.FnCall{
		Source: src,
		Nam:    name,
		Params: params,
	}, nil
}
