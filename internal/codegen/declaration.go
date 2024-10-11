package codegen

import "loxcc/internal/ast"

func (c *codeGenerator) VisitClassDeclaration(k ast.ClassDeclaration)       { panic("unimplemented") }
func (c *codeGenerator) VisitFunctionDeclaration(f ast.FunctionDeclaration) { panic("unimplemented") }
func (c *codeGenerator) VisitVarDeclaration(v ast.VarDeclaration)           { panic("unimplemented") }

func (c *codeGenerator) VisitStatementDeclaration(s ast.StatementDeclaration) {
	s.Stmt.Accept(c)
}
