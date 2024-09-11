package inspector

import (
	"fmt"
	"loxcc/internal/ast"
)

func Inspect(expr ast.Declaration) {
	expr.Accept(new(astInspector))
}

type astInspector struct {
	indent   int
	indented bool
}

func (a *astInspector) scope(title string, f func()) {
	a.printfln("%s {", title)
	a.indent++
	f()
	a.indent--
	a.printfln("}")
}

func (a *astInspector) printf(format string, v ...any) {
	if !a.indented {
		for range a.indent {
			fmt.Print("  ")
		}
		a.indented = true
	}
	fmt.Printf(format, v...)
}

func (a *astInspector) println(v ...any) {
	if !a.indented {
		for range a.indent {
			fmt.Print("  ")
		}
		a.indented = true
	}
	fmt.Println(v...)
	a.indented = false
}

func (a *astInspector) printfln(format string, v ...any) {
	a.printf(format, v...)
	a.println()
}
