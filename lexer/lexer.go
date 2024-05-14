package lexer

import (
	"fmt"
	"monkey/token"
	"strings"
)

type Lexer struct {
	input        string // the string to iterate over
	char         byte   // the char under iteration
	charIndex    int    // index of char
	readPosition int    // charIndex + 1
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	fmt.Printf("Initializing Lexer: %s\n", l)
	l.readChar() // initialize
	return l
}

func (l Lexer) String() string {
	// input is unweidly to print, so create a sliding window
	// with a nice "…e's the ⁅c⁆hara…" format
	var slidingWindow string
	{
		s := l.input
		s = strings.ReplaceAll(s, "\n", " ")

		var start string
		if l.charIndex < 6 {
			start = s[:l.charIndex]
		} else {
			start = fmt.Sprintf("…%s", s[l.charIndex-6:l.charIndex])
		}

		var end string
		if l.readPosition >= len(l.input)-6 {
			end = s[l.readPosition:]
		} else {
			end = fmt.Sprintf("%s…", s[l.readPosition:l.readPosition+6])
		}

		slidingWindow = fmt.Sprintf("%s⁅%s⁆%s",
			start, string(l.char), end)
	}

	return fmt.Sprintf(
		"char=%q charIndex=%d readPosistion=%d input=\"%s\"",
		l.char, l.charIndex, l.readPosition, slidingWindow)
}

func (l *Lexer) readChar() {
	fmt.Printf("lexer enter -> %s\n", l)
	atStrEnd := l.readPosition >= len(l.input)
	if atStrEnd {
		l.char = 0 // ASCII for "NUL"
	} else {
		l.char = l.input[l.readPosition]
	}
	l.charIndex = l.readPosition
	l.readPosition += 1
	fmt.Printf("lexer exit -> %s\n", l)
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
	fmt.Printf("%s\n", toke)

	l.readChar()
	return toke
}

func newToken(tt token.TokenType, char byte) token.Token {
	return token.Token{Type: tt, Literal: string(char)}
}
