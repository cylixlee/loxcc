package parser

import (
	"loxcc/internal/ast"
	"loxcc/internal/frontend/scanner"

	stl "github.com/chen3feng/stl4go"
)

func Parse(tokens stl.Vector[*scanner.Token]) (ast.Program, error) {
	var decls stl.Vector[ast.Declaration]
	p := parser{
		tokens:  tokens,
		current: 0,
	}

	for !p.hasReachedEnd() {
		decl, err := p.ParseDeclaration()
		if err != nil {
			return nil, err
		}
		decls.PushBack(decl)
	}
	return decls, nil
}

type parser struct {
	tokens  stl.Vector[*scanner.Token]
	current int
}

func (p parser) hasReachedEnd() bool { return p.current >= p.tokens.Len() }
