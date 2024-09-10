package parser

import (
	"errors"
	"loxcc/internal/frontend/scanner"
)

var (
	ErrEarlyEOF        = errors.New("expect some token, got EOF")
	ErrUnexpectedToken = errors.New("unexpected token")
)

func (p Parser) peek() *scanner.Token {
	if p.current < p.tokens.Len() {
		return p.tokens[p.current]
	}
	return nil
}

func (p *Parser) advance() *scanner.Token {
	if p.current < p.tokens.Len() {
		p.current++
		return p.tokens[p.current-1]
	}
	return nil
}

func (p *Parser) tryConsume(t scanner.TokenType) bool {
	if peek := p.peek(); peek != nil {
		if peek.Type == t {
			p.advance()
			return true
		}
	}
	return false
}

func (p Parser) mustPeek() (*scanner.Token, error) {
	if peek := p.peek(); peek != nil {
		return peek, nil
	}
	return nil, ErrEarlyEOF
}

func (p *Parser) mustAdvance() (*scanner.Token, error) {
	token := p.advance()
	if token == nil {
		return nil, ErrEarlyEOF
	}
	return token, nil
}

func (p *Parser) mustConsume(t scanner.TokenType) (*scanner.Token, error) {
	token, err := p.mustAdvance()
	if err != nil {
		return nil, err
	}

	if token.Type == t {
		return token, nil
	}
	return nil, ErrUnexpectedToken
}
