package internal

import (
	"fmt"
	"loxcc/internal/ast"
)

func (a *astInspector) VisitNilLiteral(n ast.NilLiteral)               { a.printfln("nil") }
func (a *astInspector) VisitBooleanLiteral(b ast.BooleanLiteral)       { a.printfln("%v", b) }
func (a *astInspector) VisitNumberLiteral(n ast.NumberLiteral)         { a.printfln("%v", n) }
func (a *astInspector) VisitStringLiteral(s ast.StringLiteral)         { a.printfln("%v", s) }
func (a *astInspector) VisitIdentifierLiteral(i ast.IdentifierLiteral) { a.printfln("%v", i) }
func (a *astInspector) VisitThisLiteral(t ast.ThisLiteral)             { a.printfln("this") }
func (a *astInspector) VisitSuperLiteral(s ast.SuperLiteral)           { a.printfln("super") }

func (a *astInspector) VisitAssignmentExpression(e ast.AssignmentExpression) {
	a.indented("assign", func() {
		e.Left.Accept(a)
		e.Right.Accept(a)
	})
}

func (a *astInspector) VisitBinaryExpression(e ast.BinaryExpression) {
	a.indented(fmt.Sprintf("binary \"%s\"", e.Operator.Lexeme), func() {
		e.Left.Accept(a)
		e.Right.Accept(a)
	})
}

func (a *astInspector) VisitUnaryExpression(e ast.UnaryExpression) {
	a.indented(fmt.Sprintf("unary \"%s\"", e.Operator.Lexeme), func() {
		e.Operand.Accept(a)
	})
}

func (a *astInspector) VisitInvocationExpression(e ast.InvocationExpression) {
	a.indented("invoke", func() {
		e.Callee.Accept(a)
		for _, v := range e.Arguments {
			v.Accept(a)
		}
	})
}
