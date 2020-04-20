package lexer

import (
	"testing"

	"github.com/ASteinheiser/amoeba-interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := `
		let four = 4;
		let seventeen = 17;

		let add = fn(a, b) {
			a + b;
		};

		let result = add(four, seventeen);

		!-/*5;
		5 < 10 > 5;

		if (5 < 10) {
			return true;
		} else {
			return false;
		}

		14 == 14;
		18 != 6;

		"something"
		"some string"

		[1, 2, 3];

		{ "key": "value" }
	`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "four"},
		{token.ASSIGN, "="},
		{token.INT, "4"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "seventeen"},
		{token.ASSIGN, "="},
		{token.INT, "17"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "a"},
		{token.COMMA, ","},
		{token.IDENT, "b"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "a"},
		{token.PLUS, "+"},
		{token.IDENT, "b"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "four"},
		{token.COMMA, ","},
		{token.IDENT, "seventeen"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "14"},
		{token.EQ, "=="},
		{token.INT, "14"},
		{token.SEMICOLON, ";"},
		{token.INT, "18"},
		{token.NOT_EQ, "!="},
		{token.INT, "6"},
		{token.SEMICOLON, ";"},
		{token.STRING, "something"},
		{token.STRING, "some string"},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.COMMA, ","},
		{token.INT, "3"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},
		{token.LBRACE, "{"},
		{token.STRING, "key"},
		{token.COLON, ":"},
		{token.STRING, "value"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, test := range tests {
		token := l.NextToken()

		if token.Type != test.expectedType {
			t.Fatalf("test [%d] - tokentype wrong. expected=%q, got=%q",
				i, test.expectedType, token.Type)
		}

		if token.Literal != test.expectedLiteral {
			t.Fatalf("test [%d] - literal wrong. expected=%q, got=%q",
				i, test.expectedLiteral, token.Literal)
		}
	}
}
