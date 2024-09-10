package parser

import (
	"loxcc/internal/ast"
	"loxcc/internal/frontend/scanner"

	stl "github.com/chen3feng/stl4go"
)

func (p *Parser) ParseStatement() (ast.Statement, error) {
	panic("unimplemented")
}

func (p *Parser) parseExpressionStatement() (ast.Statement, error) {
	expr, err := p.ParseExpression()
	if err != nil {
		return nil, err
	}

	if _, err := p.mustConsume(scanner.TokSemicolon); err != nil {
		return nil, err
	}
	return ast.ExpressionStatement{Expr: expr}, nil
}

func (p *Parser) parseForStatement() (ast.Statement, error) {
	var err error

	if _, err = p.mustConsume(scanner.TokFor); err != nil {
		return nil, err
	}
	if _, err = p.mustConsume(scanner.TokLeftParenthesis); err != nil {
		return nil, err
	}

	var initializer *ast.ForLoopInitializer
	if !p.tryConsume(scanner.TokSemicolon) {
		peek, err := p.mustPeek()
		if err != nil {
			return nil, err
		}

		if peek.Type == scanner.TokVar {
			// initializer = ast.ForLoopInitializer{
			// 	Kind:           ast.VarDecl,
			// 	VarInitializer: ast.VarDeclaration{},
			// }
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
		//nolint:staticcheck
		//
		// The [condition] variable, whether nil or not, is finally passed to the
		// [ast.ForStatement] object. It's a false positive.
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
		//nolint:staticcheck
		//
		// The [incrementer], whether nil or not, is finally passed to [ast.ForStatement]
		// object. It's false positive due to staticcheck.
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

func (p *Parser) parseIfStatement() (ast.Statement, error) {
	if _, err := p.mustConsume(scanner.TokIf); err != nil {
		return nil, err
	}
	if _, err := p.mustConsume(scanner.TokLeftParenthesis); err != nil {
		return nil, err
	}

	//nolint:staticcheck
	condition, err := p.ParseExpression()
	if err != nil {
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

func (p *Parser) parsePrintStatement() (ast.Statement, error) {
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

func (p *Parser) parseReturnStatement() (ast.Statement, error) {
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

func (p *Parser) parseWhileStatement() (ast.Statement, error) {
	if _, err := p.mustConsume(scanner.TokWhile); err != nil {
		return nil, err
	}
	if _, err := p.mustConsume(scanner.TokLeftParenthesis); err != nil {
		return nil, err
	}

	//nolint:staticcheck
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

func (p *Parser) parseBlockStatement() (ast.Statement, error) {
	if _, err := p.mustConsume(scanner.TokLeftBrace); err != nil {
		return nil, err
	}

	var declarations stl.Vector[ast.Declaration]
	for !p.tryConsume(scanner.TokRightBrace) {
		decl, err := p.ParseDeclaration()
		if err != nil {
			return nil, err
		}
		declarations.PushBack(decl)
	}
	return ast.BlockStatement{Content: declarations}, nil
}
