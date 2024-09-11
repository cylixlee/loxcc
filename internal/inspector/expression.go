package inspector

import (
	"fmt"
	"loxcc/internal/ast"
)

func (a *astInspector) VisitNilLiteral(n ast.NilLiteral)               { a.printfln("nil") }
func (a *astInspector) VisitBooleanLiteral(b ast.BooleanLiteral)       { a.printfln("%v", b) }
func (a *astInspector) VisitNumberLiteral(n ast.NumberLiteral)         { a.printfln("%v", n) }
func (a *astInspector) VisitStringLiteral(s ast.StringLiteral)         { a.printfln("%v", s) }
func (a *astInspector) VisitIdentifierLiteral(i ast.IdentifierLiteral) { a.printfln("$%v", i) }
func (a *astInspector) VisitThisLiteral(t ast.ThisLiteral)             { a.printfln("this") }
func (a *astInspector) VisitSuperLiteral(s ast.SuperLiteral)           { a.printfln("super") }

func (a *astInspector) VisitAssignmentExpression(e ast.AssignmentExpression) {
	a.scope("AssignExpr", func() {
		a.printf("left: ")
		e.Left.Accept(a)
		a.printf("right: ")
		e.Right.Accept(a)
	})
}

func (a *astInspector) VisitBinaryExpression(e ast.BinaryExpression) {
	a.scope(fmt.Sprintf("BinaryExpr (%s)", e.Operator.Lexeme), func() {
		a.printf("left: ")
		e.Left.Accept(a)
		a.printf("right: ")
		e.Right.Accept(a)
	})
}

func (a *astInspector) VisitUnaryExpression(e ast.UnaryExpression) {
	a.scope(fmt.Sprintf("UnaryExpr (%s)", e.Operator.Lexeme), func() {
		a.printf("operand: ")
		e.Operand.Accept(a)
	})
}

func (a *astInspector) VisitInvocationExpression(e ast.InvocationExpression) {
	a.scope("InvokeExpr", func() {
		a.printf("callee: ")
		e.Callee.Accept(a)
		for idx, v := range e.Arguments {
			a.printf("arg[%d]: ", idx)
			v.Accept(a)
		}
	})
}
