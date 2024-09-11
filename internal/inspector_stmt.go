package internal

import "loxcc/internal/ast"

func (a *astInspector) VisitExpressionStatement(s ast.ExpressionStatement) {
	a.indented("exprStmt", func() {
		s.Expr.Accept(a)
	})
}

func (a *astInspector) VisitForStatement(s ast.ForStatement) {
	a.indented("forStmt", func() {
		if s.Initializer != nil {
			switch s.Initializer.Kind {
			case ast.VarDecl:
				s.Initializer.VarInitializer.Accept(a)
			case ast.InitExpr:
				s.Initializer.ExprInitializer.Accept(a)
			}
		}

		if s.Condition != nil {
			s.Condition.Accept(a)
		}

		if s.Incrementer != nil {
			s.Incrementer.Accept(a)
		}

		s.Body.Accept(a)
	})
}

func (a *astInspector) VisitIfStatement(s ast.IfStatement)         {}
func (a *astInspector) VisitPrintStatement(s ast.PrintStatement)   {}
func (a *astInspector) VisitReturnStatement(s ast.ReturnStatement) {}
func (a *astInspector) VisitWhileStatement(s ast.WhileStatement)   {}
func (a *astInspector) VisitBlockStatement(s ast.BlockStatement)   {}
