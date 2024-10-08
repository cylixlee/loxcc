package parser

import (
	"loxcc/internal/ast"
	"loxcc/internal/frontend/scanner"

	stl "github.com/chen3feng/stl4go"
)

func (p *parser) ParseDeclaration() (ast.Declaration, error) {
	peek, err := p.mustPeek()
	if err != nil {
		return nil, err
	}

	switch peek.Type {
	case scanner.TokClass:
		return p.parseClassDeclaration()
	case scanner.TokFun:
		return p.parseFunDeclaration()
	case scanner.TokVar:
		return p.parseVarDeclaration()
	}

	stmt, err := p.ParseStatement()
	if err != nil {
		return nil, err
	}
	return ast.StatementDeclaration{Stmt: stmt}, nil
}

func (p *parser) parseClassDeclaration() (ast.Declaration, error) {
	if _, err := p.mustConsume(scanner.TokClass); err != nil {
		return nil, err
	}

	name, err := p.mustConsume(scanner.TokIdentifier)
	if err != nil {
		return nil, err
	}

	var baseclass *scanner.Token
	if p.tryConsume(scanner.TokLess) {
		baseclass, err = p.mustAdvance()
		if err != nil {
			return nil, err
		}
	}

	methods := stl.MakeVector[ast.Declaration]()
	if _, err := p.mustConsume(scanner.TokLeftBrace); err != nil {
		return nil, err
	}

	for !p.tryConsume(scanner.TokRightBrace) {
		method, err := p.parseMethod()
		if err != nil {
			return nil, err
		}
		methods.PushBack(method)
	}

	return ast.ClassDeclaration{
		Name:      name,
		Baseclass: baseclass,
		Methods:   methods,
	}, nil
}

func (p *parser) parseFunDeclaration() (ast.Declaration, error) {
	if _, err := p.mustConsume(scanner.TokFun); err != nil {
		return nil, err
	}
	return p.parseMethod()
}

func (p *parser) parseVarDeclaration() (ast.Declaration, error) {
	if _, err := p.mustConsume(scanner.TokVar); err != nil {
		return nil, err
	}

	name, err := p.mustConsume(scanner.TokIdentifier)
	if err != nil {
		return nil, err
	}

	var initializer ast.Expression
	if p.tryConsume(scanner.TokEqual) {
		initializer, err = p.ParseExpression()
		if err != nil {
			return nil, err
		}
	}

	if _, err := p.mustConsume(scanner.TokSemicolon); err != nil {
		return nil, err
	}

	return ast.VarDeclaration{
		Name:        name,
		Initializer: initializer,
	}, nil
}

func (p *parser) parseMethod() (ast.Declaration, error) {
	name, err := p.mustConsume(scanner.TokIdentifier)
	if err != nil {
		return nil, err
	}

	parameters, err := p.parseParameters()
	if err != nil {
		return nil, err
	}

	body, err := p.parseBlockStatement()
	if err != nil {
		return nil, err
	}

	return ast.FunctionDeclaration{
		Name:       name,
		Parameters: parameters,
		Body:       body,
	}, nil
}

func (p *parser) parseParameters() (stl.Vector[*scanner.Token], error) {
	if _, err := p.mustConsume(scanner.TokLeftParenthesis); err != nil {
		return nil, err
	}

	parameters := stl.MakeVector[*scanner.Token]()
	if peek := p.peek(); peek != nil && peek.Type == scanner.TokIdentifier {
		ident, err := p.mustConsume(scanner.TokIdentifier)
		if err != nil {
			return nil, err
		}
		parameters.PushBack(ident)

		for p.tryConsume(scanner.TokComma) {
			ident, err := p.mustConsume(scanner.TokIdentifier)
			if err != nil {
				return nil, err
			}
			parameters.PushBack(ident)
		}
	}

	if _, err := p.mustConsume(scanner.TokRightParenthesis); err != nil {
		return nil, err
	}
	return parameters, nil
}

func (p *parser) parseArguments() (stl.Vector[ast.Expression], error) {
	// consume left parenthesis
	if _, err := p.mustConsume(scanner.TokLeftParenthesis); err != nil {
		return nil, err
	}

	// create vector (slice)
	arguments := stl.MakeVector[ast.Expression]()

	if !p.tryConsume(scanner.TokRightParenthesis) {
		// parse the first argument
		expr, err := p.ParseExpression()
		if err != nil {
			return nil, err
		}
		arguments.PushBack(expr)

		// parse more arguments, each of which is lead by a [TokComma] (,).
		for p.tryConsume(scanner.TokComma) {
			expr, err = p.ParseExpression()
			if err != nil {
				return nil, err
			}
			arguments.PushBack(expr)
		}

		// parse right parenthesis
		if _, err := p.mustConsume(scanner.TokRightParenthesis); err != nil {
			return nil, err
		}
	}

	return arguments, nil
}
