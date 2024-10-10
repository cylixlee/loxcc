package ast

import (
	"loxcc/internal/frontend/scanner"

	stl "github.com/chen3feng/stl4go"
)

// Lox Declarations.
//
// According to Appendix I of the book Crafting Interpreters, there's 4 kinds of
// declarations in Lox:
//   - Class declaration
//   - Function declaration
//   - Variable declaration
//   - Statement declaration
//
// All of which can be placed at top-level of a Lox program. In fact, a Lox program is a
// series of declarations.
//
// The Declaration is defined as an interface to achieve polymorphism, and the visitor
// pattern is adopted to separate the implementation of backend from the AST. There can be
// arbitrary number of declarations, but all of them should accept a DeclarationVisitor,
// and the visitor should be able to visit any kind of declaration.
type Declaration interface {
	Accept(DeclarationVisitor)
}

// The visitor of declarations.
//
// The visitor pattern is adopted to separate the implementation of code generation from
// the AST.
type DeclarationVisitor interface {
	VisitClassDeclaration(ClassDeclaration)
	VisitFunctionDeclaration(FunctionDeclaration)
	VisitVarDeclaration(VarDeclaration)
	VisitStatementDeclaration(StatementDeclaration)
}

// Lox class declaration.
//
// Lox supports classes and inheritance. After parsing, the name, baseclass and methods of
// a class can be known.
type ClassDeclaration struct {
	Name      *scanner.Token
	Baseclass *scanner.Token
	Methods   stl.Vector[Declaration]
}

// Lox function.
//
// Function is composed of a name, some parameters and the body. Note that Lox does not
// care about the types of parameters, it just record their names for matching arguments,
// and check the number of arguments passed in.
type FunctionDeclaration struct {
	Name       *scanner.Token
	Parameters stl.Vector[*scanner.Token]
	Body       Statement
}

// Lox variable declaration.
//
// Variables can be initialized (by an initializer expression) when declared. If not, the
// variable will be initialized as nil value.
type VarDeclaration struct {
	Name        *scanner.Token
	Initializer Expression
}

// Lox statement declaration.
//
// A statement can be parsed as a declaration, just like an expression with a trailing
// semicolon is parsed as an expression statement. This declaration makes statements able
// to be written at top-level (e.g. print something).
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
