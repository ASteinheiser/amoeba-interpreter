package ast

import (
	"testing"

	"github.com/ASteinheiser/amoeba-interpreter/token"
)

func TestString(t *testing.T) {
	expectedString := "let variableName = otherGreatVar;"

	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "variableName"},
					Value: "variableName",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "otherGreatVar"},
					Value: "otherGreatVar",
				},
			},
		},
	}

	if program.String() != expectedString {
		t.Errorf("program.String() failed. got=%q", program.String())
	}
}
