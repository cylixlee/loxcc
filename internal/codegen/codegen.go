package codegen

import (
	"loxcc/assets"

	stl "github.com/chen3feng/stl4go"
)

type codeGenerator struct {
	main stl.Vector[string]
	// visitor pattern does not support return values, so we have to store it in a stack.
	returnStack *stl.DList[string]
}

func newCodeGenerator() *codeGenerator {
	return &codeGenerator{
		main:        stl.MakeVector[string](),
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
