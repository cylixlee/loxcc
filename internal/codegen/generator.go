package codegen

import "loxcc/internal/ast"

type CodeGenerator interface {
	ast.ExpressionVisitor
	ast.StatementVisitor
	ast.DeclarationVisitor

	Generate() string
}
