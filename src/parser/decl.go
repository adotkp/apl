package parser

import (
	"ast"
)

func (p *P) parseDecls() ([]ast.Decl, error) {
	var decls []ast.Decl
	for {
		decl, err := p.parseDecl()
		if err == errEof {
			return decls, nil
		}
		if err != nil {
			return nil, err
		}
		decls = append(decls, decl)
	}
}

func (p *P) parseDecl() (ast.Decl, error) {
	_, tok, err := p.consume(TokenFunc, TokenTyp)
	if err != nil {
		return nil, err
	}
	switch tok.Typ {
	case TokenFunc:
		p.tokens.unread()
		return p.parseFnDecl()
	case TokenTyp:
		p.tokens.unread()
		panic("type decl not yet implemented")
	default:
		return nil, p.errf(tok, "unexpected keyword")
	}
}

func (p *P) parseFnDecl() (ast.Decl, error) {
	_, tok, err := p.consume(TokenFunc)
	if err != nil {
		return nil, err
	}
	name, _, err := p.consumeText()
	if err != nil {
		return nil, err
	}
	args, err := p.parseFnArgs()
	if err != nil {
		return nil, err
	}
	returnType, err := p.parseFnReturn()
	if err != nil {
		return nil, err
	}
	_, _, err = p.consume(TokenBraceOpen)
	if err != nil {
		return nil, err
	}
	stmts, err := p.parseStatements()
	if err != nil {
		return nil, err
	}
	_, _, err = p.consume(TokenBraceClose)
	if err != nil {
		return nil, err
	}
	return &ast.FnDecl{
		Source:     TokenSource{tok},
		Nam:        name,
		Args:       args,
		Return:     returnType,
		Statements: stmts,
	}, nil
}

func (p *P) parseFnArgs() ([]*ast.FnArg, error) {
	_, _, err := p.consume(TokenParensOpen)
	if err != nil {
		return nil, err
	}
	var args []*ast.FnArg
	for {
		typ, tok, err := p.consumeText()
		if tok.Typ == TokenParensClose {
			break
		}
		if err != nil {
			return nil, err
		}
		name, _, err := p.consumeText()
		if err != nil {
			return nil, err
		}
		args = append(args, &ast.FnArg{
			Source: TokenSource{tok},
			Nam:    name,
			Typ:    typ,
		})
		_, tok, err = p.consume(TokenComma)
		if tok.Typ == TokenParensClose {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	return args, nil
}

func (p *P) parseFnReturn() (*ast.FnReturn, error) {
	typ, tok, err := p.consumeText()
	if tok.Typ == TokenBraceOpen {
		p.tokens.unread()
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &ast.FnReturn{Source: TokenSource{tok}, Typ: typ}, nil
}
