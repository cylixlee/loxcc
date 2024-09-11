package backend

import "loxcc/internal/ast"

func (c *codeGenerator) VisitExpressionStatement(s ast.ExpressionStatement) {
	s.Expr.Accept(c)
	c.write(";")
	c.push()
}

func (c *codeGenerator) VisitForStatement(s ast.ForStatement)       { panic("unimplemented") }
func (c *codeGenerator) VisitIfStatement(s ast.IfStatement)         { panic("unimplemented") }
func (c *codeGenerator) VisitPrintStatement(s ast.PrintStatement)   { panic("unimplemented") }
func (c *codeGenerator) VisitReturnStatement(s ast.ReturnStatement) { panic("unimplemented") }
func (c *codeGenerator) VisitWhileStatement(s ast.WhileStatement)   { panic("unimplemented") }
func (c *codeGenerator) VisitBlockStatement(s ast.BlockStatement)   { panic("unimplemented") }
