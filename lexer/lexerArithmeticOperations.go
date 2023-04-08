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

func ReplaceAllIdentsWithValue(c Token, s Token, l *Lexer) (interface{}, error) {
	return nil, nil
}
func MathExpressionTokensToEnd(c Token, s Token, l *Lexer) ([]Token, error) {
	var tokenArr []Token
	RPcount, LPcount := 0, 0

	if s.Type == RPAREN {
		tokenArr = append(tokenArr, s)
		RPcount++
	}
	switch c.Type {
	case IDENTIFIER:
		if s.Type == LPAREN {
			//function
		}

		//variable
	case BOOL, NUMBER:
		tokenArr = append(tokenArr, c)
	case LPAREN:
		tokenArr = append(tokenArr, c)
		LPcount++

	case RPAREN:
		return nil, fmt.Errorf("first token to a math expression can't be a ')'")
	}
	if LPcount != RPcount {
		return nil, fmt.Errorf("invalid left/right parentheses count")
	}
	return tokenArr, nil
}
