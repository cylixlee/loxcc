package ast

import (
	stl "github.com/chen3feng/stl4go"
)

const (
	VarDecl ForInitializerKind = iota
	InitExpr
)

type ForInitializerKind byte

type Statement interface {
	Accept(StatementVisitor)
}

type StatementVisitor interface {
	VisitExpressionStatement(ExpressionStatement)
	VisitForStatement(ForStatement)
	VisitIfStatement(IfStatement)
	VisitPrintStatement(PrintStatement)
	VisitReturnStatement(ReturnStatement)
	VisitWhileStatement(WhileStatement)
	VisitBlockStatement(BlockStatement)
}

type ExpressionStatement struct {
	Expr Expression
}

type ForLoopInitializer struct {
	Kind            ForInitializerKind
	VarInitializer  Declaration
	ExprInitializer Expression
}

type ForStatement struct {
	Initializer *ForLoopInitializer
	Condition   Expression
	Incrementer Expression
	Body        Statement
}

type IfStatement struct {
	Condition Expression
	Then      Statement
	Else      Statement
}

type PrintStatement struct {
	Value Expression
}

type ReturnStatement struct {
	Value Expression
}

type WhileStatement struct {
	Condition Expression
	Body      Statement
}

type BlockStatement struct {
	Content stl.Vector[Declaration]
}

// Visitor pattern implementations.

func (s ExpressionStatement) Accept(visitor StatementVisitor) { visitor.VisitExpressionStatement(s) }
func (s ForStatement) Accept(visitor StatementVisitor)        { visitor.VisitForStatement(s) }
func (s IfStatement) Accept(visitor StatementVisitor)         { visitor.VisitIfStatement(s) }
func (s PrintStatement) Accept(visitor StatementVisitor)      { visitor.VisitPrintStatement(s) }
func (s ReturnStatement) Accept(visitor StatementVisitor)     { visitor.VisitReturnStatement(s) }
func (s WhileStatement) Accept(visitor StatementVisitor)      { visitor.VisitWhileStatement(s) }
func (s BlockStatement) Accept(visitor StatementVisitor)      { visitor.VisitBlockStatement(s) }
