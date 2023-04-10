package interpreter

import (
	lx "RedoLanguage/lexer"
	"RedoLanguage/std"
	"fmt"
	"log"
	"strings"
)

func Interpret(input string) {
	if input[len(input)-1] != ';' {
		input += ";"
	}

	lexer := lx.NewLexer(strings.ReplaceAll(input, " ", " "))
	var secondTS lx.Token = lexer.NextToken()
	for {

		curT := secondTS
		secondT := lexer.NextToken()
		if curT.Type == lx.EOF {
			break
		}

		if lx.IsVariableExpression(curT, secondT, lexer) { // key +

			key := curT
			exp := secondT
			val := lexer.NextToken()
			_, err := lx.ParseVariableAssigningExpression(key, exp, val, lexer)
			if err != nil {
				log.Println(err)
				break
			}
			/*
				this works for
				+ and +=
			*/

		}

		if lx.IsVariable(curT) {
			fmt.Println(curT, secondT)
			_, _, err := lx.ParseVariable(curT, secondT, lexer)
			if err != nil {
				log.Println(err)
				break
			}
			continue
		}
		if lx.IsFunction(curT, secondT, lexer) {
			funcName, val, err := lx.ParseFunctionCall(curT, secondT, lexer)
			if err != nil {
				log.Println(err)
				break
			}
			if funcM, ok := std.Functions[funcName]; ok {
				funcM(val...)
			}

		}
		secondTS = secondT
	}
}
