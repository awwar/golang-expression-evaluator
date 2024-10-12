package tokenizer

import (
	"fmt"
	"strings"
)

type TokenizeError struct {
	Position   int
	Expression *string
}

func (e *TokenizeError) Error() string {
	return fmt.Sprintf("Invalid token\n%s\n%s^", *e.Expression, strings.Repeat(" ", e.Position))
}
