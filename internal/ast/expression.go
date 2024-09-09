package ast

type Expression interface {
	Accept(ExpressionVisitor)
}

type ExpressionVisitor interface {
	VisitNumberLiteral(NumberLiteral)
}

type NumberLiteral float64

// Visitor pattern implementations.

func (n NumberLiteral) Accept(visitor ExpressionVisitor) { visitor.VisitNumberLiteral(n) }
