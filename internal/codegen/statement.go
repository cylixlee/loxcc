package codegen

import (
	"loxcc/internal/ast"

	stl "github.com/chen3feng/stl4go"
)

func (c *codeGenerator) VisitExpressionStatement(e ast.ExpressionStatement) {
	e.Expr.Accept(c)
	c.push("exprstmt", c.pop())
}

func (c *codeGenerator) VisitForStatement(f ast.ForStatement) { panic("unimplemented") }
func (c *codeGenerator) VisitIfStatement(i ast.IfStatement)   { panic("unimplemented") }

func (c *codeGenerator) VisitPrintStatement(p ast.PrintStatement) {
	p.Value.Accept(c)
	c.push("print", c.pop())
}

func (c *codeGenerator) VisitReturnStatement(r ast.ReturnStatement) { panic("unimplemented") }
func (c *codeGenerator) VisitWhileStatement(w ast.WhileStatement)   { panic("unimplemented") }

func (c *codeGenerator) VisitBlockStatement(b ast.BlockStatement) {
	c.cascade++
	contents := stl.MakeVector[string]()
	for _, decl := range b.Content {
		decl.Accept(c)
		contents.PushBack(c.pop())
	}
	c.push("block", contents)
	c.cascade--
}
