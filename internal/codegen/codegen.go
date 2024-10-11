package codegen

import stl "github.com/chen3feng/stl4go"

type codeGenerator struct {
	// visitor pattern does not support return values, so we have to store it in a stack.
	returnStack *stl.DList[string]
}

func newCodeGenerator() *codeGenerator {
	return &codeGenerator{
		returnStack: stl.NewDList[string](),
	}
}

func (c *codeGenerator) push(value string) { c.returnStack.PushBack(value) }
func (c *codeGenerator) pop(value string) string {
	value, exists := c.returnStack.PopBack()
	if !exists {
		panic("stack underflow")
	}
	return value
}
