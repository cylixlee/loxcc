package inspector

import (
	"fmt"
	"loxcc/internal/ast"
	"strings"

	stl "github.com/chen3feng/stl4go"
)

func (a *astInspector) VisitClassDeclaration(d ast.ClassDeclaration) {
	a.scope(fmt.Sprintf("ClassDecl $%s < $%s", d.Name.Lexeme, d.Baseclass.Lexeme), func() {
		for _, v := range d.Methods {
			v.Accept(a)
		}
	})
}

func (a *astInspector) VisitFunctionDeclaration(d ast.FunctionDeclaration) {
	var params stl.Vector[string]
	for _, v := range d.Parameters {
		params.PushBack(fmt.Sprint("$", v.Lexeme))
	}
	paramsSignature := strings.Join(params, ", ")
	signature := fmt.Sprintf("FunDecl $%s (%s)", d.Name.Lexeme, paramsSignature)

	a.scope(signature, func() {
		a.printf("body: ")
		d.Body.Accept(a)
	})
}

func (a *astInspector) VisitVarDeclaration(d ast.VarDeclaration) {
	a.scope(fmt.Sprintf("VarDecl $%s", d.Name.Lexeme), func() {
		if d.Initializer != nil {
			d.Initializer.Accept(a)
		}
	})
}

func (a *astInspector) VisitStatementDeclaration(d ast.StatementDeclaration) {
	d.Stmt.Accept(a)
}
