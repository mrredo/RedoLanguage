package interpreter

import (
	"RedoLanguage/err"
	lx "RedoLanguage/lexer"
	"RedoLanguage/std"
	"errors"
	"fmt"
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

	//if lexer.SemErr != nil {
	//	log.Println(lexer.SemErr)
	//	return
	//}
	// var curPos, secPos = lexer.Scanner.Pos(), lexer.Scanner.Pos()
	for {

		curT = secondTS
		// curPos = lexer.Scanner.Pos()

		secondT := lexer.NextToken()
		//if lexer.SemErr != nil {
		//	log.Println(lexer.SemErr)
		//	break
		//}
		// secPos = lexer.Scanner.Pos()
		if curT.Type == lx.ILLEGAL {
			log.Println(err.NewIllegalTokenError(curT.Value, lexer.Scanner.Pos()))
			break
		} else if secondT.Type == lx.ILLEGAL {
			log.Println(err.NewIllegalTokenError(secondT.Value, lexer.Scanner.Pos()))
			break
		}
		if curT.Type == lx.EOF {
			break
		}

		//later will be added when I want to fix it
		//if curT.Type == lx.IDENTIFIER {
		//	_, ok := std.Variables[curT.Value]
		//	_, ok1 := std.Functions[curT.Value]
		//
		//	if !ok && !ok1 {
		//		errs := err.NewUndefinedError(curT.Value, lexer.Scanner.Pos())
		//		log.Println(errs)
		//		break
		//	}
		//	// if curPos.Line != secPos.Line {
		//	// 	errs := err.NewUnusedError(curT.Value, lexer.Scanner.Pos())
		//	// 	log.Println(errs)
		//	// 	break
		//	// }
		//
		//}
		if curT.Type == lx.RBRACE {
			lexer.CurrentNestingLevel--
		}
		if lx.IsIfStatement(curT) {

			tok := secondT
			i := lx.If{Position: lexer.CurrentPosition, NestingLevel: lexer.CurrentNestingLevel}
		forif:
			for {
				switch tok.Type {
				case lx.EOF:
					break forif
				case lx.LBRACE:
					lexer.CurrentPosition++
					lexer.CurrentNestingLevel++
					break forif
				}

				i.Condition += tok.Value
				tok = lexer.NextToken()
			}

			fmt.Println(i)
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
		// if curT.Type == lx.IDENTIFIER {
		// 	_, ok := std.Variables[curT.Value]
		// 	if !ok {
		// 		errS := err.NewUndefinedError(curT.Value, curT.Position)
		// 		log.Println(errS)
		// 		break
		// 	}
		// 	errS := err.NewUnusedError(curT.Value, curT.Position)
		// 	log.Println(errS)
		// 	break
		// }

		secondTS = secondT

	}
	if lexer.CurrentNestingLevel < 0 {
		log.Println(errors.New("missing start of statement"))
	}
	if lexer.CurrentNestingLevel > 0 {
		log.Println(errors.New("missing end of statement"))

	}
}
