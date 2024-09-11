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
