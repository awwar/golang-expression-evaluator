package tokenizer

import "fmt"

const (
	TypeNumber    = iota
	TypeOperation = iota
	TypeBrackets  = iota
	TypeWord      = iota
	TypeSemicolon = iota
	TypeString    = iota
)

var MapTypeToTypeName = map[int]string{TypeNumber: "number", TypeOperation: "operation", TypeBrackets: "bracket", TypeWord: "word", TypeSemicolon: "semicolon", TypeString: "string"}

type Token struct {
	Value string
	Type  int
}

func (t *Token) String() string {
	return fmt.Sprintf("{value: \"%s\", type: %s}", t.Value, MapTypeToTypeName[t.Type])
}
