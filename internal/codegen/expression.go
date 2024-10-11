package codegen

import (
	"loxcc/internal/analyzer/scanner"
	"loxcc/internal/ast"
)

var (
	// Corresponding LOXCRT function calls to binary operators.
	//
	// Due to the dynamicity of Lox language, we could not use raw operators (e.g. +) to
	// evaluate Lox expressions in C. Those runtime functions are adopted to support
	// operations on Lox values.
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

	// Corresponding LOXCRT function calls to unary operators.
	//
	// Due to the dynamicity of Lox language, we could not use raw operators (e.g. +) to
	// evaluate Lox expressions in C. Those runtime functions are adopted to support
	// operations on Lox values.
	uopFuncMap = map[scanner.TokenType]string{
		scanner.TokMinus: "LRT_Negate",
		scanner.TokBang:  "LRT_Not",
	}
)

func (c *codeGenerator) VisitNilLiteral(n ast.NilLiteral)         { c.push("nil", nil) }
func (c *codeGenerator) VisitBooleanLiteral(b ast.BooleanLiteral) { c.push("boolean", b) }
func (c *codeGenerator) VisitNumberLiteral(n ast.NumberLiteral)   { c.push("number", n) }
func (c *codeGenerator) VisitStringLiteral(s ast.StringLiteral)   { c.push("string", s) }

func (c *codeGenerator) VisitIdentifierLiteral(i ast.IdentifierLiteral) { panic("unimplemented") }
func (c *codeGenerator) VisitThisLiteral(t ast.ThisLiteral)             { panic("unimplemented") }
func (c *codeGenerator) VisitSuperLiteral(s ast.SuperLiteral)           { panic("unimplemented") }

func (c *codeGenerator) VisitAssignmentExpression(a ast.AssignmentExpression) { panic("unimplemented") }

func (c *codeGenerator) VisitBinaryExpression(b ast.BinaryExpression) {
	operatorFunc, exists := binopFuncMap[b.Operator.Type]
	if !exists {
		panic("unrecognized operator " + b.Operator.Lexeme)
	}

	b.Left.Accept(c)
	b.Right.Accept(c)

	right, left := c.pop(), c.pop()
	c.push("binary", map[string]string{
		"left":     left,
		"right":    right,
		"operator": operatorFunc,
	})
}

func (c *codeGenerator) VisitUnaryExpression(u ast.UnaryExpression) {
	operatorFunc, exists := uopFuncMap[u.Operator.Type]
	if !exists {
		panic("unrecognized operator " + u.Operator.Lexeme)
	}

	u.Operand.Accept(c)
	c.push("unary", map[string]string{
		"operand":      c.pop(),
		"operatorFunc": operatorFunc,
	})
}

func (c *codeGenerator) VisitInvocationExpression(i ast.InvocationExpression) { panic("unimplemented") }
