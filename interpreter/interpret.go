package interpreter

import (
	"RedoLanguage/err"
	lx "RedoLanguage/lexer"
	"RedoLanguage/std"
	"log"
	"strings"
)

func Interpret(input string, fileName string) {
	// if input[len(input)-1] != ';' {
	// 	input += ";"
	// }

	lexer := lx.NewLexer(strings.ReplaceAll(input, " ", " "))
	lexer.Scanner.Filename = fileName
	var secondTS lx.Token = lexer.NextToken()
	curT := secondTS
	// var curPos, secPos = lexer.Scanner.Pos(), lexer.Scanner.Pos()
	for {

		curT = secondTS
		// curPos = lexer.Scanner.Pos()

		secondT := lexer.NextToken()
		// secPos = lexer.Scanner.Pos()
		if curT.Type == lx.EOF {
			break
		}
		if curT.Type == lx.IDENTIFIER {
			_, ok := std.Variables[curT.Value]
			_, ok1 := std.Functions[curT.Value]

			if !ok && !ok1 {
				errs := err.NewUndefinedError(curT.Value, lexer.Scanner.Pos())
				log.Println(errs)
				break
			}
			// if curPos.Line != secPos.Line {
			// 	errs := err.NewUnusedError(curT.Value, lexer.Scanner.Pos())
			// 	log.Println(errs)
			// 	break
			// }

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
			_, _, err := lx.ParseVariable(curT, secondT, lexer)
			if err != nil {
				log.Println(err)
				break
			}

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
