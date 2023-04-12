package lexer

func IsVariable(token Token) bool {
	return token.Type == VAR
}

