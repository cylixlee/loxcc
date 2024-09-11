package backend

import "loxcc/internal/ast"

func (c *codeGenerator) VisitClassDeclaration(d ast.ClassDeclaration)       { panic("unimplemented") }
func (c *codeGenerator) VisitFunctionDeclaration(d ast.FunctionDeclaration) { panic("unimplemented") }
func (c *codeGenerator) VisitVarDeclaration(d ast.VarDeclaration)           { panic("unimplemented") }

func (c *codeGenerator) VisitStatementDeclaration(d ast.StatementDeclaration) {
	d.Stmt.Accept(c)
}
