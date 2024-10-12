package tokenizer

import (
	"regexp"
)

var (
	operations = map[string]bool{"-": true, "+": true, "/": true, "*": true}
	bracers    = map[string]bool{"(": true, ")": true}
)

type Tokenizer struct {
	LastType int
	Value    string
	Stream   TokenStream
}

func (ths *Tokenizer) ExpressionToStream(expression *string) (*TokenStream, error) {
	for i := 0; i < len(*expression); i++ {
		err := ths.consume(expression, i)

		if err != nil {
			return nil, err
		}
	}

	return &ths.Stream, nil
}

func (ths *Tokenizer) consume(expression *string, pos int) error {
	var CurrentType int = -1
	char := string((*expression)[pos])

	if char == " " {
		return nil
	}

	isNumber, _ := regexp.MatchString(`[0-9.]`, char)

	if isNumber {
		CurrentType = TypeNumber
	} else if operations[char] {
		CurrentType = TypeOperation
	} else if bracers[char] {
		CurrentType = TypeBrackets
	} else {
		return &TokenizeError{Position: pos, Expression: expression}
	}

	if ths.LastType == -1 {
		ths.LastType = CurrentType
	}

	if ths.LastType != CurrentType {
		ths.Stream.Push(&Token{Value: ths.Value, Type: ths.LastType})

		ths.Value = ""
		ths.LastType = CurrentType
	}

	ths.Value = ths.Value + char

	if len(*expression)-1 == pos {
		ths.Stream.Push(&Token{Value: ths.Value, Type: CurrentType})
	}

	return nil
}
