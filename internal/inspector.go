package internal

import (
	"fmt"
	"loxcc/internal/ast"
)

func Inspect(expr ast.Expression) {
	var inspector astInspector
	expr.Accept(&inspector)
}

type astInspector struct {
	indent int
}

func (a *astInspector) printfln(format string, v ...any) {
	for range a.indent {
		fmt.Print("  ")
	}
	fmt.Printf(format, v...)
	fmt.Println()
}

func (a *astInspector) VisitNilLiteral(n ast.NilLiteral)               { a.printfln("nil") }
func (a *astInspector) VisitBooleanLiteral(b ast.BooleanLiteral)       { a.printfln("%v", b) }
func (a *astInspector) VisitNumberLiteral(n ast.NumberLiteral)         { a.printfln("%v", n) }
func (a *astInspector) VisitStringLiteral(s ast.StringLiteral)         { a.printfln("%v", s) }
func (a *astInspector) VisitIdentifierLiteral(i ast.IdentifierLiteral) { a.printfln("%v", i) }
func (a *astInspector) VisitThisLiteral(t ast.ThisLiteral)             { a.printfln("this") }
func (a *astInspector) VisitSuperLiteral(s ast.SuperLiteral)           { a.printfln("super") }

func (a *astInspector) VisitAssignmentExpression(e ast.AssignmentExpression) {
	a.printfln("<assignment>")
	a.indent++
	e.Left.Accept(a)
	e.Right.Accept(a)
	a.indent--
}

func (a *astInspector) VisitBinaryExpression(e ast.BinaryExpression) {
	a.printfln("<binop %s>", e.Operator.Lexeme)
	a.indent++
	e.Left.Accept(a)
	e.Right.Accept(a)
	a.indent--
}

func (a *astInspector) VisitUnaryExpression(e ast.UnaryExpression) {
	a.printfln("<uop %s>", e.Operator.Lexeme)
	a.indent++
	e.Operand.Accept(a)
	a.indent--
}

func (a *astInspector) VisitInvocationExpression(e ast.InvocationExpression) {
	a.printfln("<invoke>")
	e.Callee.Accept(a)
	a.indent++
	for _, v := range e.Arguments {
		v.Accept(a)
	}
	a.indent--
}
