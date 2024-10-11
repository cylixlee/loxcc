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
	GlobalVar stl.Vector[map[string]string]
	Main      stl.Vector[string]

	// visitor pattern does not support return values, so we have to store it in a stack.
	// moreover, some multi-step operations may need a stack for help.
	operationStack *stl.DList[string]
	cascade        int
}

func newCodeGenerator() *codeGenerator {
	return &codeGenerator{
		GlobalVar:      stl.MakeVector[map[string]string](),
		Main:           stl.MakeVector[string](),
		operationStack: stl.NewDList[string](),
	}
}

func (c *codeGenerator) push(template string, data any) {
	c.operationStack.PushBack(assets.ApplyTemplate(template, data))
}

func (c *codeGenerator) pop() string {
	value, exists := c.operationStack.PopBack()
	if !exists {
		panic("stack underflow")
	}
	return value
}
