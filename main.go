package main

import (
	"fmt"

	"expression_parser/compiler"
	"expression_parser/parser"
	"expression_parser/tokenizer"
	"expression_parser/virtual_machine"
)

func main() {
	expression := `"result = ".uppercase() + (-1 + -2sum(3.4, 4) / 5 + 6)`

	fmt.Println(expression)

	tokenizerMachine := tokenizer.New()

	tokenStream, err := tokenizerMachine.ExpressionToStream(&expression)
	if err != nil {
		fmt.Println(err)

		return
	}

	parseMachine := parser.NewFromStream(tokenStream)

	tree, parseError := parseMachine.Parse()
	if parseError != nil {
		parseError.EnrichWithExpression(&expression)

		fmt.Println(parseError)

		return
	}

	if len(tree) != 1 {
		fmt.Println("All nodes must collapse in one node, got: ", len(tree))

		for _, rt := range tree {
			fmt.Println(rt.String(0))
		}

		return
	}

	root := tree[0]

	fmt.Println(root.String(0))

	compile := compiler.NewCompiler()

	program, err := compile.Compile(root)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(program.String())

	result, err := virtual_machine.Execute(*program)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Printf("%s", result)
}
