package interpreter

import (
	"RedoLanguage/err"
	lx "RedoLanguage/lexer"
	"RedoLanguage/std"
	"errors"
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
	forl:
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
			return
		} else if secondT.Type == lx.ILLEGAL {
			log.Println(err.NewIllegalTokenError(secondT.Value, lexer.Scanner.Pos()))
			return
		}
		if curT.Type == lx.EOF {
			return
		}


		if curT.Type == lx.RBRACE {
			lexer.CurrentNestingLevel--
		}
		if lx.IsIfStatement(curT) {
			
			switch curT.Type {
			case lx.ELSE:
				if secondT.Type == lx.IF {
					// n := lexer.NextToken()
					// if err := lx.ExecuteIf(secondT, n, lexer); err != nil {
					// 	log.Println(err)
					// 	return
					// }
					secondTS = secondT
					continue forl
					
				}
				if secondT.Type != lx.LBRACE {
					log.Println(errors.New("invalid else statement"))
					return
				}
				curIf := lexer.IfPositions[lexer.CurrentNestingLevel+1]
				b, errs := curIf.Output()
				if errs != nil {
					log.Println(errs)
					return
				}
				tok := secondT
				prevNes := lexer.CurrentNestingLevel
				if b/*b==false*/ {
					for {

						switch tok.Type {
						case lx.LBRACE:
							lexer.CurrentNestingLevel++
						case lx.RBRACE:
							lexer.CurrentNestingLevel--
						case lx.EOF:
							return
						}
						if prevNes == lexer.CurrentNestingLevel {
							break
						}
						tok = lexer.NextToken()
					}
					//skip

				}

			default:
				if err := lx.ExecuteIf(curT, secondT, lexer); err != nil {
					log.Println(err)
					return
				}
			}

			

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
