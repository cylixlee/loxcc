package scanner

const (
	// Single character tokens

	TokLeftParenthesis TokenType = iota
	TokRightParenthesis
	TokLeftBrace
	TokRightBrace
	TokComma
	TokDot
	TokMinus
	TokPlus
	TokSemicolon
	TokSlash
	TokStar

	// One or two character tokens

	TokBang
	TokEqual
	TokGreater
	TokLess
	TokBangEqual
	TokEqualEqual
	TokGreaterEqual
	TokLessEqual

	// Literals

	TokIdentifier
	TokString
	TokNumber

	// Keywords

	TokAnd
	TokClass
	TokElse
	TokFalse
	TokFor
	TokFun
	TokIf
	TokNil
	TokOr
	TokPrint
	TokReturn
	TokSuper
	TokThis
	TokTrue
	TokVar
	TokWhile
)

type TokenType byte

type Token struct {
	Type   TokenType
	Lexeme string
	Lineno int
}
