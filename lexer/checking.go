package lexer

func IsVariable(token Token) bool {
	return token.Type == VAR
}

//func TestIsVariable() {
//	lx := NewLexer(`
//var hello = "hello world"
//`)
//	curT := lx.NextToken()
//	fmt.Println(IsVariable(curT))
//	fmt.Println(ParseVariable(curT, lx))
//	//fmt.Println(ParseVariable(lx))
//}
