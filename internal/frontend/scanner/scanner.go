package scanner

var (
	// Mapping from single character to a concrete TokenType.
	//
	// This is used to reduce switch-case boilerplate, although it may get slower than the
	// original version. Trade speed for maintainability is worthwhile.
	sTokensMapping = map[rune]TokenType{
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
	dTokensMapping = map[rune]dtype{
		'!': {'=', TokBangEqual},
		'=': {'=', TokEqualEqual},
		'>': {'=', TokGreaterEqual},
		'<': {'=', TokLessEqual},
	}
)

// Type for process some token with higher priority.
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

func (s *Scanner) Scan() (*Token, error) {
	if err := s.skipWhitespace(); err != nil {
		return nil, err
	}
	s.start = s.current

	c, err := s.advance()
	if err != nil {
		return nil, err
	}

	if d, exists := dTokensMapping[c]; exists {
		if s.match(d.Expect) {
			return s.makeToken(d.Then)
		}
	}

	if t, exists := sTokensMapping[c]; exists {
		return s.makeToken(t)
	}
	return nil, UnexpectedCharacterError(s.source[s.current])
}

// Peeks the current character (rune), without moving s.current forward.
//
// This function may return [ErrEndOfSource] as error when there's no more character. No
// other errors will be returned.
func (s Scanner) peek() (rune, error) {
	if s.current >= len(s.source) {
		return 0, ErrEndOfSource
	}
	return s.source[s.current], nil
}

// Peeks the next character without moving s.current forward.
func (s Scanner) peekNext() (rune, error) {
	if s.current+1 >= len(s.source) {
		return 0, ErrEndOfSource
	}
	return s.source[s.current+1], nil
}

// Advances the s.current pointer, returns the consumed character, or an error if any.
//
// This function can and can only return [ErrEndOfSource] as error.
func (s Scanner) advance() (rune, error) {
	peek, err := s.peek()
	if err != nil {
		return 0, err
	}
	s.current++
	return peek, nil
}

// Returns true when the next character is as expected.
//
// This function will never fail. It only depends on peekNext, which can and can only
// return [ErrEndOfSource] when there's no more character. In that case, this function
// returns false, because the "next character" is not as expected.
func (s *Scanner) match(expect rune) bool {
	peek, err := s.peekNext()
	if err != nil {
		return false
	}

	if peek != expect {
		return false
	}
	s.current++
	return true
}

func (s *Scanner) skipWhitespace() error {
	for {
		c, err := s.peek()
		if err != nil {
			if err == ErrEndOfSource {
				return nil
			}
			return err
		}

		switch c {
		case '\n':
			s.lineno++
			fallthrough
		case ' ', '\t', '\r':
			s.advance()
		case '/':
			if s.current+1 < len(s.source) {
				next, err := s.peekNext()
				if err != nil {
					if err == ErrEndOfSource {
						return nil
					}
					return err
				}

				if next == '/' {
					for {
						peek, err := s.peek()
						if err != nil {
							if err == ErrEndOfSource {
								break
							}
							return err
						}

						if peek == '\n' {
							break
						}
						s.advance()
					}
				} else {
					return nil
				}
			}
		default:
			return nil
		}
	}
}

// Returns a token with given type and source[start:current] as lexeme. An extra nil value
// (as [error] type) is returned in order to reduce boilerplate.
func (s Scanner) makeToken(t TokenType) (*Token, error) {
	return &Token{
		Type:   t,
		Lexeme: string(s.source[s.start:s.current]),
		Lineno: s.lineno,
	}, nil
}
