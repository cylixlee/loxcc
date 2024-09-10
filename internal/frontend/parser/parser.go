package parser

import (
	"loxcc/internal/frontend/scanner"

	stl "github.com/chen3feng/stl4go"
)

type Parser struct {
	tokens  stl.Vector[*scanner.Token]
	current int
}

func NewParser(tokens stl.Vector[*scanner.Token]) Parser {
	return Parser{
		tokens:  tokens,
		current: 0,
	}
}
