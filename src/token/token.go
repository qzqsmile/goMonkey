package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	SLASH    = "/"
	ASTERISK = "*"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	LT = "<"
	GT = ">"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

	EQ = "=="
	NOT_EQ = "!="
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true": TRUE,
	"false": FALSE,
}

func LookupIndent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
