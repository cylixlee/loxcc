package ast

import (
	"loxcc/internal/frontend/scanner"

	stl "github.com/chen3feng/stl4go"
)

type Declaration interface {
	Accept(DeclarationVisitor)
}

type DeclarationVisitor interface {
	VisitClassDeclaration(ClassDeclaration)
	VisitFunctionDeclaration(FunctionDeclaration)
	VisitVarDeclaration(VarDeclaration)
	VisitStatementDeclaration(StatementDeclaration)
}

type ClassDeclaration struct {
	Name      *scanner.Token
	Baseclass *scanner.Token
	Methods   stl.Vector[Declaration]
}

type FunctionDeclaration struct {
	Name       *scanner.Token
	Parameters stl.Vector[*scanner.Token]
	Body       Statement
}

type VarDeclaration struct {
	Name        *scanner.Token
	Initializer Expression
}

type StatementDeclaration struct {
	Stmt Statement
}

// Visitor pattern implementations.
func (d ClassDeclaration) Accept(visitor DeclarationVisitor)    { visitor.VisitClassDeclaration(d) }
func (d FunctionDeclaration) Accept(visitor DeclarationVisitor) { visitor.VisitFunctionDeclaration(d) }
func (d VarDeclaration) Accept(visitor DeclarationVisitor)      { visitor.VisitVarDeclaration(d) }
func (d StatementDeclaration) Accept(visitor DeclarationVisitor) {
	visitor.VisitStatementDeclaration(d)
}
