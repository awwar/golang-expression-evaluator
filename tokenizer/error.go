package tokenizer

import (
	"fmt"
	"strings"
)

type TokenizeError struct {
	Position   int
	Expression string
}

func (e *TokenizeError) Error() string {
	pre := e.Expression[:e.Position]
	post := e.Expression[e.Position:]

	offset := strings.LastIndex(pre, "\n")
	if offset == -1 {
		offset = e.Position
	} else {
		offset = e.Position - offset - 1
	}

	nextNewLine := strings.Index(post, "\n")
	if nextNewLine == -1 {
		nextNewLine = len(post)
	}

	expr := pre + post[:nextNewLine] + "\n" + strings.Repeat(" ", offset) + "^" + post[nextNewLine:]

	return fmt.Sprintf("Invalid token at position %d \n%s", e.Position, expr)
}
