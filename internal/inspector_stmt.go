package internal

import "loxcc/internal/ast"

func (a *astInspector) VisitExpressionStatement(s ast.ExpressionStatement) {
	a.scope("exprStmt", func() {
		s.Expr.Accept(a)
	})
}

func (a *astInspector) VisitForStatement(s ast.ForStatement) {
	a.scope("forStmt", func() {
		if s.Initializer != nil {
			a.printf("initializer: ")
			switch s.Initializer.Kind {
			case ast.VarDecl:
				s.Initializer.VarInitializer.Accept(a)
			case ast.InitExpr:
				s.Initializer.ExprInitializer.Accept(a)
			}
		}

		if s.Condition != nil {
			a.printf("condition: ")
			s.Condition.Accept(a)
		}

		if s.Incrementer != nil {
			a.printf("incrementer: ")
			s.Incrementer.Accept(a)
		}

		a.printf("body: ")
		s.Body.Accept(a)
	})
}

func (a *astInspector) VisitIfStatement(s ast.IfStatement) {
	a.scope("ifStmt", func() {
		a.printf("condition: ")
		s.Condition.Accept(a)
		a.printf("then: ")
		s.Then.Accept(a)
		if s.Else != nil {
			a.printf("else: ")
			s.Else.Accept(a)
		}
	})
}

func (a *astInspector) VisitPrintStatement(s ast.PrintStatement) {
	a.scope("printStmt", func() {
		a.printf("expr: ")
		s.Value.Accept(a)
	})
}

func (a *astInspector) VisitReturnStatement(s ast.ReturnStatement) {
	a.scope("returnStmt", func() {
		a.printf("expr: ")
		s.Value.Accept(a)
	})
}

func (a *astInspector) VisitWhileStatement(s ast.WhileStatement) {
	a.scope("while", func() {
		a.printf("condition: ")
		s.Condition.Accept(a)
		a.printf("body: ")
		s.Body.Accept(a)
	})
}

func (a *astInspector) VisitBlockStatement(s ast.BlockStatement) {
	a.scope("blockStmt", func() {
		for _, v := range s.Content {
			v.Accept(a)
		}
	})
}
