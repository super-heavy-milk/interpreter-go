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
	// with a nice "…e's the [c]hara…" format
	var slidingWindow string
	{
		s := l.input
		s = strings.ReplaceAll(s, "\n", " ")
		winSz := 10

		var start string
		var end string
		beginSeg := l.charIndex < winSz
		endSeg := l.readPosition > len(s)-winSz

		switch true {
		// short string
		case beginSeg && endSeg:
			start = s[:l.charIndex]
			end = s[l.readPosition:]
		// start
		case beginSeg && !endSeg:
			start = s[:l.charIndex]
			end = fmt.Sprintf("%s…", s[l.readPosition:l.readPosition+winSz])
		// middle
		case !beginSeg && !endSeg:
			start = fmt.Sprintf("…%s", s[l.charIndex-winSz:l.charIndex])
			end = fmt.Sprintf("%s…", s[l.readPosition:l.readPosition+winSz])
		// end
		case !beginSeg && endSeg:
			start = fmt.Sprintf("…%s", s[l.charIndex-winSz:l.charIndex])
			end = s[l.readPosition:]
		}

		slidingWindow = fmt.Sprintf("%s[%s]%s",
			start, string(s[l.charIndex]), end)
	}

	return fmt.Sprintf(
		"char=%-6q charIndex=%-6d readPosition=%-6d input=\"%s\"",
		l.char, l.charIndex, l.readPosition, slidingWindow)
}

func (l *Lexer) readChar() {
	// fmt.Printf("lexer enter -> %s\n", l)
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
	// fmt.Printf("%s\n", toke)

	l.readChar()
	return toke
}

func newToken(tt token.TokenType, char byte) token.Token {
	return token.Token{Type: tt, Literal: string(char)}
}
