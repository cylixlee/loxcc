package internal

import "loxcc/internal/ast"

func (a *astInspector) VisitClassDeclaration(d ast.ClassDeclaration)         {}
func (a *astInspector) VisitFunctionDeclaration(d ast.FunctionDeclaration)   {}
func (a *astInspector) VisitVarDeclaration(d ast.VarDeclaration)             {}
func (a *astInspector) VisitStatementDeclaration(d ast.StatementDeclaration) {}
