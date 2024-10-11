package analyzer

import (
	"loxcc/internal/analyzer/parser"
	"loxcc/internal/analyzer/scanner"
	"loxcc/internal/ast"
)

// Shorthand for calling scanner & parser.
func Analyze(source string) (ast.Program, error) {
	tokens, err := scanner.Scan(source)
	if err != nil {
		return nil, err
	}

	return parser.Parse(tokens)
}
