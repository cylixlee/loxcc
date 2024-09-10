package scanner

import (
	"errors"
	"fmt"
)

var (
	// Mapping from single character to a concrete [TokenType].
	//
	// This is used to reduce switch-case boilerplate, although it may get slower than the
	// original version. Trade speed for maintainability is worthwhile.
	sTokenMap = map[rune]TokenType{
		'(': TokLeftParenthesis,
		')': TokRightParenthesis,
		'{': TokLeftBrace,
		'}': TokRightBrace,
		';': TokSemicolon,
		',': TokComma,
		'.': TokDot,
		'-': TokMinus,
		'+': TokPlus,
		'/': TokSlash,
		'*': TokStar,
		'!': TokBang,
		'=': TokEqual,
		'>': TokGreater,
		'<': TokLess,
	}

	// Mapping for processing one or two character tokens.
	// See [dtype] for details.
	dTokenMap = map[rune]dtype{
		'!': {'=', TokBangEqual},
		'=': {'=', TokEqualEqual},
		'>': {'=', TokGreaterEqual},
		'<': {'=', TokLessEqual},
	}

	// Mapping from keywords to a concrete [TokenType].
	//
	// Keywords are special identifiers, whose [TokenType] is determined after making the
	// token itself. We can just lookup the map for a specific [TokenType].
	keywordMap = map[string]TokenType{
		"and":    TokAnd,
		"class":  TokClass,
		"else":   TokElse,
		"false":  TokFalse,
		"for":    TokFor,
		"fun":    TokFun,
		"if":     TokIf,
		"nil":    TokNil,
		"or":     TokOr,
		"print":  TokPrint,
		"return": TokReturn,
		"super":  TokSuper,
		"this":   TokThis,
		"true":   TokTrue,
		"var":    TokVar,
		"while":  TokWhile,
	}

	ErrUnterminatedString = errors.New("unterminated string")
)

// Type for process some token with higher priority. [dtype] is the abbreviation of
// double-character token helper type.
//
// For example, the greater sign [>] and a following equal sign [=] can form a
// [TokGreaterEqual], instead of two tokens [TokGreater] and [TokEqual].
type dtype struct {
	Expect rune
	Then   TokenType
}

type Scanner struct {
	source  []rune
	start   int
	current int
	lineno  int
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source:  []rune(source),
		start:   0,
		current: 0,
		lineno:  1,
	}
}

// Scans one single [Token] and returns.
//
// There's three patterns of return values:
//  1. *Token and nil: Successfully scanned a token without any errors.
//  2. nil and nil: Scanner has reached EOF, there's no more [Token].
//  3. nil and error: Error occurred while scanning.
//
// If there's an error, it'll be of those kinds:
//  1. Unexpected character. When encountered with an unexpected character, returns a
//     distinct error value, even if the character is identical.
//  2. [ErrUnterminatedString]. When a string isn't terminated with the double quotation
//     mark '"'.
func (s *Scanner) Scan() (*Token, error) {
	// skips whitespace
	s.skipWhitespace()
	s.start = s.current

	// advances (consumes) a character
	c, ok := s.advance()
	if !ok {
		return nil, nil
	}

	// check whether it's a double character token
	// more specifically, an operator (e.g. [TokGreaterEqual] >=)
	if d, exists := dTokenMap[c]; exists {
		if s.match(d.Expect) {
			return s.makeToken(d.Then), nil
		}
	}

	// check whether it's a single character token
	if t, exists := sTokenMap[c]; exists {
		return s.makeToken(t), nil
	}

	// try scanning the token as literals
	switch {
	case c == '"':
		return s.scanString()
	case isDigit(c):
		return s.scanNumber()
	case isAlpha(c):
		return s.scanIdentifier()
	}

	// finally we can say the character is not recognizable
	return nil, fmt.Errorf("unexpected character %c", c)
}

// Scans a double-quoted string literal.
//
// This function is called when the leading double quotation mark is consumed, and it
// consumes the tailing one. If the string is not terminated with '"',
// [ErrUnterminatedString] is returned.
func (s *Scanner) scanString() (*Token, error) {
	for {
		peek, ok := s.peek()
		if !ok || peek == '"' {
			break
		}

		if peek == '\n' {
			s.lineno++
		}
		s.advance()
	}

	terminator, ok := s.peek()
	if !ok || terminator != '"' {
		return nil, ErrUnterminatedString
	}
	s.advance()
	return s.makeToken(TokString), nil
}

// Scans the current token as a number.
//
// Numbers consist of an integer part and an optional decimal part.
func (s *Scanner) scanNumber() (*Token, error) {
	for {
		peek, ok := s.peek()
		if !ok || !isDigit(peek) {
			break
		}
		s.advance()
	}

	peek, ok1 := s.peek()
	next, ok2 := s.peekNext()
	if ok1 && ok2 && peek == '.' && isDigit(next) {
		s.advance()
		for {
			peek, ok1 = s.peek()
			if !ok1 || !isDigit(peek) {
				break
			}
			s.advance()
		}
	}
	return s.makeToken(TokNumber), nil
}

// Scans the current token as an identifier.
//
// Note that a keyword is a valid identifier, so we need to re-determine its [TokenType]
// after the token is made.
func (s *Scanner) scanIdentifier() (*Token, error) {
	for {
		peek, ok := s.peek()
		if !ok || !isAlplanumeric(peek) {
			break
		}
		s.advance()
	}

	token := s.makeToken(TokIdentifier)
	if t, exists := keywordMap[token.Lexeme]; exists {
		token.Type = t
	}
	return token, nil
}

// Peeks the current character (rune), without moving s.current forward.
//
// If there's no more character available, this function returns 0 as rune and a false
// value indicating the operation is not successful.
func (s Scanner) peek() (rune, bool) {
	if s.current >= len(s.source) {
		return 0, false
	}
	return s.source[s.current], true
}

// Peeks the next character without moving s.current forward.
//
// If there's no more character at next position, this function returns 0 as rune and a
// false value indicating the operation is not successful.
func (s Scanner) peekNext() (rune, bool) {
	if s.current+1 >= len(s.source) {
		return 0, false
	}
	return s.source[s.current+1], true
}

// Advances the s.current pointer, returns the consumed character, or false if there's no
// character ahead.
func (s *Scanner) advance() (rune, bool) {
	peek, ok := s.peek()
	if !ok {
		return 0, false
	}
	s.current++
	return peek, true
}

// Returns true when the next character is as expected.
//
// This function never fails. It only depends on peekNext, which may return a false value
// indicating there's no character at next position. In that case, this function returns
// false, because the "next character" is, indeed, not as expected.
func (s *Scanner) match(expect rune) bool {
	peek, ok := s.peekNext()
	if !ok {
		return false
	}

	if peek != expect {
		return false
	}
	s.current++
	return true
}

// Skips all whitespaces (including comments) before Lox source.
//
// This function never fails. It depends on peek, peekNext and advance, which can and can
// only return [ErrEndOfSource]. It is legal to skip all whitespaces until EOF.
func (s *Scanner) skipWhitespace() {
	for {
		c, ok := s.peek()
		if !ok {
			return
		}

		switch c {
		// newline
		case '\n':
			s.lineno++
			fallthrough

		// regular whitespaces
		case ' ', '\t', '\r':
			s.advance()

		// comment (only line comment is supported)
		case '/':
			if s.current+1 < len(s.source) {
				next, ok := s.peekNext()
				if !ok {
					return
				}

				if next == '/' {
					for {
						peek, ok := s.peek()
						if !ok {
							return
						}

						if peek == '\n' {
							break
						}
						s.advance()
					}
				} else {
					return
				}
			}
		default:
			return
		}
	}
}

// Returns a token with given type and source[start:current] as lexeme.
func (s Scanner) makeToken(t TokenType) *Token {
	return &Token{
		Type:   t,
		Lexeme: string(s.source[s.start:s.current]),
		Lineno: s.lineno,
	}
}
