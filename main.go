package main

import (
	"RedoLanguage/interpreter"
)

func main() {
	//	interpreter.Interpret(`
	//var hel = 10
	////hello
	///*
	//eeee
	//*/
	//print(modulo(hel,3))
	//	`)
	interpreter.Interpret(`
	hello *= 20
print(hello)
`)

	//	lexer := lx.NewLexer(`
	//hello %= 2
	//`)
	//	key := lexer.NextToken()
	//	expression := lexer.NextToken()
	//	value := lexer.NextToken()
	//
	//	fmt.Println(lx.ParseVariableAssigningExpression(key, expression, value, lexer))
}
