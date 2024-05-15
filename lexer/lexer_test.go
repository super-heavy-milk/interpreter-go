package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
`

	cases := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.COMMA, ","},
		{token.EOF, ""},
		{token.LBRACE, "{"},
		{token.LPAREN, "("},
		{token.PLUS, "+"},
		{token.RBRACE, "}"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
	}

	// todo: implement
	tokenReader := New(input)

	for i, tcase := range cases {
		currentToken := tokenReader.NextToken()
		t.Log(currentToken)

		if currentToken.Type != tcase.expectedType {
			t.Logf("tcase[%d] wrong type expected=%q got=%q",
				i, tcase.expectedType, currentToken.Type)
			t.Fail()
		}

		if currentToken.Literal != tcase.expectedLiteral {
			t.Logf("tcase[%d] wrong literal expected=%q got=%q",
				i, tcase.expectedLiteral, currentToken.Literal)
			t.Fail()
		}
	}

	// these test String() to see if panic
	// not great but whatever i'm a noob at Go
	t3 := New("hey how it is going?\n what's up?")
	for i := 0; i < len(t3.input)-1; i++ {
		_ = t3.NextToken()
		t.Log(t3)
	}
	t4 := New("hey\nsup")
	for i := 0; i < len(t4.input)-1; i++ {
		_ = t4.NextToken()
		t.Log(t4)
	}
}
