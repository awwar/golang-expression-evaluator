package tokenizer

import "fmt"

const (
	TypeEmpty     = iota
	TypeNumber    = iota
	TypeOperation = iota
	TypeBrackets  = iota
	TypeWord      = iota
	TypeSemicolon = iota
	TypeString    = iota
	TypeEOL       = iota
)

var MapTypeToTypeName = map[int]string{
	TypeEmpty:     "empty",
	TypeNumber:    "number",
	TypeOperation: "operation",
	TypeBrackets:  "bracket",
	TypeWord:      "word",
	TypeSemicolon: "semicolon",
	TypeString:    "string",
	TypeEOL:       "EOL",
}

type Token struct {
	Value    string
	Type     int
	Position int
}

func (t *Token) String() string {
	return fmt.Sprintf("{value: \"%s\", type: %s, pos: %d}", t.Value, MapTypeToTypeName[t.Type], t.Position)
}
