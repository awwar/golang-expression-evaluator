package utility

import (
	"fmt"
	"os"
	"strings"
)

type Error struct {
	Message    string
	Position   int
	Expression string
}

func NewError(position int, expression string, message string, params ...any) *Error {
	return &Error{Position: position, Message: fmt.Sprintf(message, params...), Expression: expression}
}

func (e *Error) Error() string {
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

	return fmt.Sprintf("%s\n%s", e.Message, expr)
}

func Must[T any](val T, err error) T {
	if err != nil {
		fmt.Printf("%v\n", err)

		os.Exit(1)
	}

	return val
}
