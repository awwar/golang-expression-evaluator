package tokenizer

import (
	"fmt"
	"strings"
)

type TokenStream struct {
	Tokens []*Token
}

func (ths *TokenStream) Length() int {
	return len(ths.Tokens)
}

func (ths *TokenStream) Push(token *Token) {
	ths.Tokens = append(ths.Tokens, token)
}

func (ths *TokenStream) Get(index int) *Token {
	var token *Token

	if ths.Length() > index {
		token = ths.Tokens[index]
	}

	return token
}

func (ths *TokenStream) SearchIdxOfClosedBracer(startBracer int) int {
	var bracersCount int = 0
	var currenPosition int = startBracer

	for {
		value := ths.Get(currenPosition)

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

func (ths *TokenStream) New() *TokenStream { return &TokenStream{} }

func (ths *TokenStream) String() string {
	var output string = ""

	for i := range ths.Tokens {
		var value Token = *ths.Tokens[i]

		output = fmt.Sprintf("%s\n	%s", output, value.String())
	}

	output = strings.TrimLeft(output, ",  \n")

	return fmt.Sprintf("[\n%s\n]", output)
}
