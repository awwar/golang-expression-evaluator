package tokenizer

import (
	"strings"
)

var (
	operations   = map[string]bool{"-": true, "+": true, "/": true, "*": true, ".": true}
	bracers      = map[string]bool{"(": true, ")": true}
	numbers      = "0123456789"
	wordChars    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
	singleTokens = map[int]bool{TypeSemicolon: true, TypeBrackets: true, TypeOperation: true}
)

type Tokenizer struct {
	LastType        int
	Value           string
	Stream          TokenStream
	CurrentPosition int
}

func New() *Tokenizer {
	return &Tokenizer{LastType: TypeEmpty}
}

func (t *Tokenizer) ExpressionToStream(expression *string) (*TokenStream, error) {
	for i := 0; i < len(*expression); i++ {
		char := string((*expression)[i])
		t.CurrentPosition = i

		if char == `"` {
			if t.LastType == TypeString {
				t.changeTokenType(TypeEmpty)
			} else {
				t.changeTokenType(TypeString)
			}

			continue
		} else if t.LastType == TypeString {
		} else if char == " " {
			t.changeTokenType(TypeEmpty)
			continue
		} else if strings.Contains(numbers, char) {
			t.changeTokenType(TypeNumber)
		} else if operations[char] {
			t.changeTokenType(TypeOperation)
		} else if char == "," {
			t.changeTokenType(TypeSemicolon)
		} else if bracers[char] {
			t.changeTokenType(TypeBrackets)
		} else if strings.Contains(wordChars, char) {
			t.changeTokenType(TypeWord)
		} else {
			return nil, &TokenizeError{Position: i, Expression: expression}
		}

		t.Value = t.Value + char
	}

	t.changeTokenType(TypeEOL)

	return &t.Stream, nil
}

func (t *Tokenizer) changeTokenType(newType int) {
	if t.LastType == newType && !singleTokens[newType] {
		return
	}

	if t.LastType != TypeEmpty {
		t.Stream.Push(&Token{Value: t.Value, Type: t.LastType, Position: t.CurrentPosition})
	}

	t.Value = ""
	t.LastType = newType
}

func (t *Tokenizer) swapTokenType(newType int) {
	t.LastType = newType
}
