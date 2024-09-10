package parser

import (
	"loxcc/internal/ast"
	"loxcc/internal/frontend/scanner"

	stl "github.com/chen3feng/stl4go"
)

func (p *Parser) parseArguments(terminator scanner.TokenType) (stl.Vector[ast.Expression], error) {
	// create vector (slice)
	var arguments stl.Vector[ast.Expression]

	// handle the situation of 0 argument(s)
	if p.tryConsume(terminator) {
		return arguments, nil
	}

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
