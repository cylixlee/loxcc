package codegen

import (
	"loxcc/internal/ast"

	stl "github.com/chen3feng/stl4go"
)

func (c *codeGenerator) VisitExpressionStatement(e ast.ExpressionStatement) {
	e.Expr.Accept(c)
	c.push("exprstmt", c.pop())
}

func (c *codeGenerator) VisitForStatement(f ast.ForStatement) {
	c.cascade++
	f.Body.Accept(c)
	data := map[string]string{
		"body": c.pop(),
	}
	if f.Initializer != nil {
		switch f.Initializer.Kind {
		case ast.VarDecl:
			f.Initializer.VarInitializer.Accept(c)
		case ast.InitExpr:
			f.Initializer.ExprInitializer.Accept(c)
		default:
			panic("unreachable code in parsing for-loop initializer.")
		}
		data["initializer"] = c.pop()
	}

	if f.Condition != nil {
		f.Condition.Accept(c)
		data["condition"] = c.pop()
	}

	if f.Incrementer != nil {
		f.Incrementer.Accept(c)
		data["incrementer"] = c.pop()
	}
	c.push("for", data)
	c.cascade--
}

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
