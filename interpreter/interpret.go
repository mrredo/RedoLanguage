package interpreter

import (
	lx "RedoLanguage/lexer"
	"RedoLanguage/std"
	"log"
)

func Interpret(input string) {
	lexer := lx.NewLexer(input)
	for {
		curT := lexer.NextToken()
		if curT.Type == lx.EOF {
			break
		}
		if lx.IsVariable(curT) {
			_, _, err := lx.ParseVariable(curT, lexer)
			if err != nil {
				log.Fatal(err)
				break
			}
			continue
		}
		if lx.IsFunction(curT, lexer) {
			funcName, val, err := lx.ParseFunctionCall(curT, lexer)
			if err != nil {
				log.Fatal(err)
				break
			}
			if funcM, ok := std.Functions[funcName]; ok {
				funcM(val...)
			}

		}
	}
}
