package main

import (
	"expression_parser/parser"
	"expression_parser/tokenizer"
	"fmt"
)

func main() {
	var expression string = "3 * (222 + 1)"

	var tokenizerMachine *tokenizer.Tokenizer = &tokenizer.Tokenizer{}

	tokenStream, err := tokenizerMachine.ExpressionToStream(&expression)

	if err != nil {
		fmt.Println(err)

		return
	}
	fmt.Println(tokenStream)

	pars := parser.NewFromStream(tokenStream)

	tree, err := pars.Parse()

	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(tree)
}
