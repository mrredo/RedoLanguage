package lexer

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func ParseArithmeticExpressions(expression string) (any, error) {
	// Create new expression with default token factory
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return 0, fmt.Errorf("Error parsing expression: %v", err)
	}

	// Evaluate expression with empty parameter map
	result, err := expr.Evaluate(nil)
	if err != nil {
		return 0, fmt.Errorf("Error evaluating expression: %v", err)
	}

	// Convert result to int and return
	if result == "true" || result == "false" {
		return result == "true", nil
	}
	if val, ok := result.(float64); ok {
		return int(val), nil
	}
	return 0, fmt.Errorf("Error converting result to int")
}
func IsMathExpression(curT Token, secondT Token, lexer *Lexer) bool { //10 +

	switch curT.Type {
	case LPAREN:
		switch secondT.Type {
		case IDENTIFIER, NUMBER, BOOL:
			return true
		default:
			return false
		}
	case NUMBER:
		switch secondT.Type {
		//case :
		}
	}
	return false
}

//func ReplaceAllIdentsWithValue(c Token, s Token, l *Lexer) (interface{}, error) {
//	switch c.Type {
//	case IDENTIFIER:
//		if s.Type == RPAREN {
//			funName, args, err := ParseFunctionCall(c, s, l)
//			if err != nil {
//				return nil, err
//			}
//
//		}
//		va, ok := std.Variables[c.Value]
//		if !ok {
//			return nil, fmt.Errorf("'%s' is not defined", c.Value)
//		}
//		return va, nil
//
//	default:
//		return nil, nil
//	}
//}
