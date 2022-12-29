package lexer

import "monkey-interpreter/internal/token"

type Lexer struct {
	input        string
	position     int  // Current position in input (points to current char).
	readPosition int  // Current reading position in input (after current char).
	ch           byte // Current char under examination.
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	defer l.readChar()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, string(l.ch))
	case ';':
		tok = newToken(token.SEMICOLON, string(l.ch))
	case '(':
		tok = newToken(token.LPAREN, string(l.ch))
	case ')':
		tok = newToken(token.RPAREN, string(l.ch))
	case ',':
		tok = newToken(token.COMMA, string(l.ch))
	case '+':
		tok = newToken(token.PLUS, string(l.ch))
	case '{':
		tok = newToken(token.LBRACE, string(l.ch))
	case '}':
		tok = newToken(token.RBRACE, string(l.ch))
	case 0:
		tok = newToken(token.EOF, "")
	}

	return tok
}

func newToken(tokenType token.TokenType, ch string) token.Token {
	return token.Token{Type: tokenType, Literal: ch}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for the "NUL".
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}
