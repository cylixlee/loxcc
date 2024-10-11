package codegen

import "loxcc/internal/ast"

func (c *codeGenerator) VisitClassDeclaration(k ast.ClassDeclaration)       { panic("unimplemented") }
func (c *codeGenerator) VisitFunctionDeclaration(f ast.FunctionDeclaration) { panic("unimplemented") }

func (c *codeGenerator) VisitVarDeclaration(v ast.VarDeclaration) {
	if v.Initializer != nil {
		v.Initializer.Accept(c)
	} else {
		c.push("nil", nil)
	}
	initializer := c.pop()

	c.VarDecl.PushBack(map[string]string{
		"name":        v.Name.Lexeme,
		"initializer": initializer,
	})
}

func (c *codeGenerator) VisitStatementDeclaration(s ast.StatementDeclaration) {
	s.Stmt.Accept(c)
}
