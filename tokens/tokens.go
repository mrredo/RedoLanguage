package tokens
 import lx "RedoLanguage/lexer"

var keywords = map[string]lx.TokenType{
	"if": lx.IF,
	"else if": lx.ELSE_IF,
	"else": lx.ELSE,
	"var": lx.VAR,
	"true": lx.BOOL,
	"false": lx.BOOL,
}