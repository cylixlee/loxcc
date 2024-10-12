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

func (c *codeGenerator) VisitIfStatement(i ast.IfStatement) {
	i.Then.Accept(c)
	i.Condition.Accept(c)

	condition, then := c.pop(), c.pop()
	data := map[string]string{
		"condition": condition,
		"then":      then,
	}
	if i.Else != nil {
		i.Else.Accept(c)
		data["else"] = c.pop()
	}
	c.push("if", data)
}

func (c *codeGenerator) VisitPrintStatement(p ast.PrintStatement) {
	p.Value.Accept(c)
	c.push("print", c.pop())
}

func (c *codeGenerator) VisitReturnStatement(r ast.ReturnStatement) { panic("unimplemented") }

func (c *codeGenerator) VisitWhileStatement(w ast.WhileStatement) {
	w.Body.Accept(c)
	w.Condition.Accept(c)
	condition, body := c.pop(), c.pop()
	c.push("while", map[string]string{
		"condition": condition,
		"body":      body,
	})
}

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
