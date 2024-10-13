package main

import (
	"expression_parser/parser"
	"expression_parser/tokenizer"
	"expression_parser/virtual_machine"
	"fmt"
)

func main() {
	var expression string = "1 + 2 * (3 + 4) / 5 + 6"

	fmt.Println(expression)

	var tokenizerMachine *tokenizer.Tokenizer = &tokenizer.Tokenizer{}

	tokenStream, err := tokenizerMachine.ExpressionToStream(&expression)

	if err != nil {
		fmt.Println(err)

		return
	}
	fmt.Println(tokenStream)

	parseMachine := parser.NewFromStream(tokenStream)

	tree, err := parseMachine.Parse()

	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(tree)

	result, err := virtual_machine.Invoke(tree)

	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Printf("result = %f", result)
}
