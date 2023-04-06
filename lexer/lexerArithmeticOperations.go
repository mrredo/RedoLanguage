package lexer

func ParseArithmeticExpression(curT Token, lexer *Lexer) {

}

var precedence = map[TokenType]int{
	MULTIPLY: 3,
	DIVIDE:   3,
	MODULO:   3,
	PLUS:     2,
	SUBTRACT: 2,
	LPAREN:   1,
	RPAREN:   0,
}
