package parser

import (
	"loxcc/internal/ast"
	"loxcc/internal/frontend/scanner"

	stl "github.com/chen3feng/stl4go"
)

func (p *Parser) ParseDeclaration() (ast.Declaration, error) { panic("unimplemented") }

func (p *Parser) parseArguments() (stl.Vector[ast.Expression], error) {
	// create vector (slice)
	var arguments stl.Vector[ast.Expression]

	// parse the first argument
	expr, err := p.ParseExpression()
	if err != nil {
		return nil, err
	}
	arguments.PushBack(expr)

	// parse more arguments, each of which is lead by a [TokComma] (,).
	for p.tryConsume(scanner.TokComma) {
		expr, err = p.ParseExpression()
		if err != nil {
			return nil, err
		}
		arguments.PushBack(expr)
	}
	return arguments, nil
}
