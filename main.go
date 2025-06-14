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

	tokenStream := utility.Must(tokenizer.New().ExpressionToStream(&input))

	tree := utility.Must(parser.NewFromStream(tokenStream).ParseProgram())

	program := utility.Must(compiler.NewCompiler().Compile(tree))

	result := utility.Must(virtual_machine.Execute(*program))

	//fmt.Println(input)
	fmt.Println(tokenStream)
	fmt.Println(tree.String(0))
	fmt.Println(program.String())
	fmt.Printf("%s", result)
}
