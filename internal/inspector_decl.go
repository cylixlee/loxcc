package internal

import (
	"fmt"
	"loxcc/internal/ast"
)

func (a *astInspector) VisitClassDeclaration(d ast.ClassDeclaration) {
	a.scope(fmt.Sprintf("classDecl %s < %s", d.Name.Lexeme, d.Baseclass.Lexeme), func() {
		for _, v := range d.Methods {
			v.Accept(a)
		}
	})
}

func (a *astInspector) VisitFunctionDeclaration(d ast.FunctionDeclaration) {
	signature := fmt.Sprintf("funDecl %s (", d.Name.Lexeme)
	for _, v := range d.Parameters {
		signature = fmt.Sprint(signature, v.Lexeme)
	}
	signature = fmt.Sprint(signature, ")")

	a.scope(signature, func() {
		d.Body.Accept(a)
	})
}

func (a *astInspector) VisitVarDeclaration(d ast.VarDeclaration) {
	a.scope(fmt.Sprintf("varDecl %s", d.Name.Lexeme), func() {
		if d.Initializer != nil {
			d.Initializer.Accept(a)
		}
	})
}

func (a *astInspector) VisitStatementDeclaration(d ast.StatementDeclaration) {
	a.scope("stmtDecl", func() {
		d.Stmt.Accept(a)
	})
}
