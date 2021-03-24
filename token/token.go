package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	IDENT = "IDENT"
	INT = "INT"
	STRING = "STRING"

	ASSIGN = "="

	PLUS = "+"
	MINUS = "-"
	GT = ">"
	LT = "<"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"
	EQ = "=="
	NEQ = "!="

	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	LBRACE = "{"
	RBRACE = "}"

	READ = "READ"
	PRINT = "PRINT"

	IF = "IF"
	ELSE = "ELSE"

	ID_INT = "ID_INT"
	ID_STRING = "ID_STRING"
	ID_BOOL = "ID_BOOL"

	TRUE = "TRUE"
	FALSE = "FALSE"
)

var keywords = map[string]TokenType {
	"read": READ,
	"print": PRINT,
	"if": IF,
	"else": ELSE,
	"int": ID_INT,
	"string": ID_STRING,
	"bool": ID_BOOL,
	"true": TRUE,
	"false": FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}