package lexer

import "fmt"

func IsVariable(token Token) bool {
	if token.Type == VAR {
		return true
	}
	return false
}

func TestIsVariable() {
	lx := NewLexer(`
var hello = "hello world"
`)
	curT := lx.NextToken()
	fmt.Println(IsVariable(curT))
	fmt.Println(ParseVariable(curT, lx))
	//fmt.Println(ParseVariable(lx))
}
