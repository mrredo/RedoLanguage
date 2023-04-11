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

func MathExpressionTokensToEnd(c Token, l *Lexer, function ...bool) (string, Token, error) {

	var finalStr string
	var firstfirstTok = c
	var firstfirstExpired = false
	var curType TokenType = -1
	var nestingLevel int
	//var nestingLevelModified bool
	for {
		if c.Type == SEMICOLON || c.Type == NEW_LINE || c.Type == EOF /*|| c.Type == COMMA*/ {
			break
		}
		//if p := l.Scanner.Pos(); p.Offset == len(l.Input)-2 && len(function) >= 1 {
		//	break
		//}
		if firstfirstTok.Type == COMMA && !firstfirstExpired {
			nestingLevel++
			finalStr += "("
			firstfirstExpired = true
		}

		if p := l.Scanner.Peek(); p == ';' || p == '\n' {
			break
		}
		switch c.Type {
		case COMMA:
			fmt.Println(firstfirstExpired)
			if firstfirstExpired {
				nestingLevel--
				finalStr += ")"
				fmt.Println(nestingLevel)
				if nestingLevel == 0 {
					break
				}
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
				if vas, ok1 := va.(string); ok1 {
					finalStr += fmt.Sprintf(`"%s"`, vas)
				} else {
					finalStr += fmt.Sprint(va)
				}

				if curType == -1 {
					curType = ConvertToTokenType(reflect.TypeOf(fmt.Sprint(va)).String())
				}
				if curType != ConvertToTokenType(reflect.TypeOf(fmt.Sprint(va)).String()) {
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
			if nestingLevel < 0 {
				// unbalanced parentheses
				return "", c, err.NewSyntaxError(err.UnbalancedParentheses, l.Scanner.Pos())
			}
			finalStr += ")"
		default:

			finalStr += c.Value
		}
		if nestingLevel == 0 {
			break
		}
		c = l.NextToken()

		//if c.Type != RPAREN || c.Type != LPAREN {
		//	OperatorTurn = !OperatorTurn

	}
	//if LPcount != RPcount {
	//	return "", fmt.Errorf("invalid left/right parentheses count")
	//}
	return finalStr, c, nil
}
