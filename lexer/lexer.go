package lexer

import "monkey/token"

type Lexer struct {
	input        string // the string to interate over
	char         byte   // the char under iteration
	charIndex    int    // index of char
	readPosition int    // charIndex + 1
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // initialize
	return l
}

func (l *Lexer) readChar() {
	atStrEnd := l.readPosition >= len(l.input)
	if atStrEnd {
		l.char = 0 // ASCII for "NUL"
	} else {
		l.char = l.input[l.readPosition]
	}
	l.charIndex = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var toke token.Token

	switch l.char {
	case '=':
		toke = newToken(token.ASSIGN, l.char)
	case ';':
		toke = newToken(token.SEMICOLON, l.char)
	case '(':
		toke = newToken(token.LBRACE, l.char)
	case '{':
		toke = newToken(token.LPAREN, l.char)
	case ')':
		toke = newToken(token.RPAREN, l.char)
	case '}':
		toke = newToken(token.RBRACE, l.char)
	case '+':
		toke = newToken(token.PLUS, l.char)
	case ',':
		toke = newToken(token.COMMA, l.char)
	case 0:
		toke.Literal = ""
		toke.Type = token.EOF
	}

	l.readChar()
	return toke
}

func newToken(tt token.TokenType, b byte) token.Token {
	return token.Token{Type: tt, Literal: string(b)}
}
