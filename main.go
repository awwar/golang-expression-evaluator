package main

import (
	"expression_parser/parser"
	"expression_parser/tokenizer"
	"expression_parser/virtual_machine"
	"fmt"
)

func main() {
	//expression := "1 + 2 * sum(3 + 4) / 5 + 6"
	expression := "sum(1 + 2, 3.4, sum(5 + 6, 7))"

	fmt.Println(expression)

	tokenizerMachine := tokenizer.New()

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

	if len(tree) != 1 {
		fmt.Println("All nodes must collapse in one node, got: ", len(tree))

		return
	}

	root := tree[0]

	fmt.Println(root.String(0))

	result, err := virtual_machine.Invoke(root)

	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Printf("result = %s", result)
}
