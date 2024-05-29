package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	// This simulates some monkey variable declarations
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	lex := lexer.New(input)
	psr := New(lex)

	program := psr.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements expected=3 got=%d", program.Statements)
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, expected string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("stmt.TokenLiteral() not 'let', got=%s", s.TokenLiteral())
		return false
	}

	letSt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement, got=%s", s)
		return false
	}

	if letSt.Name.Value != expected {
		t.Errorf("letSt.Name.Value not %q, got=%s", expected, letSt.Name.Value)
		return false
	}

	if letSt.Name.TokenLiteral() != expected {
		t.Errorf("letSt.Name.TokenLiteral() not %q, got=%s", expected, letSt.Name.TokenLiteral())
		return false
	}

	return true
}
