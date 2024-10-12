package tokenizer

import "fmt"

const (
	TypeNumber    = iota
	TypeOperation = iota
	TypeBrackets  = iota
)

var MapTypeToTypeName = map[int]string{TypeNumber: "number", TypeOperation: "operation", TypeBrackets: "bracket"}

type Token struct {
	Value string
	Type  int
}

func (ths *Token) String() string {
	return fmt.Sprintf("{value: \"%s\", type: %s}", ths.Value, MapTypeToTypeName[ths.Type])
}
