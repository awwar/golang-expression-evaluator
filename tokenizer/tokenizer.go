package tokenizer

import (
	"strings"
)

var (
	operations = map[string]bool{"-": true, "+": true, "/": true, "*": true}
	bracers    = map[string]bool{"(": true, ")": true}
	numbers    = "0123456789."
	wordChars  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
)

type Tokenizer struct {
	LastType int
	Value    string
	Stream   TokenStream
}

func New() *Tokenizer {
	return &Tokenizer{}
}

func (t *Tokenizer) ExpressionToStream(expression *string) (*TokenStream, error) {
	for i := 0; i < len(*expression); i++ {
		err := t.consume(expression, i)

		if err != nil {
			return nil, err
		}
	}

	return &t.Stream, nil
}

func (t *Tokenizer) consume(expression *string, pos int) error {
	var CurrentType int = -1
	char := string((*expression)[pos])

	if char == " " {
		return nil
	}

	if strings.Contains(numbers, char) {
		CurrentType = TypeNumber

		if t.LastType == TypeWord {
			CurrentType = TypeWord
		}
	} else if operations[char] {
		CurrentType = TypeOperation
	} else if bracers[char] {
		CurrentType = TypeBrackets
	} else if strings.Contains(wordChars, char) {
		CurrentType = TypeWord

		if t.LastType == TypeNumber {
			t.LastType = TypeWord
		}
	} else {
		return &TokenizeError{Position: pos, Expression: expression}
	}

	if t.LastType == -1 {
		t.LastType = CurrentType
	}

	if t.LastType != CurrentType {
		t.Stream.Push(&Token{Value: t.Value, Type: t.LastType})

		t.Value = ""
		t.LastType = CurrentType
	}

	t.Value = t.Value + char

	if len(*expression)-1 == pos {
		t.Stream.Push(&Token{Value: t.Value, Type: CurrentType})
	}

	return nil
}
