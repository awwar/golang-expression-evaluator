package tokenizer

import (
	"fmt"
	"strings"
)

type TokenStream struct {
	Tokens []*Token
}

func (t *TokenStream) Length() int {
	return len(t.Tokens)
}

func (t *TokenStream) Push(token *Token) {
	t.Tokens = append(t.Tokens, token)
}

func (t *TokenStream) Get(index int) *Token {
	var token *Token

	if t.Length() > index {
		token = t.Tokens[index]
	}

	return token
}

func (t *TokenStream) SearchIdxOfClosedBracer(startBracer int) int {
	var bracersCount int = 0
	var currenPosition int = startBracer

	for {
		value := t.Get(currenPosition)

		if value == nil {
			break
		}

		token := *value

		if token.Value == "(" {
			bracersCount++
		} else if token.Value == ")" {
			bracersCount--
		}

		if bracersCount == 0 {
			return currenPosition
		}

		currenPosition++
	}

	return -1
}

func (t *TokenStream) NextTokenIsBracer(position int) bool {
	nextToken := t.Get(position + 1)

	if nextToken == nil {
		return false
	}

	return nextToken.Type == TypeBrackets
}

func (t *TokenStream) String() string {
	var output string = ""

	for i := range t.Tokens {
		var value Token = *t.Tokens[i]

		output = fmt.Sprintf("%s\n	%s", output, value.String())
	}

	output = strings.TrimLeft(output, ",  \n")

	return fmt.Sprintf("[\n%s\n]", output)
}
