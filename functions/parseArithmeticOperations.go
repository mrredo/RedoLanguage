package functions

import (
	"RedoLanguage/lexer"
	"RedoLanguage/err"
	"github.com/Knetic/govaluate"
)

func ParseArithmeticExpressions(expression string, l *lexer.Lexer) (any, error) {
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
	if val, ok := result.(string); ok {
		return val, nil
	}
	if result == true || result == false {
		return result, nil
	}
	if val, ok := result.(float64); ok {
		return int(val), nil
	}
	// Convert result to int and return

	return result, err.NewExpressionError(err.ErrorConvertingResultToInt, "", l.Scanner.Pos()) //fmt.Errorf("error converting result to int")
}