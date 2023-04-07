package lexer

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func ParseArithmeticExpressions(expression string) (int, error) {
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
	if val, ok := result.(float64); ok {
		return int(val), nil
	}
	return 0, fmt.Errorf("Error converting result to int")
}
