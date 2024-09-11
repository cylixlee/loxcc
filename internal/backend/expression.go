package backend

import "loxcc/internal/ast"

func (c *codeGenerator) VisitNilLiteral(n ast.NilLiteral)         { panic("unimplemented") }
func (c *codeGenerator) VisitBooleanLiteral(b ast.BooleanLiteral) { panic("unimplemented") }

func (c *codeGenerator) VisitNumberLiteral(n ast.NumberLiteral) {
	c.writef("%v", n)
}

func (c *codeGenerator) VisitStringLiteral(s ast.StringLiteral)         { panic("unimplemented") }
func (c *codeGenerator) VisitIdentifierLiteral(i ast.IdentifierLiteral) { panic("unimplemented") }
func (c *codeGenerator) VisitThisLiteral(t ast.ThisLiteral)             { panic("unimplemented") }
func (c *codeGenerator) VisitSuperLiteral(s ast.SuperLiteral)           { panic("unimplemented") }

func (c *codeGenerator) VisitAssignmentExpression(e ast.AssignmentExpression) { panic("unimplemented") }

func (c *codeGenerator) VisitBinaryExpression(e ast.BinaryExpression) {
	// add parentheses to make sure operator precedence.
	c.write("(")
	defer c.write(")")

	e.Left.Accept(c)
	c.write(e.Operator.Lexeme)
	e.Right.Accept(c)
}

func (c *codeGenerator) VisitUnaryExpression(e ast.UnaryExpression) {
	c.write("(")
	defer c.write(")")

	c.write(e.Operator.Lexeme)
	e.Operand.Accept(c)
}

func (c *codeGenerator) VisitInvocationExpression(e ast.InvocationExpression) { panic("unimplemented") }
