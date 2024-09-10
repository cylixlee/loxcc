package backend

import "loxcc/internal/ast"

type CodeGenerator interface {
	Generate(ast.Program) string
}
