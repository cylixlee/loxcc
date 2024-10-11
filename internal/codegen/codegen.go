package codegen

import (
	"loxcc/assets"
	"loxcc/internal/ast"

	stl "github.com/chen3feng/stl4go"
)

func Generate(program ast.Program) string {
	generator := newCodeGenerator()
	for _, decl := range program {
		decl.Accept(generator)
	}
	return assets.ApplyTemplate("entrypoint", generator)
}

type codeGenerator struct {
	Main stl.Vector[string]
	// visitor pattern does not support return values, so we have to store it in a stack.
	returnStack *stl.DList[string]
}

func newCodeGenerator() *codeGenerator {
	return &codeGenerator{
		Main:        stl.MakeVector[string](),
		returnStack: stl.NewDList[string](),
	}
}

func (c *codeGenerator) push(template string, data any) {
	c.returnStack.PushBack(assets.ApplyTemplate(template, data))
}

func (c *codeGenerator) pop() string {
	value, exists := c.returnStack.PopBack()
	if !exists {
		panic("stack underflow")
	}
	return value
}
