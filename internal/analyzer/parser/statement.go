package parser

import (
	"loxcc/internal/analyzer/scanner"
	"loxcc/internal/ast"

	stl "github.com/chen3feng/stl4go"
)

func (p *parser) ParseStatement() (ast.Statement, error) {
	peek, err := p.mustPeek()
	if err != nil {
		return nil, err
	}

	switch peek.Type {
	case scanner.TokFor:
		return p.parseForStatement()
	case scanner.TokIf:
		return p.parseIfStatement()
	case scanner.TokPrint:
		return p.parsePrintStatement()
	case scanner.TokReturn:
		return p.parseReturnStatement()
	case scanner.TokWhile:
		return p.parseWhileStatement()
	case scanner.TokLeftBrace:
		return p.parseBlockStatement()
	}
	return p.parseExpressionStatement()
}

func (p *parser) parseExpressionStatement() (ast.Statement, error) {
	expr, err := p.ParseExpression()
	if err != nil {
		return nil, err
	}

	if _, err := p.mustConsume(scanner.TokSemicolon); err != nil {
		return nil, err
	}
	return ast.ExpressionStatement{Expr: expr}, nil
}

func (p *parser) parseForStatement() (ast.Statement, error) {
	var err error

	if _, err := p.mustConsume(scanner.TokFor); err != nil {
		return nil, err
	}
	if _, err := p.mustConsume(scanner.TokLeftParenthesis); err != nil {
		return nil, err
	}

	var initializer *ast.ForLoopInitializer
	if !p.tryConsume(scanner.TokSemicolon) {
		peek, err := p.mustPeek()
		if err != nil {
			return nil, err
		}

		if peek.Type == scanner.TokVar {
			decl, err := p.parseVarDeclaration()
			if err != nil {
				return nil, err
			}

			initializer = &ast.ForLoopInitializer{
				Kind:           ast.VarDecl,
				VarInitializer: decl,
			}
		} else {
			expr, err := p.ParseExpression()
			if err != nil {
				return nil, err
			}

			initializer = &ast.ForLoopInitializer{
				Kind:            ast.InitExpr,
				ExprInitializer: expr,
			}
		}
	}

	var condition ast.Expression
	if !p.tryConsume(scanner.TokSemicolon) {
		condition, err = p.ParseExpression()
		if err != nil {
			return nil, err
		}

		_, err = p.mustConsume(scanner.TokSemicolon)
		if err != nil {
			return nil, err
		}
	}

	var incrementer ast.Expression
	if !p.tryConsume(scanner.TokRightParenthesis) {
		incrementer, err = p.ParseExpression()
		if err != nil {
			return nil, err
		}

		if _, err := p.mustConsume(scanner.TokRightParenthesis); err != nil {
			return nil, err
		}
	}

	body, err := p.ParseStatement()
	if err != nil {
		return nil, err
	}
	return ast.ForStatement{
		Initializer: initializer,
		Condition:   condition,
		Incrementer: incrementer,
		Body:        body,
	}, nil
}

func (p *parser) parseIfStatement() (ast.Statement, error) {
	if _, err := p.mustConsume(scanner.TokIf); err != nil {
		return nil, err
	}
	if _, err := p.mustConsume(scanner.TokLeftParenthesis); err != nil {
		return nil, err
	}

	condition, err := p.ParseExpression()
	if err != nil {
		return nil, err
	}

	if _, err := p.mustConsume(scanner.TokRightParenthesis); err != nil {
		return nil, err
	}

	then, err := p.ParseStatement()
	if err != nil {
		return nil, err
	}

	var elseBranch ast.Statement
	if p.tryConsume(scanner.TokElse) {
		elseBranch, err = p.ParseStatement()
		if err != nil {
			return nil, err
		}
	}
	return ast.IfStatement{
		Condition: condition,
		Then:      then,
		Else:      elseBranch,
	}, nil
}

func (p *parser) parsePrintStatement() (ast.Statement, error) {
	if _, err := p.mustConsume(scanner.TokPrint); err != nil {
		return nil, err
	}

	expr, err := p.ParseExpression()
	if err != nil {
		return nil, err
	}

	if _, err := p.mustConsume(scanner.TokSemicolon); err != nil {
		return nil, err
	}
	return ast.PrintStatement{Value: expr}, nil
}

func (p *parser) parseReturnStatement() (ast.Statement, error) {
	var err error

	if _, err := p.mustConsume(scanner.TokReturn); err != nil {
		return nil, err
	}

	var returnValue ast.Expression
	if !p.tryConsume(scanner.TokSemicolon) {
		returnValue, err = p.ParseExpression()
		if err != nil {
			return nil, err
		}
	}

	if _, err = p.mustConsume(scanner.TokSemicolon); err != nil {
		return nil, err
	}
	return ast.ReturnStatement{Value: returnValue}, nil
}

func (p *parser) parseWhileStatement() (ast.Statement, error) {
	if _, err := p.mustConsume(scanner.TokWhile); err != nil {
		return nil, err
	}
	if _, err := p.mustConsume(scanner.TokLeftParenthesis); err != nil {
		return nil, err
	}

	condition, err := p.ParseExpression()
	if err != nil {
		return nil, err
	}

	if _, err := p.mustConsume(scanner.TokRightParenthesis); err != nil {
		return nil, err
	}

	body, err := p.ParseStatement()
	if err != nil {
		return nil, err
	}
	return ast.WhileStatement{
		Condition: condition,
		Body:      body,
	}, nil
}

func (p *parser) parseBlockStatement() (ast.Statement, error) {
	if _, err := p.mustConsume(scanner.TokLeftBrace); err != nil {
		return nil, err
	}

	declarations := stl.MakeVector[ast.Declaration]()
	for !p.tryConsume(scanner.TokRightBrace) {
		decl, err := p.ParseDeclaration()
		if err != nil {
			return nil, err
		}
		declarations.PushBack(decl)
	}
	return ast.BlockStatement{Content: declarations}, nil
}
