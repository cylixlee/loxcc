package internal

import "loxcc/internal/ast"

func (a *astInspector) VisitExpressionStatement(s ast.ExpressionStatement) {
	s.Expr.Accept(a)
}

func (a *astInspector) VisitForStatement(s ast.ForStatement) {
	a.scope("ForStmt", func() {
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
	a.scope("IfStmt", func() {
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
	a.scope("PrintStmt", func() {
		a.printf("expr: ")
		s.Value.Accept(a)
	})
}

func (a *astInspector) VisitReturnStatement(s ast.ReturnStatement) {
	a.scope("ReturnStmt", func() {
		a.printf("expr: ")
		s.Value.Accept(a)
	})
}

func (a *astInspector) VisitWhileStatement(s ast.WhileStatement) {
	a.scope("WhileStmt", func() {
		a.printf("condition: ")
		s.Condition.Accept(a)
		a.printf("body: ")
		s.Body.Accept(a)
	})
}

func (a *astInspector) VisitBlockStatement(s ast.BlockStatement) {
	a.scope("BlockStmt", func() {
		for _, v := range s.Content {
			v.Accept(a)
		}
	})
}
