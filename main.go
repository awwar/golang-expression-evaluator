package main

import (
	"fmt"
	"os"

	"expression_parser/tokenizer"
	"expression_parser/utility"
)

func main() {
	input := string(utility.Must(os.ReadFile(".example/index.mp")))

	//fmt.Println(input)

	tokenizerMachine := tokenizer.New()

	tokenStream := utility.Must(tokenizerMachine.ExpressionToStream(&input))

	fmt.Println(tokenStream)

	//parseMachine := expression.NewFromStream(tokenStream)
	//
	//tree, parseError := parseMachine.Parse()
	//if parseError != nil {
	//	parseError.EnrichWithExpression(&input)
	//
	//	fmt.Println(parseError)
	//
	//	return
	//}
	//
	//if len(tree) != 1 {
	//	fmt.Println("All nodes must collapse in one node, got: ", len(tree))
	//
	//	for _, rt := range tree {
	//		fmt.Println(rt.String(0))
	//	}
	//
	//	return
	//}
	//
	//root := tree[0]
	//
	//fmt.Println(root.String(0))
	//
	//compile := compiler.NewCompiler()
	//
	//program, err := compile.Compile(root)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(program.String())
	//
	//result, err := virtual_machine.Execute(*program)
	//if err != nil {
	//	fmt.Println(err)
	//
	//	return
	//}
	//
	//fmt.Printf("%s", result)
}
