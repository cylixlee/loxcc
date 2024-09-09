package internal

import (
	"loxcc/internal/ast"

	stl "github.com/chen3feng/stl4go"
)

// Lox programs consist of declarations.
type Program struct {
	definitions stl.Vector[ast.Declaration]
}
