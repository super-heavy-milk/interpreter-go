package lexer

import (
	"fmt"
	"monkey/token"
	"strings"
)

type Lexer struct {
	input   string // the string to iterate over
	char    byte   // the char under iteration
	charPos int    // index of char
	readPos int    // charIndex + 1
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
		beginSeg := l.charPos < winSz
		endSeg := l.readPos > len(s)-winSz

		switch true {
		// short string
		case beginSeg && endSeg:
			start = s[:l.charPos]
			end = s[l.readPos:]
		// start
		case beginSeg && !endSeg:
			start = s[:l.charPos]
			end = fmt.Sprintf("%s…", s[l.readPos:l.readPos+winSz])
		// middle
		case !beginSeg && !endSeg:
			start = fmt.Sprintf("…%s", s[l.charPos-winSz:l.charPos])
			end = fmt.Sprintf("%s…", s[l.readPos:l.readPos+winSz])
		// end
		case !beginSeg && endSeg:
			start = fmt.Sprintf("…%s", s[l.charPos-winSz:l.charPos])
			end = s[l.readPos:]
		}

		slidingWindow = fmt.Sprintf("%s[%s]%s",
			start, string(s[l.charPos]), end)
	}

	return fmt.Sprintf(
		"char=%-6q charIndex=%-6d readPosition=%-6d input=\"%s\"",
		l.char, l.charPos, l.readPos, slidingWindow)
}

func (l *Lexer) readChar() {
	atStrEnd := l.readPos >= len(l.input)
	if atStrEnd {
		l.char = 0 // ASCII for "NUL"
	} else {
		l.char = l.input[l.readPos]
	}
	l.charPos = l.readPos
	l.readPos += 1
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
