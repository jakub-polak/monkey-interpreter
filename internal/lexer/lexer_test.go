package lexer

import (
	"testing"

	"monkey-interpreter/internal/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for _, test := range tests {
		tok := l.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("wrong tokentype: expected=%q, got=%q",
				test.expectedType, tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Fatalf("wrong literal: expected=%q, got=%q",
				test.expectedLiteral, tok.Literal)
		}
	}
}
