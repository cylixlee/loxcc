package codegen

import "loxcc/internal/ast"

func (c *codeGenerator) VisitExpressionStatement(e ast.ExpressionStatement) {
	e.Expr.Accept(c)
	c.Main.PushBack(c.pop() + ";")
}

func (c *codeGenerator) VisitForStatement(f ast.ForStatement) { panic("unimplemented") }
func (c *codeGenerator) VisitIfStatement(i ast.IfStatement)   { panic("unimplemented") }

func (c *codeGenerator) VisitPrintStatement(p ast.PrintStatement) {
	p.Value.Accept(c)
	c.push("print", c.pop())
	c.Main.PushBack(c.pop())
}

func (c *codeGenerator) VisitReturnStatement(r ast.ReturnStatement) { panic("unimplemented") }
func (c *codeGenerator) VisitWhileStatement(w ast.WhileStatement)   { panic("unimplemented") }
func (c *codeGenerator) VisitBlockStatement(b ast.BlockStatement)   { panic("unimplemented") }
