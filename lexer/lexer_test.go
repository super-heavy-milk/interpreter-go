package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

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
	tokenReader := token.New(input)

	for i, tcase := range cases {
		currentToken := tokenReader.NextToken()

		if currentToken.Type != tcase.expectedType {
			t.Fatalf("tcase[%d] wrong type expected=%q got=%q",
				i, tcase.expectedType, currentToken.Type)
		}

		if currentToken.Literal != tcase.expectedLiteral {
			t.Fatalf("tcase[%d] wrong literal expected=%q got=%q",
				i, tcase.expectedLiteral, currentToken.Literal)
		}
	}
}
