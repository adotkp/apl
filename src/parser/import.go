package parser

import (
	"ast/statement"
)

func (p *P) parseImports() ([]*statement.Import, error) {
	var ret []*statement.Import
	for {
		imp, err := p.parseImport()
		if err != nil {
			return nil, err
		}
		if imp == nil {
			return ret, nil
		}
		ret = append(ret, imp)
	}
}

func (p *P) parseImport() (*statement.Import, error) {
	mismatch, _, err := p.consume(TokenImport)
	if err != nil {
		if mismatch {
			p.tokens.unread()
			return nil, nil
		}
		return nil, err
	}
	name, _, err := p.consumeText()
	if err != nil {
		return nil, err
	}
	_, _, err = p.consume(TokenSemicolon)
	if err != nil {
		return nil, err
	}
	return &statement.Import{
		Name: name,
	}, nil
}
