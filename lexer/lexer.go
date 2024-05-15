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

func (l *Lexer) String() string {
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
	var t token.Token

	l.skipWhitespace()

	switch l.char {
	case '=':
		if l.peekChar() == '=' {
			ch := l.char
			l.readChar()
			literal := string(ch) + string(l.char)
			t = token.Token{Type: token.EQ, Literal: literal}
		} else {
			t = newToken(token.ASSIGN, l.char)
		}
	case ';':
		t = newToken(token.SEMICOLON, l.char)
	case '(':
		t = newToken(token.LPAREN, l.char)
	case '{':
		t = newToken(token.LBRACE, l.char)
	case ')':
		t = newToken(token.RPAREN, l.char)
	case '}':
		t = newToken(token.RBRACE, l.char)
	case '+':
		t = newToken(token.PLUS, l.char)
	case ',':
		t = newToken(token.COMMA, l.char)
	case '*':
		t = newToken(token.ASTERISK, l.char)
	case '!':
		if l.peekChar() == '=' {
			ch := l.char
			l.readChar()
			literal := string(ch) + string(l.char)
			t = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			t = newToken(token.BANG, l.char)
		}
	case '/':
		t = newToken(token.SLASH, l.char)
	case '<':
		t = newToken(token.LT, l.char)
	case '>':
		t = newToken(token.GT, l.char)
	case '-':
		t = newToken(token.MINUS, l.char)
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isLetter(l.char) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdentifier(t.Literal)
			return t
		} else if isDigit(l.char) {
			t.Literal = l.readNumber()
			t.Type = token.INT
			return t
		} else {
			t = newToken(token.ILLEGAL, l.char)
		}
	}
	// fmt.Printf("%s\n", toke)

	l.readChar()
	return t
}

// readIdentifier advances the lexer till the end of a letter sequence,
// and then returns the letter sequence.
func (l *Lexer) readIdentifier() string {
	idStart := l.charPos
	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[idStart:l.charPos]
}

// readNumber the lexer till the end of a digit sequence,
// and then returns the digit sequence.
func (l *Lexer) readNumber() string {
	numStart := l.charPos
	for isDigit(l.char) {
		l.readChar()
	}
	return l.input[numStart:l.charPos]
}

func (l *Lexer) peekChar() byte {
	atEnd := l.readPos >= len(l.input)
	if atEnd {
		return 0
	} else {
		return l.input[l.readPos]
	}
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' ||
		l.char == '\t' ||
		l.char == '\n' ||
		l.char == '\r' {
		l.readChar()
	}
}

func newToken(tt token.TokenType, char byte) token.Token {
	return token.Token{Type: tt, Literal: string(char)}
}

func isLetter(char byte) bool {
	return 'a' <= char && char >= 'z' ||
		'A' <= char && char >= 'Z' ||
		char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}
