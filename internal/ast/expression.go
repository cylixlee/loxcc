package ast

type Expression interface {
	Accept(ExpressionVisitor)
}

type ExpressionVisitor interface {
	VisitNumberLiteral(NumberLiteral)
	VisitBooleanLiteral(BooleanLiteral)
}

type NumberLiteral float64
type BooleanLiteral bool

// Visitor pattern implementations.

func (n NumberLiteral) Accept(visitor ExpressionVisitor)  { visitor.VisitNumberLiteral(n) }
func (b BooleanLiteral) Accept(visitor ExpressionVisitor) { visitor.VisitBooleanLiteral(b) }
