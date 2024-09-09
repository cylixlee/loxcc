package ast

import stl "github.com/chen3feng/stl4go"

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
	Name      string
	Baseclass string
	Methods   stl.Vector[FunctionDeclaration]
}

type FunctionDeclaration struct {
	Name       string
	Parameters stl.Vector[string]
	Body       BlockStatement
}

type VarDeclaration struct {
	Name        string
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
