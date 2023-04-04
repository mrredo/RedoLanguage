package interpreter

import (
	lx "RedoLanguage/lexer"
	"RedoLanguage/std"
	"log"
	"strings"
)

func Interpret(input string) {
	lexer := lx.NewLexer(strings.ReplaceAll(input, " ", ""))
	for {

		curT := lexer.NextToken()
		if curT.Type == lx.EOF {
			break
		}
		if lx.IsVariableExpression(curT, lexer) { // hello +
			key := curT
			exp := lexer.NextToken()
			val := lexer.NextToken()
			_, err := lx.ParseVariableAssigningExpression(key, exp, val, lexer)
			if err != nil {
				log.Fatal(err)
				break
			}
			/*
				this works for
				+ and +=
			*/

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
