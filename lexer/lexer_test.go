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
}
