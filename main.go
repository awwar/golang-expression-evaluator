package main

import (
	"fmt"
	"os"

	"expression_parser/compiler"
	"expression_parser/parser"
	"expression_parser/tokenizer"
	"expression_parser/utility"
	"expression_parser/virtual_machine"
)

func main() {
	input := string(utility.Must(os.ReadFile(".example/index.mp")))

	//fmt.Println(input)

	tokenStream := utility.Must(tokenizer.New().ExpressionToStream(&input))

	fmt.Println(tokenStream)

	tree := utility.Must(parser.NewFromStream(tokenStream).ParseProgram())

	fmt.Println(tree.String(0))

	program := utility.Must(compiler.NewCompiler().Compile(tree))

	fmt.Println(program.String())

	result := utility.Must(virtual_machine.Execute(*program))

	fmt.Printf("%s", result)
}
