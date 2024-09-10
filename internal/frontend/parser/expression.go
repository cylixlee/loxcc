package parser

import (
	"errors"
	"loxcc/internal/ast"
	"loxcc/internal/frontend/scanner"
	"strconv"
)

const (
	None Precedence = iota
	Assignment
	ConditionalOr
	ConditionalAnd
	Equality
	Relational
	Additive
	Multiplicative
	Invocation
	Property
	Impossible
)

var (
	precedenceMap = map[scanner.TokenType]Precedence{
		scanner.TokEqual:           Assignment,
		scanner.TokOr:              ConditionalOr,
		scanner.TokAnd:             ConditionalAnd,
		scanner.TokEqualEqual:      Equality,
		scanner.TokBangEqual:       Equality,
		scanner.TokGreater:         Relational,
		scanner.TokLess:            Relational,
		scanner.TokGreaterEqual:    Relational,
		scanner.TokLessEqual:       Relational,
		scanner.TokPlus:            Additive,
		scanner.TokMinus:           Additive,
		scanner.TokStar:            Multiplicative,
		scanner.TokSlash:           Multiplicative,
		scanner.TokLeftParenthesis: Invocation,
		scanner.TokDot:             Property,
	}

	ErrInvalidPrefix = errors.New("invalid prefix expression")
)

type Precedence byte

func (p Precedence) increase() Precedence {
	if p >= Impossible {
		return Impossible
	}
	return p + 1
}

func precedenceOf(t scanner.TokenType) Precedence {
	if p, exists := precedenceMap[t]; exists {
		return p
	}
	return None
}

func (p *Parser) ParseExpression() (ast.Expression, error) {
	return p.parsePrecedence(Assignment)
}

func (p *Parser) parsePrecedence(precedence Precedence) (ast.Expression, error) {
	// consume prefix token
	prefix, err := p.mustAdvance()
	if err != nil {
		return nil, err
	}

	// parse prefix expression
	var expr ast.Expression

	switch prefix.Type {
	case scanner.TokNil:
		expr = ast.NilLiteral{}
	case scanner.TokTrue:
		expr = ast.BooleanLiteral(true)
	case scanner.TokFalse:
		expr = ast.BooleanLiteral(false)
	case scanner.TokNumber:
		number, err := strconv.ParseFloat(prefix.Lexeme, 64)
		if err != nil {
			return nil, err
		}
		expr = ast.NumberLiteral(number)
	case scanner.TokString:
		expr = ast.StringLiteral(prefix.Lexeme)
	case scanner.TokIdentifier:
		expr = ast.IdentifierLiteral(prefix.Lexeme)
	case scanner.TokThis:
		expr = ast.ThisLiteral{}
	case scanner.TokSuper:
		expr = ast.SuperLiteral{}
	case scanner.TokLeftParenthesis:
		expr, err = p.parseParenthesized()
		if err != nil {
			return nil, err
		}
	case scanner.TokBang, scanner.TokMinus:
		expr, err = p.parseUnary()
		if err != nil {
			return nil, err
		}
	default:
		return nil, ErrInvalidPrefix
	}

	for {
		infix := p.peek()
		if infix == nil || precedenceOf(infix.Type) < precedence {
			break
		}

		switch infix.Type {
		case scanner.TokLeftParenthesis:
			expr, err = p.parseInvocation(expr)
			if err != nil {
				return nil, err
			}
		case scanner.TokEqual:
			expr, err = p.parseAssignment(expr)
			if err != nil {
				return nil, err
			}
		default:
			expr, err = p.parseBinary(expr)
			if err != nil {
				return nil, err
			}
		}
	}
	return expr, nil
}

func (p *Parser) parseParenthesized() (ast.Expression, error) {
	// // consume left parenthesis
	// if _, err := p.mustConsume(scanner.TokLeftParenthesis); err != nil {
	// 	return nil, err
	// }

	// parse expression
	expr, err := p.ParseExpression()
	if err != nil {
		return nil, err
	}

	// consume right parenthesis
	if _, err := p.mustConsume(scanner.TokRightParenthesis); err != nil {
		return nil, err
	}
	return expr, nil
}

func (p *Parser) parseInvocation(callee ast.Expression) (ast.Expression, error) {
	// consume left parenthesis
	if _, err := p.mustConsume(scanner.TokLeftParenthesis); err != nil {
		return nil, err
	}

	// parse arguments
	arguments, err := p.parseArguments(scanner.TokRightParenthesis)
	if err != nil {
		return nil, err
	}

	// parse right parenthesis
	if _, err := p.mustConsume(scanner.TokRightParenthesis); err != nil {
		return nil, err
	}
	return ast.InvocationExpression{
		Callee:    callee,
		Arguments: arguments,
	}, nil
}

func (p *Parser) parseUnary() (ast.Expression, error) {
	operator, err := p.mustAdvance()
	if err != nil {
		return nil, err
	}

	operand, err := p.ParseExpression()
	if err != nil {
		return nil, err
	}

	return ast.UnaryExpression{
		Operator: operator,
		Operand:  operand,
	}, nil
}

func (p *Parser) parseAssignment(left ast.Expression) (ast.Expression, error) {
	// consume operator
	operator, err := p.mustAdvance()
	if err != nil {
		return nil, err
	}

	// parse precedence (right associative)
	precedence := precedenceOf(operator.Type)
	right, err := p.parsePrecedence(precedence)
	if err != nil {
		return nil, err
	}

	return ast.AssignmentExpression{
		Left:  left,
		Right: right,
	}, nil
}

func (p *Parser) parseBinary(left ast.Expression) (ast.Expression, error) {
	// consume operator
	operator, err := p.mustAdvance()
	if err != nil {
		return nil, err
	}

	// parse precedence (left associative)
	precedence := precedenceOf(operator.Type).increase()
	right, err := p.parsePrecedence(precedence)
	if err != nil {
		return nil, err
	}

	return ast.BinaryExpression{
		Left:     left,
		Operator: operator,
		Right:    right,
	}, nil
}
