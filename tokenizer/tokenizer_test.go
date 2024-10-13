package tokenizer

import (
	"testing"
)

func TestWhenSimpleAddition(t *testing.T) {
	expression := "1 + 2"

	tokenizer := Tokenizer{}

	stream, err := tokenizer.ExpressionToStream(&expression)

	if err != nil {
		t.Fatal(err)

		return
	}

	expected := []Token{
		{
			Value: "1",
			Type:  TypeNumber,
		},
		{
			Value: "+",
			Type:  TypeOperation,
		},
		{
			Value: "2",
			Type:  TypeNumber,
		},
	}

	for i, expectedToken := range expected {
		actualToken := *stream.Tokens[i]

		if expectedToken.Value != actualToken.Value {
			t.Fatalf("Expected: %v\nActual: %v", expectedToken.Value, actualToken.Value)
		}

		if expectedToken.Type != actualToken.Type {
			t.Fatalf("Expected type: %v\nActual type: %v", expectedToken.Type, actualToken.Type)
		}
	}
}
