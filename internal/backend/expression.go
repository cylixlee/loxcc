package backend

import (
	"loxcc/internal/ast"
	"loxcc/internal/frontend/scanner"
)

var (
	binopFuncMap = map[scanner.TokenType]string{
		scanner.TokPlus:         "LRT_Add",
		scanner.TokMinus:        "LRT_Subtract",
		scanner.TokStar:         "LRT_Multiply",
		scanner.TokSlash:        "LRT_Divide",
		scanner.TokEqualEqual:   "LRT_Equal",
		scanner.TokGreater:      "LRT_Greater",
		scanner.TokLess:         "LRT_Less",
		scanner.TokBangEqual:    "LRT_NotEqual",
		scanner.TokLessEqual:    "LRT_LessEqual",
		scanner.TokGreaterEqual: "LRT_GreaterEqual",
	}

	uopFuncMap = map[scanner.TokenType]string{
		scanner.TokMinus: "LRT_Negate",
		scanner.TokBang:  "LRT_Not",
	}
)

func (c *codeGenerator) VisitNilLiteral(n ast.NilLiteral) { c.write("NIL") }

func (c *codeGenerator) VisitBooleanLiteral(b ast.BooleanLiteral) {
	c.writef("BOOLEAN(%v)", b)
}

func (c *codeGenerator) VisitNumberLiteral(n ast.NumberLiteral) {
	c.writef("NUMBER(%v)", n)
}

func (c *codeGenerator) VisitStringLiteral(s ast.StringLiteral) {
	c.writef("OBJECT(LRT_NewString(%v, %v))", s, len(s)-2)
}

func (c *codeGenerator) VisitIdentifierLiteral(i ast.IdentifierLiteral) { panic("unimplemented") }
func (c *codeGenerator) VisitThisLiteral(t ast.ThisLiteral)             { panic("unimplemented") }
func (c *codeGenerator) VisitSuperLiteral(s ast.SuperLiteral)           { panic("unimplemented") }

func (c *codeGenerator) VisitAssignmentExpression(e ast.AssignmentExpression) { panic("unimplemented") }

func (c *codeGenerator) VisitBinaryExpression(e ast.BinaryExpression) {
	// add parentheses to make sure operator precedence.
	c.write("(")
	defer c.write(")")

	if f, exists := binopFuncMap[e.Operator.Type]; exists {
		c.write(f)
	} else {
		panic("unrecognized binary operator")
	}

	c.write("(")
	defer c.write(")")

	e.Left.Accept(c)
	c.write(", ")
	e.Right.Accept(c)
}

func (c *codeGenerator) VisitUnaryExpression(e ast.UnaryExpression) {
	c.write("(")
	defer c.write(")")

	if f, exists := uopFuncMap[e.Operator.Type]; exists {
		c.write(f)
	} else {
		panic("unrecognized unary operator")
	}

	c.write("(")
	defer c.write(")")

	e.Operand.Accept(c)
}

func (c *codeGenerator) VisitInvocationExpression(e ast.InvocationExpression) {
	panic("unimplemented")
}
