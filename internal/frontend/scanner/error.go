package scanner

import (
	"errors"
	"fmt"
)

var (
	ErrEndOfSource error = errors.New("end of source")
)

type UnexpectedCharacterError rune

func (u UnexpectedCharacterError) Error() string {
	return fmt.Sprintf("unexpected character %c", u)
}
