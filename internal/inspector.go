package internal

import (
	"fmt"
	"loxcc/internal/ast"
)

func Inspect(expr ast.Expression) {
	expr.Accept(new(astInspector))
}

type astInspector struct {
	indent int
}

func (a *astInspector) indented(title string, f func()) {
	a.printfln("<%s>", title)
	a.indent++
	f()
	a.indent--
}

func (a *astInspector) printfln(format string, v ...any) {
	for range a.indent {
		fmt.Print("\t")
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

func (a *astInspector) VisitAssignmentExpression(as ast.AssignmentExpression) {
	a.indented("assign", func() {
		as.Left.Accept(a)
		as.Right.Accept(a)
	})
}

func (a *astInspector) VisitBinaryExpression(b ast.BinaryExpression) {
	a.indented(fmt.Sprintf("binary \"%s\"", b.Operator.Lexeme), func() {
		b.Left.Accept(a)
		b.Right.Accept(a)
	})
}

func (a *astInspector) VisitUnaryExpression(u ast.UnaryExpression) {
	a.indented(fmt.Sprintf("unary \"%s\"", u.Operator.Lexeme), func() {
		u.Operand.Accept(a)
	})
}

func (a *astInspector) VisitInvocationExpression(i ast.InvocationExpression) {
	a.indented("invoke", func() {
		i.Callee.Accept(a)
		for _, v := range i.Arguments {
			v.Accept(a)
		}
	})
}
