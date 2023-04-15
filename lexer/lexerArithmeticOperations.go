package lexer

import (
	"RedoLanguage/err"
	"RedoLanguage/std"
	"fmt"
	"github.com/Knetic/govaluate"
	"reflect"
)

func ParseArithmeticExpressions(expression string, l *Lexer) (any, error) {
	// Create new expression with default token factory

	expr, errs := govaluate.NewEvaluableExpression(expression)
	if errs != nil {
		return 0, err.NewExpressionError(err.ErrorParsingExpression, errs.Error(), l.Scanner.Pos()) //fmt.Errorf("error parsing expression: %v", err)
	}

	// Evaluate expression with empty parameter map
	result, errss := expr.Evaluate(nil)
	if errss != nil {
		return 0, err.NewExpressionError(err.ErrorEvaluatingExpression, errss.Error(), l.Scanner.Pos()) //fmt.Errorf("error evaluating expression: %v", err)
	}

	// Convert result to int and return
	if val, ok := result.(string); ok {
		return val, nil
	}
	if result == true || result == false {
		return result, nil
	}
	if val, ok := result.(float64); ok {
		return int(val), nil
	}
	return 0, err.NewExpressionError(err.ErrorConvertingResultToInt, "", l.Scanner.Pos()) //fmt.Errorf("error converting result to int")
}
func isOperator(op Token) bool {
	_, ok := OperatorNumToString[op.Type]
	return ok
}
func MathExpressionTokensToEnd(c Token, l *Lexer, function ...bool) (string, Token, error) {

	var finalStr string
	var firstfirstTok = c
	var firstfirstExpired = false
	var curType TokenType = -1
	var nestingLevel int

	//var isOperatorTurn bool
	//var nestingLevelModified bool
loop:
	for {
		if c.Type == SEMICOLON || c.Type == NEW_LINE || c.Type == EOF /*|| c.Type == COMMA*/ {
			break
		}
		//fmt.Println(isOperator(c), isOperatorTurn, c)
		//if isOperator(c) {
		//	if !isOperatorTurn {
		//		break
		//
		//	}
		//	isOperatorTurn = false
		//
		//} else {
		//	if isOperatorTurn {
		//		break
		//	}
		//	isOperatorTurn = true
		//}

		if len(function) > 0 {
			if firstfirstTok.Type == COMMA && !firstfirstExpired {
				nestingLevel++
				finalStr += "("
				firstfirstExpired = true
				c = l.NextToken()
				continue loop
			}

		}

		// if p := l.Scanner.Peek(); p == ';' || p == '\n' {
		// 	break
		// }
		switch c.Type {
		case COMMA:
			nestingLevel--
			finalStr += ")"
			if nestingLevel == 0 {
				break loop
			}

		case STRING:
			if curType == -1 {
				curType = c.Type
			}
			if curType != c.Type {
				return "", c, err.NewTypeError(l.Scanner.Pos())
			}
			finalStr += c.Value
		case BOOL, NUMBER:

			if curType == -1 {
				curType = c.Type
			}
			if curType != c.Type {
				return "", c, err.NewTypeError(l.Scanner.Pos())
			}
			finalStr += c.Value

		case IDENTIFIER:
			if p := l.Scanner.Peek(); p == '(' {
				s := l.NextToken()
				f, args, errs := ParseFunctionCall(c, s, l)
				if errs != nil {
					return "", c, errs
				}
				out, ok := std.Functions[f]

				if !ok {
					return "", c, err.NewUndefinedError(c.Value, l.Scanner.Pos()) //fmt.Errorf("'%s' function is not defined", c.Value)
				}
				o := out(args...)

				if o == nil {
					return "", c, err.NewFunctionVoidError(c.Value, l.Scanner.Pos()) //fmt.Errorf("invalid function call: '%s' function does not return a value", c.Value)
				}
				if vas, ok1 := o.(string); ok1 {
					finalStr += fmt.Sprintf(`"%s"`, vas)
				} else {
					finalStr += fmt.Sprint(o)
				}

				if curType == -1 {
					curType = ConvertToTokenType(reflect.TypeOf(fmt.Sprint(o)).String())
				}
				if curType != ConvertToTokenType(reflect.TypeOf(fmt.Sprint(o)).String()) {
					if curType != c.Type {
						return "", c, err.NewTypeError(l.Scanner.Pos())
					}
				}

				c = l.NextToken()

				continue
			} else {
				va, ok := std.Variables[c.Value]
				if !ok {
					return "", c, err.NewUndefinedError(c.Value, l.Scanner.Pos()) //fmt.Errorf("'%s' is not defined", c.Value)
				}
				if vas, ok1 := va.Value.(string); ok1 {
					finalStr += fmt.Sprintf(`"%s"`, vas)
				} else {
					finalStr += fmt.Sprint(va.Value)
				}

				if curType == -1 {
					curType = ConvertToTokenType(reflect.TypeOf(fmt.Sprint(va.Value)).String())
				}
				if curType != ConvertToTokenType(reflect.TypeOf(fmt.Sprint(va.Value)).String()) {
					if curType != c.Type {
						return "", c, err.NewTypeError(l.Scanner.Pos())
					}

				}
			}
		case OR, AND:
			curType = -1
			finalStr += c.Value
		case LPAREN:
			//nestingLevelModified = true
			nestingLevel++
			finalStr += "("
		case RPAREN:
			nestingLevel--
			if nestingLevel != 0 {

				return "", c, err.NewSyntaxError(err.UnbalancedParentheses, l.Scanner.Pos())
			}
			finalStr += ")"
		default:

			finalStr += c.Value
		}
		if len(function) > 0 {
			if nestingLevel == 0 {
				break
			}
		}

		c = l.NextToken()

	}

	return finalStr, c, nil
}
