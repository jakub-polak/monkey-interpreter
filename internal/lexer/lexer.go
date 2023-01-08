package lexer

import (
	"monkey-interpreter/internal/token"
)

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
	var t token.Token

	l.skipWhitespaces()

	switch l.ch {
	case '=':
		t = newToken(token.ASSIGN, string(l.ch))
	case ';':
		t = newToken(token.SEMICOLON, string(l.ch))
	case '(':
		t = newToken(token.LPAREN, string(l.ch))
	case ')':
		t = newToken(token.RPAREN, string(l.ch))
	case ',':
		t = newToken(token.COMMA, string(l.ch))
	case '+':
		t = newToken(token.PLUS, string(l.ch))
	case '-':
		t = newToken(token.MINUS, string(l.ch))
	case '!':
		t = newToken(token.BANG, string(l.ch))
	case '/':
		t = newToken(token.SLASH, string(l.ch))
	case '*':
		t = newToken(token.ASTERISK, string(l.ch))
	case '<':
		t = newToken(token.LT, string(l.ch))
	case '>':
		t = newToken(token.GT, string(l.ch))
	case '{':
		t = newToken(token.LBRACE, string(l.ch))
	case '}':
		t = newToken(token.RBRACE, string(l.ch))
	case 0:
		t = newToken(token.EOF, "")
	default:
		if isLetter(l.ch) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdent(t.Literal)

			return t
		} else if isDigit(l.ch) {
			t.Type = token.INT
			t.Literal = l.readNumber()
			return t
		} else {
			t = newToken(token.ILLEGAL, string(l.ch))
		}
	}

	l.readChar()
	return t
}

func newToken(tokenType token.Type, ch string) token.Token {
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

func (l *Lexer) readIdentifier() string {
	position := l.position // Save the starting position of an identifier.

	for isLetter(l.ch) {
		l.readChar() // Read all letters.
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position // Save the starting position of a number.

	for isDigit(l.ch) { // Read all digits.
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespaces() {
	for isSpace(l.ch) {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}
