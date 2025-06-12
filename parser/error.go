package parser

import (
	"fmt"
	"strings"
)

type Error struct {
	Message    string
	Position   int
	Expression *string
}

func NewError(position int, message string, params ...any) *Error {
	return &Error{Position: position, Message: fmt.Sprintf(message, params...), Expression: nil}
}

func (e *Error) EnrichWithExpression(expression *string) {
	e.Expression = expression
}

func (e *Error) Error() string {
	expressionHint := ""

	if e.Expression != nil {
		expressionHint = fmt.Sprintf("%s\n%s^", *e.Expression, strings.Repeat(" ", e.Position))
	}

	return fmt.Sprintf("%s\n%s", e.Message, expressionHint)
}
