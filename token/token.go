package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// 标识符 + 字面量
	IDENT = "IDENT" // add foobar x y
	INT   = "INT"   // 123142531
	// 运算符
	ASSIGN = "="
	PLUS   = "+"
	// 分隔符
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	// 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
