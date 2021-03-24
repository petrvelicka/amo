package lexer

import (
	"github.com/petrvelicka/amo/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
int five;
string str;
five = 5;
read var;
print var + 5;
ten = 10;
bool x;
x = true;
x = false;
if (ten > five) {
	print "hello";
} else {
	print "not";
}
+-<>/*!
10 == 10;
5 != 4;
`

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.ID_INT, "int"},
		{token.IDENT, "five"},
		{token.SEMICOLON, ";"},
		{token.ID_STRING, "string"},
		{token.IDENT, "str"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "five"},
		{token.ASSIGN,"="},
		{token.INT,"5"},
		{token.SEMICOLON, ";"},
		{token.READ, "read"},
		{token.IDENT, "var"},
		{token.SEMICOLON, ";"},
		{token.PRINT, "print"},
		{token.IDENT, "var"},
		{token.PLUS, "+"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "ten"},
		{token.ASSIGN,"="},
		{token.INT,"10"},
		{token.SEMICOLON, ";"},
		{token.ID_BOOL, "bool"},
		{token.IDENT, "x"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "x"},
		{token.ASSIGN,"="},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "x"},
		{token.ASSIGN,"="},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "ten"},
		{token.GT, ">"},
		{token.IDENT, "five"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.PRINT, "print"},
		{token.STRING, "hello"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.PRINT, "print"},
		{token.STRING, "not"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.BANG, "!"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.NEQ, "!="},
		{token.INT, "4"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong, expected %q, got %q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - token literal wrong, expected %q, got %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
