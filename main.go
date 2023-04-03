package main

import (
	lx "RedoLanguage/lexer"
	"fmt"
)

func main() {
	//	lexer := lx.NewLexer(`
	//println(modulo(10, 2),subtract(10, 2), multiply(10, 2), divide(10, 2), add(10,2))
	//println(modulo(add(divide(multiply(10, 100), 1000000), 2),2))
	//
	//`)
	lexer := lx.NewLexer(`
var hello = "hello"
`)

	fmt.Println(lx.ParseVariable(lexer))
	fmt.Println(lx.Variables["hello"])
	//for {
	//	// Parse the next function call
	//
	//	fmt.Println(lexer.NextToken())
	//	if lexer.NextToken().Type == 0 {
	//		break
	//	}
	//	//funcName, args, err := lx.ParseFunctionCall(lexer)
	//	//
	//	//if err == io.EOF {
	//	//	break // End of input
	//	//}
	//	//if err != nil {
	//	//	log.Fatal(err)
	//	//}
	//	//if funcName == "" {
	//	//	break
	//	//}
	//	//// Look up the function and call it with the arguments
	//	//fn, ok := std.Functions[funcName]
	//	//if !ok {
	//	//
	//	//	log.Fatalf("undefined function '%s'", funcName)
	//	//}
	//	//fn(args...)
	//}
}
