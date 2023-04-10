package lexer

import (
	"RedoLanguage/err"
	"RedoLanguage/std"
	"fmt"
	"github.com/Knetic/govaluate"
	"reflect"
)

func ParseArithmeticExpressions(expression string) (any, error) {
	// Create new expression with default token factory

	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return 0, fmt.Errorf("error parsing expression: %v", err)
	}

	// Evaluate expression with empty parameter map
	result, err := expr.Evaluate(nil)
	if err != nil {
		return 0, fmt.Errorf("error evaluating expression: %v", err)
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
	return 0, fmt.Errorf("error converting result to int")
}
func MathExpressionTokensToEndFunctionArgument(c Token, l *Lexer) (string, Token, error) {
	var tokenArr []Token

	var RPcount = 0
	var LPcount = 0
	var finalStr string
	for {
		if c.Type == SEMICOLON || c.Type == NEW_LINE || c.Type == EOF || c.Type == COMMA {
			break
		}
		if p := l.Scanner.Pos(); p.Offset == len(l.Input)-2 {
			break
		}
		if p := l.Scanner.Peek(); p == ';' || p == '\n' {
			break
		}

		switch c.Type {
		case IDENTIFIER:
			if p := l.Scanner.Peek(); p == '(' {
				s := l.NextToken()
				f, args, err := ParseFunctionCall(c, s, l)
				if err != nil {
					return "", c, err
				}
				out, ok := std.Functions[f]

				if !ok {
					return "", c, fmt.Errorf("'%s' function is not defined", c.Value)
				}
				o := out(args...)
				if o == nil {
					return "", c, fmt.Errorf("invalid function call: '%s' function does not return a value", c.Value)
				}
				finalStr += fmt.Sprint(o)
				c = l.NextToken()
				continue
			} else {
				va, ok := std.Variables[c.Value]
				if !ok {
					return "", c, fmt.Errorf("'%s' is not defined", c.Value)
				}
				finalStr += fmt.Sprint(va)
				tokenArr = append(tokenArr, c)
			}

		case LPAREN:
			LPcount++
			finalStr += "("
			tokenArr = append(tokenArr, c)
		case RPAREN:

			RPcount++
			finalStr += ")"
			tokenArr = append(tokenArr, c)
		default:
			finalStr += c.Value
			tokenArr = append(tokenArr, c)
		}

		c = l.NextToken()
		//if c.Type != RPAREN || c.Type != LPAREN {
		//	OperatorTurn = !OperatorTurn
		//}
	}
	//if LPcount != RPcount {
	//	return "", fmt.Errorf("invalid left/right parentheses count")
	//}
	return finalStr, c, nil
}

func MathExpressionTokensToEnd(c Token, l *Lexer, function ...bool) (string, Token, error) {
	var tokenArr []Token

	var RPcount = 0
	var LPcount = 0
	var finalStr string
	var curType TokenType = -1

	for {
		if c.Type == SEMICOLON || c.Type == NEW_LINE || c.Type == EOF || c.Type == COMMA {
			break
		}
		fmt.Println(c)
		if p := l.Scanner.Pos(); p.Offset == len(l.Input)-2 && len(function) >= 1 {
			break
		}
		if p := l.Scanner.Peek(); p == ';' || p == '\n' {
			break
		}
		switch c.Type {
		case STRING:
			if curType == -1 {
				curType = c.Type
			}
			if curType != c.Type {
				return "", c, err.NewTypeError(l.Scanner.Pos())
			}
			finalStr += c.Value
			tokenArr = append(tokenArr, c)
		case BOOL, NUMBER:

			if curType == -1 {
				curType = c.Type
			}
			if curType != c.Type {
				return "", c, err.NewTypeError(l.Scanner.Pos())
			}
			finalStr += c.Value
			tokenArr = append(tokenArr, c)

		case IDENTIFIER:
			if p := l.Scanner.Peek(); p == '(' {
				s := l.NextToken()
				f, args, errs := ParseFunctionCall(c, s, l)
				if errs != nil {
					return "", c, errs
				}
				out, ok := std.Functions[f]

				if !ok {
					return "", c, fmt.Errorf("'%s' function is not defined", c.Value)
				}
				o := out(args...)
				if o == nil {
					return "", c, fmt.Errorf("invalid function call: '%s' function does not return a value", c.Value)
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
					return "", c, fmt.Errorf("'%s' is not defined", c.Value)
				}
				if vas, ok1 := va.(string); ok1 {
					finalStr += fmt.Sprintf(`"%s"`, vas)
				} else {
					finalStr += fmt.Sprint(va)
				}

				tokenArr = append(tokenArr, c)
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
			tokenArr = append(tokenArr, c)
		case LPAREN:
			LPcount++
			finalStr += "("
			tokenArr = append(tokenArr, c)
		case RPAREN:

			RPcount++
			finalStr += ")"
			tokenArr = append(tokenArr, c)
		default:

			finalStr += c.Value
			tokenArr = append(tokenArr, c)
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
