package ast

import (
	"loxcc/internal/analyzer/scanner"

	stl "github.com/chen3feng/stl4go"
)

// Lox expressions.
//
// Expression is defined as an interface for polymorphism, and visitor pattern is
// introduced to separate the code generation from AST.
type Expression interface {
	Accept(ExpressionVisitor)
}

// Visitor of Lox expressions.
type ExpressionVisitor interface {
	VisitNilLiteral(NilLiteral)
	VisitBooleanLiteral(BooleanLiteral)
	VisitNumberLiteral(NumberLiteral)
	VisitStringLiteral(StringLiteral)
	VisitIdentifierLiteral(IdentifierLiteral)
	VisitThisLiteral(ThisLiteral)
	VisitSuperLiteral(SuperLiteral)

	VisitAssignmentExpression(AssignmentExpression)
	VisitBinaryExpression(BinaryExpression)
	VisitUnaryExpression(UnaryExpression)
	VisitInvocationExpression(InvocationExpression)
}

// Literal expressions

type NilLiteral struct{}
type BooleanLiteral bool
type NumberLiteral float64
type StringLiteral string
type IdentifierLiteral string
type ThisLiteral struct{}
type SuperLiteral struct{}

// Structural expressions

type AssignmentExpression struct {
	Left  Expression
	Right Expression
}

type BinaryExpression struct {
	Left     Expression
	Operator *scanner.Token
	Right    Expression
}

type UnaryExpression struct {
	Operator *scanner.Token
	Operand  Expression
}

type InvocationExpression struct {
	Callee    Expression
	Arguments stl.Vector[Expression]
}

// Visitor pattern implementations.

func (n NilLiteral) Accept(visitor ExpressionVisitor)        { visitor.VisitNilLiteral(n) }
func (b BooleanLiteral) Accept(visitor ExpressionVisitor)    { visitor.VisitBooleanLiteral(b) }
func (n NumberLiteral) Accept(visitor ExpressionVisitor)     { visitor.VisitNumberLiteral(n) }
func (s StringLiteral) Accept(visitor ExpressionVisitor)     { visitor.VisitStringLiteral(s) }
func (i IdentifierLiteral) Accept(visitor ExpressionVisitor) { visitor.VisitIdentifierLiteral(i) }
func (t ThisLiteral) Accept(visitor ExpressionVisitor)       { visitor.VisitThisLiteral(t) }
func (s SuperLiteral) Accept(visitor ExpressionVisitor)      { visitor.VisitSuperLiteral(s) }

func (a AssignmentExpression) Accept(visitor ExpressionVisitor) { visitor.VisitAssignmentExpression(a) }
func (b BinaryExpression) Accept(visitor ExpressionVisitor)     { visitor.VisitBinaryExpression(b) }
func (u UnaryExpression) Accept(visitor ExpressionVisitor)      { visitor.VisitUnaryExpression(u) }
func (i InvocationExpression) Accept(visitor ExpressionVisitor) { visitor.VisitInvocationExpression(i) }
