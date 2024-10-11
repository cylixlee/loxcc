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

	if c.cascade == 0 { // global var
		c.GlobalVar.PushBack(map[string]string{
			"name":        v.Name.Lexeme,
			"initializer": initializer,
		})
	} else { // local var
		c.push("localvar", map[string]string{
			"name":        v.Name.Lexeme,
			"initializer": initializer,
		})
		// let block statement to pop the statement out
	}
}

func (c *codeGenerator) VisitStatementDeclaration(s ast.StatementDeclaration) {
	s.Stmt.Accept(c)
	if c.cascade == 0 {
		c.Main.PushBack(c.pop())
	}
}
